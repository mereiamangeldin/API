package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mereiamangeldin/API/models"
	"github.com/mereiamangeldin/API/platform/toDo"
	"log"
	"net/http"
)

type toDoPostRequest struct {
	ID       int    `json:"id"`
	Message  string `json:"message"`
	Complete bool   `json:"complete"`
}

func ToDoPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := toDoPostRequest{}
		c.Bind(&requestBody)

		db, err := models.GetDB()
		if err != nil {
			log.Fatal(err)
		}
		var lastTask toDo.Task
		err = db.Last(&lastTask).Error
		var lastTaskId int

		if err != nil {
			lastTaskId = 1
		} else {
			lastTaskId = lastTask.ID + 1
		}

		task := toDo.Task{
			ID:       lastTaskId,
			Message:  requestBody.Message,
			Complete: requestBody.Complete,
		}
		result := db.Create(&task)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		c.Status(http.StatusNoContent)
	}
}
