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
	engine.POST("/books", controllers.CreateBook)

	engine.Run()
}