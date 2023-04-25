package http

import (
	"awesomeProject/internal/infrastructure/http"

	"github.com/gin-gonic/gin"
)

type mappers struct {
	catCreateHandler http.CatCreateHandler
}

// create func to create mappers
func NewMappers(catCreateHandler http.CatCreateHandler) *mappers {
	return &mappers{catCreateHandler: catCreateHandler}
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

	// return the server
	return r
}
