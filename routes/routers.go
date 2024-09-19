package routes

import (
	"github.com/galantixa/go-gin-gorm/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/books", controllers.GetBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books:id", controllers.GetBookByID)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	return r

}
