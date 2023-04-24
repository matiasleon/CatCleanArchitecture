package database

import "awesomeProject/internal/domain"

// generate a new cat repository using local memory
func NewCatRepository() *CatRepository {
	return &CatRepository{}
}

// CatRepository is a repository for cats
type CatRepository struct {
	cats []domain.Cat
}

// create a new cat
func (cr *CatRepository) CreateCat(cat *domain.Cat) error {
	cr.cats = append(cr.cats, *cat)
	return nil
}

// get cat by id
func (cr *CatRepository) GetCatById(id int) (*domain.Cat, error) {
	for _, cat := range cr.cats {
		if cat.Id == id {
			return &cat, nil
		}
	}
	return nil, nil
}

// get all cats
func (cr *CatRepository) GetAllCats() ([]*domain.Cat, error) {
	var cats []*domain.Cat
	for _, cat := range cr.cats {
		cats = append(cats, &cat)
	}

	return cats, nil
}
