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

type CreateUserInput struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type JWTInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
