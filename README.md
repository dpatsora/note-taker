# Note Taking API


# Description
This is a simple API that allows you to create, read, update, and delete notes. It is built using Golang and PostgreSQL.

All the endpoints are secured using JWT tokens. The API is documented using Swagger.
You can find OpenAPI specification in `open-api` folder

# Pre-requisites
- Docker
- Docker Compose
- Make


# How to run the API
```bash
make start
```

This command will start the API and the database in Docker. The API will be available at `http://localhost:8080`


If you want to run application without docker, you can run the following commands:
```bash
go run . db create
go run . db migrate
go run . start
```


Example JWT token - `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.FhQnRoRWaTyZ5_GNKhX-6XGjB-L9JUGOdCbxDd-8ffk`

# How to run Swagger UI
```bash
make swagger-ui
```

This command will start the Swagger UI and it will be available at `http://localhost:3000`


# Project Structure

### Adapters
Adapters package is responsible for interacting with the database. It contains the implementation of the repository interface.

### Config
Config directory stores docker-compose file and migration files for the database.

### Domain 
Domain package contains the business logic of the application. It contains the definition of the repository interface and the use case.

### Open-API
Open-API directory contains the OpenAPI specification of the API.

### Ports
Ports package contains the implementation of the HTTP handlers and OpenAPI generated types, server interface and client.

### Pkg
Pkg directory contains the common utilities and middlewares used in the application.

### Pkg/Auth
Auth package contains the implementation of the JWT middleware.

### Pkg/config
Config package contains the configuration of the application.

### Pkg/errors
Errors package contains the custom error types used in the application.

### Pkg/logs
Logs package contains the logger configuration.

### Pkg/server
Server package contains the setup of the basic http router.