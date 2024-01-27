package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Bala de goma", "BALADEGOMA", 1500)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.Name)
	assert.NotEmpty(t, product.Code)
	assert.True(t, product.ValidateProduct() == nil)
}

func TestProduct_ValidateProduct(t *testing.T) {
	product, err := NewProduct("Bala de goma", "BALADEGOMA", 1500)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.True(t, product.ValidateProduct() == nil)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", "BALADEGOMA", 1500)

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenCodeIsRequired(t *testing.T) {
	product, err := NewProduct("Bala de goma", "", 1500)

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, ErrCodeIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	product, err := NewProduct("Bala de goma", "BALADEGOMA", -1500)

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidPrice, err)
}
