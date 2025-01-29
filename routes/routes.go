package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/LucasMacena09/GoLibrary/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	
	router.GET("/books", handlers.GetBooks)
	router.GET("/books/:id", handlers.BookById)
	router.POST("/books", handlers.CreateBook)

	return router
}