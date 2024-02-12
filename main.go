package main

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"

	"github.com/dpatsora/note-taker/adapters"
	"github.com/dpatsora/note-taker/pkg/config"
	"github.com/dpatsora/note-taker/pkg/logs"
	"github.com/dpatsora/note-taker/pkg/server"
	"github.com/dpatsora/note-taker/ports"

	"github.com/jmoiron/sqlx"
	"github.com/openware/pkg/database"
	"github.com/openware/pkg/kli"
	"github.com/pressly/goose/v3"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var (
	//go:embed config/migrations/*/*.sql
	embedMigrations embed.FS
)

func main() {
	logs.Init()

	cli := kli.NewCli("note-taker", "Note taking API", "3.0.0")
	cli.NewSubCommand("start", "Starts the API").Action(startAPI)

	dbCli := cli.NewSubCommand("db", "DB actions")
	dbCli.NewSubCommand("create", "Create database").Action(cmdCreateDB)
	dbCli.NewSubCommand("migrate", "Migrate database").Action(cmdMigrateDB)
	dbCli.NewSubCommand("drop", "Drop database").Action(cmdDropDB)

	if err := cli.Run(); err != nil {
		logrus.Fatal(err)
	}
}

func startAPI() error {
	cnf, err := config.Load()
	if err != nil {
		return err
	}

	db, err := setupDB(cnf.Database)
	if err != nil {
		return err
	}

	noteRepository := adapters.NewNotePostgresqlRepository(db)

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(ports.NewHttpServer(noteRepository), router)
	})

	return nil
}

func cmdCreateDB() error {
	cnf, err := config.Load()
	if err != nil {
		return err
	}

	logrus.Info("Creating schema")
	dbConf := cnf.Database
	dbConf.Schema = ""
	db, err := sqlx.Connect(dbConf.Driver, dbConf.URL())
	if err != nil {
		return err
	}

	queryDbCheck := fmt.Sprintf("SELECT 1 FROM information_schema.schemata WHERE schema_name='%s'", cnf.Database.Schema)
	if res, err := db.Exec(queryDbCheck); err != nil {
		return fmt.Errorf("error while checking schema existance: %s", err.Error())
	} else if rows, err := res.RowsAffected(); err != nil {
		return fmt.Errorf("error while checking schema existance: %s", err.Error())
	} else if rows > 0 {
		logrus.Warn("Schema already exist", "schema", cnf.Database.Schema)
		return nil
	}

	if _, err = db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s AUTHORIZATION %s", cnf.Database.Schema, cnf.Database.Username)); err != nil {
		return fmt.Errorf("error while creating database: %s", err.Error())
	}

	logrus.Info("Schema created")
	return nil
}

func cmdMigrateDB() error {
	cnf, err := config.Load()
	if err != nil {
		return err
	}

	db, err := goose.OpenDBWithDriver(cnf.Database.Driver, cnf.Database.URL())
	if err != nil {
		return err
	}

	logrus.Info("Applying database migrations")
	goose.SetBaseFS(embedMigrations)
	if err := goose.Up(db, "config/migrations/"+cnf.Database.Driver); err != nil {
		panic(err)
	}

	logrus.Info("Applied migrations")
	return nil
}

func cmdDropDB() error {
	cnf, err := config.Load()
	if err != nil {
		return err
	}

	logrus.Debug("Dropping schema")
	dbConf := cnf.Database
	dbConf.Schema = ""
	db, err := sqlx.Connect(dbConf.Driver, dbConf.URL())
	if err != nil {
		return err
	}

	if _, err = db.Exec(fmt.Sprintf("DROP SCHEMA %s CASCADE", cnf.Database.Schema)); err != nil {
		return err
	}

	logrus.Info("Dropped schema")
	return nil
}

func setupDB(cnf config.Database) (*gorm.DB, error) {
	dbConf := database.Config{
		Driver: cnf.Driver,
		Host:   cnf.Host,
		Port:   fmt.Sprint(cnf.Port),
		Name:   cnf.Name,
		User:   cnf.Username,
		Pass:   cnf.Password,
		Pool:   10,
		Schema: cnf.Schema,
	}

	db, err := database.Connect(&dbConf)
	if err != nil {
		return nil, err
	}

	return db, nil
}
