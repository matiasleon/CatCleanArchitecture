package usecases

import "awesomeProject/internal/domain"

type CatRepository interface { // define a repository interface
	GetAllCats() ([]*domain.Cat, error)
	CreateCat(cat *domain.Cat) error
}
