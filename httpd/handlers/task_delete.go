package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mereiamangeldin/API/models"
	"github.com/mereiamangeldin/API/platform/toDo"
	"log"
	"net/http"
)

func TaskDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println(id)
		db, err := models.GetDB()
		if err != nil {
			log.Fatal(err)
		}
		var task toDo.Task
		result := db.Where("id=?", id).First(&task)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		err = db.Unscoped().Delete(&task).Error
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
	}
}
