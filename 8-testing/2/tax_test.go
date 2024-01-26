package tax

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.0)

	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = CalculateTax(0)
	assert.NotNil(t, err, "eita :(")
	assert.Equal(t, 0.0, tax)
}

func TestCalculateTaxAndSave(t *testing.T) {
	repositoryMock := &RepositoryMock{}

	repositoryMock.On("SaveTax", 10.0).Return(nil)
	repositoryMock.On("SaveTax", 0.0).Return(errors.New("amount must be greater than 0"))

	err := CalculateTaxAndSave(1000.0, repositoryMock)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(-1, repositoryMock)
	assert.NotNil(t, err)

	repositoryMock.AssertExpectations(t)
	repositoryMock.AssertNumberOfCalls(t, "SaveTax", 2)
}
