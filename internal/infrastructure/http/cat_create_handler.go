package http

import (
	"awesomeProject/internal/application/dtos"

	"github.com/gin-gonic/gin"
)

// CatCreateHandler is a handler for creating a new cat
type CatCreateHandler struct {
	catCreateUseCase CatCreateUseCase
}

type CatCreateUseCase interface { // define a usecase interface
	CreateCat(cat *dtos.CreateCatRequest) error
}

// create a new cat handler
func NewCatCreateHandler(catCreateUseCase CatCreateUseCase) *CatCreateHandler {
	return &CatCreateHandler{catCreateUseCase: catCreateUseCase}
}

// create a new cat handler
func (cch *CatCreateHandler) NewCatCreateHandler(c *gin.Context) {

	var createCatRequest dtos.CreateCatRequest
	if err := c.ShouldBindJSON(&createCatRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// create a new cat using the usecase
	if err := cch.catCreateUseCase.CreateCat(&createCatRequest); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
