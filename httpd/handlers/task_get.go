package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mereiamangeldin/API/models"
	"github.com/mereiamangeldin/API/platform/toDo"
	"log"
	"net/http"
)

func TaskGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		db, err := models.GetDB()
		if err != nil {
			log.Fatal(err)
		}
		var task toDo.Task
		result := db.Where("id=?", id).First(&task)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		c.JSON(http.StatusOK, task)
	}
}
