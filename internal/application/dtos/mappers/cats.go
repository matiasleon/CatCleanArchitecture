package mappers

import (
	"awesomeProject/internal/application/dtos"
	"awesomeProject/internal/domain"
)

func mapCatsToCatResponse(cats []*domain.Cat) []*dtos.CatResponse {
	catsResponse := make([]*dtos.CatResponse, len(cats))
	for i, cat := range cats {
		catsResponse[i] = &dtos.CatResponse{
			Id:   cat.Id,
			Name: cat.Name,
			Age:  cat.Age,
		}
	}
	return catsResponse
}

func ToCatsResponse(cats []*domain.Cat) *dtos.CatsResponse {
	return &dtos.CatsResponse{
		Cats: mapCatsToCatResponse(cats),
	}
}
