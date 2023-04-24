package dtos

type CreateCatRequest struct {
	Name *string `json:"name" binding:"required" required:"true"` // Name of the cat
	Age  *int    `json:"age" binding:"required" required:"true"`  // Age of the cat
}
