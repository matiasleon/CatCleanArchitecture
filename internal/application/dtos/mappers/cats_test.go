package mappers

import (
	"awesomeProject/internal/domain"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMapCatsToCatsResponseShouldFullfillAllFields(t *testing.T) {
	// given
	cats := []*domain.Cat{
		{
			Id:   1,
			Name: "name",
			Age:  1,
		},
	}
	// when
	catsResponse := mapCatsToCatResponse(cats)
	// then
	assert.Equal(t, cats[0].Id, catsResponse[0].Id)
	assert.Equal(t, cats[0].Name, catsResponse[0].Name)
	assert.Equal(t, cats[0].Age, catsResponse[0].Age)

}
