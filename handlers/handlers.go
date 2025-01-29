package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/LucasMacena09/GoLibrary/models"
	"github.com/LucasMacena09/GoLibrary/services"
)

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Books)
}

func BookById(c *gin.Context) {
	id := c.Param("id")
	book, err := services.GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.Books = append(models.Books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}
