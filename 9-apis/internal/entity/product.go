package entity

import (
	"errors"
	"github.com/jfromjefferson/gi-course-9/pkg/entity"
	"time"
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}

var (
	ErrNameIsRequired = errors.New("name is required")
	ErrCodeIsRequired = errors.New("code is required")
	ErrInvalidPrice   = errors.New("invalid price")
)

func NewProduct(name, code string, price int) (*Product, error) {
	product := Product{
		ID:        entity.NewID(),
		Name:      name,
		Code:      code,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := product.ValidateProduct()
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (product *Product) ValidateProduct() error {
	if product.Name == "" {
		return ErrNameIsRequired
	}

	if product.Code == "" {
		return ErrCodeIsRequired
	}

	if product.Price <= 0 {
		return ErrInvalidPrice
	}

	return nil
}
