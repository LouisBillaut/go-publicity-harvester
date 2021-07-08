# 🌾 Go Publicity Harvester 🌾
## 🧐 About this repository
Go Publicity Harvester is a GO API which allow to push some publicity event to a DB and
then collect these data using some filters.
It is using Clean architecture. 

### Developed with
* [Go](https://golang.org/)
* [MySQL](https://www.mysql.com/)

### Dependencies used
* [gin](https://github.com/gin-gonic/gin)
* [mysql](https://github.com/go-sql-driver/mysql)
* [uuid](https://github.com/google/uuid)
* [useragent](https://github.com/mileusna/useragent)
* [testify](https://github.com/stretchr/testify)

## 🤓 Getting started
### Prerequisite
Use `go mod download` command to install dependencies.
### DB
A SQL file is available at `/ops/db/init.sql` to instantiate database.
### DB Configuration
The configuration file for database is available at `/config/config.go`
### Run
The easiest way to run application is to run the `/api/main.go` file.


## 💻 API Endpoints
### Create event
`/createEvent` **GET** method:
allows to create a publicity event.
⚠️ You must add a `type` value in url params.

Example: `/createEvent?type=visible`
#### Available types:
* visible
* click
* enum_impression

### Get by OS
`/getByOS` **GET** method:
allows getting publicity by OS.
⚠️ You must add an `os` value in url params.

Example: `/getByOS?os=macOS`

### Get by Timestamp
`/getByTimestamp` **GET** method:
allows getting publicity by Timestamp.
⚠️ You must add a `from` and a `to` values in url params.

Example: `/getByTimestamp?from=2021-07-08&to=2021-07-09`
