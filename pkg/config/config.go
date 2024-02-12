package config

import (
	"fmt"

	"github.com/openware/pkg/ika"
)

type Config struct {
	Database  Database `yaml:"database"`
	JWTSecret string   `yaml:"jwt_secret" env:"JWT_SECRET" env-default:"mock_secret"`
	Port      uint64   `yaml:"app_port" env:"PORT" env-default:"8080"`
}

type Database struct {
	Name     string `yaml:"name" env:"DATABASE_NAME" env-default:"postgres"`
	Schema   string `yaml:"schema" env:"DATABASE_SCHEMA" env-default:"postgres"`
	Driver   string `yaml:"driver" env:"DATABASE_DRIVER" env-default:"postgres"`
	Username string `yaml:"username" env:"DATABASE_USERNAME" env-default:"postgres"`
	Password string `yaml:"password" env:"DATABASE_PASSWORD" env-default:"mock_secret"`
	Host     string `yaml:"host" env:"DATABASE_HOST" env-default:"0.0.0.0"`
	Port     uint64 `yaml:"port" env:"DATABASE_PORT" env-default:"5432"`
}

func Load() (*Config, error) {
	cnf := &Config{
		Database: Database{},
	}

	if err := ika.ReadConfig("", cnf); err != nil {
		return nil, fmt.Errorf("failed to parse config: %s", err.Error())
	}
	return cnf, nil
}

func (db *Database) URL() string {
	var dsn string

	if db.Name == "" {
		switch db.Driver {
		case "mysql":
			dsn = fmt.Sprintf("%s:%s@(%s:%d)?multiStatements=true&autocommit=true&parseTime=true",
				db.Username, db.Password, db.Host, db.Port,
			)
		case "postgres":
			dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
				db.Host, db.Port, db.Username, db.Password,
			)
		}
	} else {
		switch db.Driver {
		case "mysql":
			dsn = fmt.Sprintf("%s:%s@(%s:%d)/%s?multiStatements=true&autocommit=true&parseTime=true",
				db.Username, db.Password, db.Host, db.Port, db.Name,
			)
		case "postgres":
			dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable ",
				db.Host, db.Port, db.Username, db.Password, db.Name,
			)

			// if schema specified
			if db.Schema != "" {
				dsn = fmt.Sprintf("%s search_path=%s", dsn, db.Schema)
			}
		}
	}

	return dsn
}
