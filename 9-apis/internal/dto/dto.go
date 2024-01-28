package dto

type CreateProductInput struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Price int    `json:"price"`
}

type UpdateProductInput struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Price int    `json:"price"`
}
