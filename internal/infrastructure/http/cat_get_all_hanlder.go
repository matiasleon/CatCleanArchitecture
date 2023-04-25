package http

import (
	"awesomeProject/internal/application/dtos"

	"github.com/gin-gonic/gin"
)

// CatGetAllHandler is a handler for getting all cats
type CatGetAllHandler struct {
	catGetAllUseCase CatGetAllUseCase
}

type CatGetAllUseCase interface { // define a usecase interface
	GetAllCats() (*dtos.CatsResponse, error)
}

// create a new cat handler
func NewCatGetAllHandler(catGetAllUseCase CatGetAllUseCase) *CatGetAllHandler {
	return &CatGetAllHandler{catGetAllUseCase: catGetAllUseCase}
}

// create get all cats handler
func (cgh *CatGetAllHandler) CatGetAllHandler(c *gin.Context) {

	// get all cats using the usecase
	cats, err := cgh.catGetAllUseCase.GetAllCats()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"cats": cats,
	})
}
