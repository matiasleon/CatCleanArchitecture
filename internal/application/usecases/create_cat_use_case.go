package usecases

import (
	"awesomeProject/internal/application/dtos"
	"awesomeProject/internal/domain"
)

type CatCreateUseCase struct {
	catRepository CatRepository
}

func NewCatCreateUseCase(catRepository CatRepository) *CatCreateUseCase {
	return &CatCreateUseCase{catRepository: catRepository}
}

func (ccu *CatCreateUseCase) CreateCat(createCatRequest *dtos.CreateCatRequest) error {

	cat := domain.Cat{
		Name: *createCatRequest.Name,
		Age:  *createCatRequest.Age,
	}

	err := ccu.catRepository.CreateCat(&cat)
	if err != nil {
		return err
	}

	return nil
}
