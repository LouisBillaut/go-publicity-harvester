package _type

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/louisbillaut/test/config"
	"github.com/louisbillaut/test/entity"
	_type "github.com/louisbillaut/test/repository/type"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"testing"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func initDB() (*sql.DB, error){
	//must use a different database config for testing
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)
	return sql.Open("mysql", dataSourceName)
}

func TestCreate(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	db, err := initDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	typeRepo := _type.NewTypeMySQL(db)
	typeService := NewService(typeRepo)
	_, err = typeService.CreateType("visible")
	assert.NotNil(t, err)
	_, err = typeService.CreateType(randStringRunes(30))
	assert.Nil(t, err)
}

func TestGetByName(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	db, err := initDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	typeRepo := _type.NewTypeMySQL(db)
	typeService := NewService(typeRepo)
	_, err = typeService.GetTypeByName(randStringRunes(30))
	assert.NotNil(t, err)
	_, err = typeService.GetTypeByName("visible")
	assert.Nil(t, err)
}

func TestGetByID(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	db, err := initDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	typeRepo := _type.NewTypeMySQL(db)
	typeService := NewService(typeRepo)
	typeCreatedID, _ := typeService.CreateType(randStringRunes(30))
	_, err = typeService.GetTypeByID(typeCreatedID)
	assert.Nil(t, err)
	_, err = typeService.GetTypeByID(entity.NewID())
	assert.NotNil(t, err)
}
