
# How to Run, and Needed

## Getting started

### Requirements
- go version >= 1.20
- Makefile
- mockery v2.32.0
- swag


## Usage

### Config
Clone config file `config.yaml.example` from directory `/config/files`, put it on the same directory and rename it to `config.yaml`

You can also define config from env.


### API Server
running http server using
```
make run-api
```
or go command
```
go run main.go serve-http
```



### DB Migration
Reference:
- [go-migrate](https://github.com/golang-migrate/migrate)
- [go-migrate for PostgreSQL](https://github.com/golang-migrate/migrate/tree/master/database/postgres)

#### Install Go Migrate
```
$ go install -tags "postgres" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```
#### Create Migration File
```
$ migrate create -ext sql -dir db/migrations [name_of_migration_file]
```
#### Migrate up
```
$ migrate -database "postgres://[user]:[password]@[host]:[port]/[dbname]?query" -path db/migrations up
```
#### Migrate down
```
$ migrate -database "postgres://[user]:[password]@[host]:[port]/[dbname]?query" -path db/migrations down
```


accessing swagger docs using
```
http://localhost:8900/docs/index.html
```
