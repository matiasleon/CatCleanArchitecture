package dtos

// cats array response
type CatsResponse struct {
	Cats []*CatResponse `json:"cats"`
}

// cat response
type CatResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
