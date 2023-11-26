package domain

import (
	"shopping/internal/ddd"

	"github.com/stackus/errors"
)

type Product struct {
	ddd.AggregateBase
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

	product := &Product{
		AggregateBase: ddd.AggregateBase{
			ID: id,
		},
		Name:        name,
		Price:       price,
		Description: description,
	}

	product.AddEvents(&ProductCreated{
		Product: product,
	})

	return product, nil
}

func (p *Product) Update() error {

	p.AddEvents(&ProductUpdated{Product: p})

	return nil
}

func (p *Product) Delete() error {
	p.AddEvents(&ProductDeleted{
		Product: p,
	})

	return nil
}
