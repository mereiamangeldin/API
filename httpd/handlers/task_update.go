package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mereiamangeldin/API/models"
	"github.com/mereiamangeldin/API/platform/toDo"
	"log"
)

func TaskUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := toDoPostRequest{}
		c.Bind(&requestBody)

		db, err := models.GetDB()
		if err != nil {
			log.Fatal(err)
		}
		var task toDo.Task
		result := db.Where("id=?", requestBody.ID).First(&task)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		task.Message = requestBody.Message
		task.Complete = requestBody.Complete
		fmt.Println(task.Message)
		db.Save(&task)
	}
}
