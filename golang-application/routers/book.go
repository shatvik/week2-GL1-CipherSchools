package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shatvik/week2-GL1-CipherSchools.git/database"
	"github.com/shatvik/week2-GL1-CipherSchools.git/handler"
)

func Router() *gin.Engine {
	router := gin.Default()
	api := handler.Handler{
		DB: database.GetDB(),
	}
	router.GET("/books", api.GetBooks)
	router.POST("/books", api.SaveBook)
	router.DELETE("/books", api.DeleteBookByID)
	router.PUT("/books", api.UpdateBookByID)
	return router
}
