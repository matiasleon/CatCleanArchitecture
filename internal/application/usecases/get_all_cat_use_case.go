package usecases

import (
	"awesomeProject/internal/application/dtos"
	"awesomeProject/internal/application/dtos/mappers"
)

// create get all cats usecase
func NewCatGetAllUseCase(catRepository CatRepository) *CatGetAllUseCase {
	return &CatGetAllUseCase{catRepository: catRepository}
}

// CatGetAllUseCase is a usecase for getting all cats
type CatGetAllUseCase struct {
	catRepository CatRepository
}

// get all cats
func (cgu *CatGetAllUseCase) GetAllCats() (*dtos.CatsResponse, error) {
	cats, err := cgu.catRepository.GetAllCats()

	catsResponse := mappers.ToCatsResponse(cats)

	if err != nil {
		return nil, err
	}

	return catsResponse, nil
}
