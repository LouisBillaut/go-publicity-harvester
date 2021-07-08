package handler

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/louisbillaut/test/api/presenter"
	publicity2 "github.com/louisbillaut/test/repository/publicity"
	"github.com/louisbillaut/test/usecase/publicity"
	"net/http"
	"time"
)

func createPublicity(publicityService publicity.UseCase) func(*gin.Context){
	return func(c *gin.Context) {
		if len(c.Request.URL.Query()) == 0 || len(c.Request.URL.Query()["type"]) == 0 {
			c.String(http.StatusInternalServerError, "you must give a type")
			return
		}
		pub, err := publicityService.CreatePublicity(c.Request.URL.Query()["type"][0], c.Request.UserAgent(), c.ClientIP())
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
			return
		}
		getPublicity, err := publicityService.GetPublicity(pub)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
			return
		}
		c.JSON(http.StatusCreated, presenter.ResponsePublicityCreated(getPublicity))
	}
}

func getByOS(publicityService publicity.UseCase) func(*gin.Context){
	return func(c *gin.Context) {
		if len(c.Request.URL.Query()) == 0 || len(c.Request.URL.Query()["os"]) == 0 {
			c.String(http.StatusInternalServerError, "you must give an os")
			return
		}
		pub, err := publicityService.GetByOS(c.Request.URL.Query()["os"][0])
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
			return
		}
		c.JSON(http.StatusOK, presenter.ResponsePublicityList(pub))
	}
}

func getByTimestamp(publicityService publicity.UseCase) func(*gin.Context){
	return func(c *gin.Context) {
		if len(c.Request.URL.Query()) == 0 || len(c.Request.URL.Query()["from"]) == 0 || len(c.Request.URL.Query()["to"]) == 0{
			c.String(http.StatusInternalServerError, "you must give a `from` timestamp and a `to` timestamp")
			return
		}
		parsedFrom, err := time.Parse("2006-01-02", c.Request.URL.Query()["from"][0])
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
			return
		}
		parsedTo, err := time.Parse("2006-01-02", c.Request.URL.Query()["to"][0])
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
			return
		}
		pub, err := publicityService.GetByTimeStamp(parsedFrom.UTC(), parsedTo.UTC())
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
			return
		}
		c.JSON(http.StatusOK, presenter.ResponsePublicityList(pub))
	}
}

func MakePublicityHandler(r *gin.Engine, db *sql.DB) {
	publicityRepo := publicity2.NewPublicityMySQL(db)
	publicityService := publicity.NewService(publicityRepo)

	r.GET("/createEvent", createPublicity(publicityService))

	r.GET("/getByOS", getByOS(publicityService))

	r.GET("/getByTimestamp", getByTimestamp(publicityService))
}
