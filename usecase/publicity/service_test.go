package publicity

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/louisbillaut/test/config"
	"github.com/louisbillaut/test/entity"
	"github.com/louisbillaut/test/repository/publicity"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func initDB() (*sql.DB, error){
	//must use a different database config for testing
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)
	return sql.Open("mysql", dataSourceName)
}

func TestCreate(t *testing.T) {
	db, err := initDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	publicityRepo := publicity.NewPublicityMySQL(db)
	publicityService := NewService(publicityRepo)

	badType := "nonexistent type"
	goodType := "visible"
	ua := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36"
	ip := "localhost"

	_, err = publicityService.CreatePublicity(badType, ua, ip)
	assert.NotNil(t, err)
	_, err = publicityService.CreatePublicity(goodType, ua, ip)
	assert.Nil(t, err)
}

func TestGetPublicity(t *testing.T) {
	db, err := initDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	publicityRepo := publicity.NewPublicityMySQL(db)
	publicityService := NewService(publicityRepo)

	typeTest := "visible"
	ua := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36"
	ip := "localhost"
	pubID, _ := publicityService.CreatePublicity(typeTest, ua, ip)
	_, err = publicityService.GetPublicity(entity.NewID())
	assert.NotNil(t, err)
	publicity, err := publicityService.GetPublicity(pubID)
	assert.NotNil(t, pubID)
	assert.Equal(t, publicity.ID, pubID)
	assert.Equal(t, publicity.Type, typeTest)
	assert.Equal(t, publicity.Ua, ua)
	assert.Equal(t, publicity.Ip, ip)
}

func TestGetByOS(t *testing.T) {
	db, err := initDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	publicityRepo := publicity.NewPublicityMySQL(db)
	publicityService := NewService(publicityRepo)

	typeTest := "visible"
	macUa := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36"
	linuxUa := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36"
	ip := "localhost"
	os := "macOS"

	_, _ = publicityService.CreatePublicity(typeTest, macUa, ip)
	_, _ = publicityService.CreatePublicity(typeTest, linuxUa, ip)

	byOS, err := publicityService.GetByOS(os)
	assert.Nil(t, err)
	for _, b := range byOS {
		assert.Equal(t, os, b.Os)
	}
}

func TestGetByTimeStamp(t *testing.T) {
	db, err := initDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	publicityRepo := publicity.NewPublicityMySQL(db)
	publicityService := NewService(publicityRepo)

	typeTest := "visible"
	ua := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36"
	ip := "localhost"
	date := "2021-07-07"
	from, _ := time.Parse("2006-01-02", date)
	to := time.Now().Add(time.Hour * 48)
	_, _ = publicityService.CreatePublicity(typeTest, ua, ip)
	byTimeStamp, err := publicityService.GetByTimeStamp(from, to)
	assert.Nil(t, err)
	for _, b := range byTimeStamp {
		assert.True(t, b.CreatedAt.Before(to))
	}
}