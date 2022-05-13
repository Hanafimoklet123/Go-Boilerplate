package main

import (
	"bookselft/controllers"
	"bookselft/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()

	router.GET("/books", controllers.AllBooks)
	router.POST("/books", controllers.CreateBooks)
	router.GET("/books/:id", controllers.FindBooks)
	router.PUT("/books/:id", controllers.UpdateBooks)
	router.DELETE("/books/:id", controllers.DeleteBooks)

	router.Run(":8000")
}
