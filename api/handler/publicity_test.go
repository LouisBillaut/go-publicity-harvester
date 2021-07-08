package handler

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/louisbillaut/test/config"
	publicity2 "github.com/louisbillaut/test/repository/publicity"
	"github.com/louisbillaut/test/usecase/publicity"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func initDB() (*sql.DB, error){
	//must use a different database config for testing
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)
	return sql.Open("mysql", dataSourceName)
}

func TestCreatePublicity(t *testing.T) {
	db, err := initDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	publicityRepo := publicity2.NewPublicityMySQL(db)
	publicityService := publicity.NewService(publicityRepo)
	typeTest := "visible"
	r := gin.Default()
	r.GET("/createEvent", createPublicity(publicityService))
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/createEvent?type="+typeTest, nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/createEvent?ttt="+typeTest, nil)
	r.ServeHTTP(w2, req2)
	assert.Equal(t, http.StatusInternalServerError, w2.Code)
}

func TestGetByOS(t *testing.T) {
	db, err := initDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	publicityRepo := publicity2.NewPublicityMySQL(db)
	publicityService := publicity.NewService(publicityRepo)
	os := "visible"
	r := gin.Default()
	r.GET("/getByOS", getByOS(publicityService))
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getByOS?os="+os, nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/getByOS?s=somethings", nil)
	r.ServeHTTP(w2, req2)
	assert.Equal(t, http.StatusInternalServerError, w2.Code)
}

func TestGetByTimeStamp(t *testing.T) {
	db, err := initDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	publicityRepo := publicity2.NewPublicityMySQL(db)
	publicityService := publicity.NewService(publicityRepo)
	r := gin.Default()
	r.GET("/getByTimestamp", getByTimestamp(publicityService))
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getByTimestamp?from=2021-07-08&to=2021-07-09", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/getByTimestamp?from=2021-07-08", nil)
	r.ServeHTTP(w2, req2)
	assert.Equal(t, http.StatusInternalServerError, w2.Code)
}