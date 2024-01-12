# go-starter

Golang starter project.

## Stack

- [task](https://taskfile.dev) for task management
- [goose](https://pressly.github.io/goose/) for database migrations
- [pgx](https://github.com/jackc/pgx) as a PostgreSQL driver
- [sqlc](https://sqlc.dev/) for generating code from SQL queries
- [chi](https://go-chi.io/) for routing
- [validator](https://pkg.go.dev/github.com/go-playground/validator/v10) for validating input data
- [templ](https://templ.guide/) as a template engine
- [picocss](https://picocss.com/) for styles


## Install dev dependencies

```bash
chmod +x ./scripts/install-deps.sh
./scripts/install-deps.sh
```

## Create migration file

```bash
task goose -- create <migration_name> sql
```

## Start server in watch mode

```bash
task
```
