package main

import (
	"github.com/FelipeMP0/go-rest-api-gin/v2/controllers"
	"github.com/FelipeMP0/go-rest-api-gin/v2/models"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	
	models.ConnectDatabase()

	engine.GET("/books", controllers.GetBooks)
	engine.GET("/books/:id", controllers.GetBook)
	engine.POST("/books", controllers.CreateBook)
	engine.PATCH("/books/:id", controllers.UpdateBook)
	engine.DELETE("/books/:id", controllers.DeleteBook)

	engine.Run()
}