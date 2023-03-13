package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mereiamangeldin/API/httpd/handlers"
)

func main() {

	r := gin.Default()
	r.GET("/", handlers.ToDoGet())
	r.GET("/:id", handlers.TaskGet())
	r.DELETE("/:id", handlers.TaskDelete())
	r.PUT("/:id", handlers.TaskUpdate())
	r.POST("/", handlers.ToDoPost())

	r.Run(":5000")
}
