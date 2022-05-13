package controllers

import (
	"bookselft/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBooks(c *gin.Context) {
	var inputBooks models.CreateBookInput
	if err := c.ShouldBindJSON(&inputBooks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{
		Title:  inputBooks.Title,
		Author: inputBooks.Author,
	}

	models.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindBooks(c *gin.Context) {
	var book models.Book

	if error := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBooks(c *gin.Context) {
	var book models.Book

	if error := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Record Not Found"})
		return
	}

	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	models.DB.Model(&book).Updates(&input)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBooks(c *gin.Context) {
	var book models.Book
	if error := models.DB.Where("id = ? ", c.Param("id")).First(&book).Error; error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	models.DB.Delete(&book)
}
