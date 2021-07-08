package publicity_type

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/louisbillaut/test/config"
	"github.com/louisbillaut/test/entity"
	"github.com/louisbillaut/test/repository/publicity_type"
	"github.com/stretchr/testify/assert"
	"log"

	"testing"
)

func initDB() (*sql.DB, error){
	//must use a different database config for testing
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)
	return sql.Open("mysql", dataSourceName)
}

func TestGetPublicityTypeByPublicityID(t *testing.T) {
	db, err := initDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	publicityTypeRepo := publicity_type.NewPublicityTypeMySQL(db)
	publicityTypeService := NewService(publicityTypeRepo)
	randomID := entity.NewID()
	_, err = publicityTypeService.GetPublicityTypeByPublicityID(randomID)
	assert.NotNil(t, err)
}
