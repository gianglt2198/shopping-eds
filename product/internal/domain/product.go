package domain

import (
	"github.com/stackus/errors"
)

type Product struct {
	ID          string
	Name        string
	Description string
	Price       float64
}

func CreateProduct(id, name, description string, price float64) (*Product, error) {
	if id == "" {
		return nil, errors.Wrap(errors.ErrBadRequest, "the product id cannot be blank")
	}

	if name == "" {
		return nil, errors.Wrap(errors.ErrBadRequest, "the product name cannot be blank")
	}

	if price == 0 {
		return nil, errors.Wrap(errors.ErrBadRequest, "the price cannot be blank")
	}

	return &Product{
		ID:          id,
		Name:        name,
		Price:       price,
		Description: description,
	}, nil
}

func UpdateProduct(id, name, description string, price float64) (*Product, error) {
	if id == "" {
		return nil, errors.Wrap(errors.ErrBadRequest, "the product id cannot be blank")
	}

	return &Product{
		ID:          id,
		Name:        name,
		Price:       price,
		Description: description,
	}, nil

}
