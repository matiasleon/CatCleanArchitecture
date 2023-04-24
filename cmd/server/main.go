package server

import (
	"awesomeProject/internal/application/usecases"
	"awesomeProject/internal/infrastructure/database"
	"awesomeProject/internal/interfaces/http"
)

func main() {

	// create create cat handler with usecase
	catRepository := database.NewCatRepository()
	createCatUseCase := usecases.NewCatCreateUseCase(catRepository)
	createCatHandler := http.NewCatCreateHandler(createCatUseCase)

	// create get all cats handler with usecase and repository
	catGetAllUseCase := usecases.NewCatGetAllUseCase(catRepository)
	catGetAllHandler := http.NewCatGetAllHandler(catGetAllUseCase)

	mappers := NewMappers(*createCatHandler, *catGetAllHandler)
	server := NewServer(mappers)
	server.Run(":8080")
}
