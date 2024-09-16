# Web App Template
> Go API with React frontend template

Template for full-stack web applications using Gin, Postgres, React, and Shadcn/Tailwind.

## About this Template
The main driver lives in `cmd/main.go`. This is where the logger, config, and store are
initialized and used to initialize an instance of the `models` package, which contains
all business logic. The `models` package defines the structs for models as necessary for 
requests, responses, and business logic methods. An implementation of the store interface
is required to instantiate the `models` package.

## Running with Docker
Create a file in the root of this directory named `.env.docker` and populate it based on `.env.template`.
```sh
docker compose --profile all up
```

## Running
Be sure to copy `.env.template`, rename to `.env`, and replace the relevant values.

```sh
# start Postgres and PGAdmin
docker compose --profile postgres up

# start API
go mod download
go run ./cmd/main.go

# start UI
cd ./web
npm install
npm start
```