package server

import (
	"awesomeProject/internal/interfaces/http"

	"github.com/gin-gonic/gin"
)

type mappers struct {
	catCreateHandler http.CatCreateHandler
	catGetAllHanlder http.CatGetAllHandler
}

// create func to create mappers
func NewMappers(catCreateHandler http.CatCreateHandler, catGetAllHanlder http.CatGetAllHandler) *mappers {
	return &mappers{
		catCreateHandler: catCreateHandler,
		catGetAllHanlder: catGetAllHanlder,
	}
}

func NewServer(mappers *mappers) *gin.Engine {
	r := gin.Default()

	// define a routes
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/cats", mappers.catCreateHandler.NewCatCreateHandler)

	r.GET("/cats", mappers.catGetAllHanlder.CatGetAllHandler)

	// return the server
	return r
}
