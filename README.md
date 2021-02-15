# todo-app
A sample todo app example written in go

Run `postgres` in docker container

```bash
docker run --name todo-db -e POSTGRES_PASSWORD=qwert -p 5432:5432 -d postgres
```

## Database migration

[Migrate](https://github.com/golang-migrate/migrate) is used for Database migrations.

Create migration files

```bash
migrate create -ext sql -dir db/migrations -seq init
```

Run DB migrations

```bash
migrate -path db/migrations -database 'postgres://postgres:qwert@localhost:5432/postgres?sslmode=disable' up
```

In case of error `error: Dirty database version 1. Fix and force version.` during migration run

```bash
migrate -path db/migrations -database 'postgres://postgres:qwert@localhost:5432/postgres?sslmode=disable' force 1
```

where `1` is a failed version of migration (see the name of sql file with migrations in `db/migrations`)

More details here: https://github.com/golang-migrate/migrate/issues/282#issuecomment-530743258

## Swagger

[Swagg](https://github.com/swaggo/swag) is used for Swagger documentation.

Generate swagger

```bash
swag init -g cmd/main.go
```

