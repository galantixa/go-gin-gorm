package controllers

import (
	"github.com/galantixa/go-gin-gorm/database"
	"github.com/galantixa/go-gin-gorm/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	if err := database.DB.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": book})
}
func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.Where("id = ?", id).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.Where("id = ?", id).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book Not Found!"})
		return
	}
	database.DB.Save(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.Where("id = ?", id).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book Not Found!"})
		return
	}
	if err := database.DB.Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book Deleted!"})
}
