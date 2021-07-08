package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/louisbillaut/test/api/handler"
	"github.com/louisbillaut/test/config"
	"log"
)

func main() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	r := gin.Default()
	handler.MakePublicityHandler(r, db)
	r.Run()
}