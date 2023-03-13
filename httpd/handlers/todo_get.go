package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mereiamangeldin/API/models"
	"github.com/mereiamangeldin/API/platform/toDo"
	"log"
)

import "net/http"

func ToDoGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := models.GetDB()
		if err != nil {
			log.Fatal(err)
		}
		var tasks []toDo.Task
		result := db.Order("id asc").Find(&tasks)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		c.JSON(http.StatusOK, tasks)
	}
}
