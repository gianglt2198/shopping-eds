package domain

import (
	"shopping/internal/ddd"
	"shopping/internal/es"

	"github.com/stackus/errors"
)

const ProductAggregate = "products.Product"

var (
	ErrProductNameIsBlank     = errors.Wrap(errors.ErrBadRequest, "the product name cannot be blank")
	ErrProductPriceIsNegative = errors.Wrap(errors.ErrBadRequest, "the product price cannot be negative")
	ErrNotAPriceIncrease      = errors.Wrap(errors.ErrBadRequest, "the price change would be a decrease")
	ErrNotAPriceDecrease      = errors.Wrap(errors.ErrBadRequest, "the price change would be an increase")
)

type Product struct {
	es.Aggregate
	Name        string
	Description string
	Price       float64
}

var _ interface {
	es.EventApplier
	es.Snapshotter
} = (*Product)(nil)

func NewProduct(id string) *Product {
	return &Product{
		Aggregate: es.NewAggregate(id, ProductAggregate),
	}
}

func (Product) Key() string { return ProductAggregate }

func CreateProduct(id, name, description string, price float64) (*Product, error) {
	if name == "" {
		return nil, ErrProductNameIsBlank
	}

	if price < 0 {
		return nil, ErrProductPriceIsNegative
	}

	product := NewProduct(id)

	product.AddEvent(ProductCreatedEvent, &ProductCreated{
		Name:        name,
		Description: description,
		Price:       price,
	})

	return product, nil
}

func (p *Product) Delete() error {
	p.AddEvent(ProductDeletedEvent, &ProductDeleted{})

	return nil
}

func (p *Product) IncreasePrice(price float64) error {
	if price < p.Price {
		return ErrNotAPriceIncrease
	}

	p.AddEvent(ProductInscreasedPriceEvent, &ProductPriceChanged{
		Delta: price - p.Price,
	})

	return nil
}

func (p *Product) DecreasePrice(price float64) error {
	if price > p.Price {
		return ErrNotAPriceDecrease
	}

	p.AddEvent(ProductInscreasedPriceEvent, &ProductPriceChanged{
		Delta: price - p.Price,
	})

	return nil
}

func (p *Product) ApplyEvent(event ddd.Event) error {
	switch payload := event.Payload().(type) {
	case *ProductCreated:
		p.Name = payload.Name
		p.Description = payload.Description
		p.Price = payload.Price
	case *ProductDeleted:
		//
	case *ProductPriceChanged:
		p.Price = p.Price + payload.Delta
	default:
		return errors.ErrInternal.Msgf("%T received the event %s with unexpected payload %T", p, event.EventName(), payload)
	}

	return nil
}

func (p *Product) ApplySnapshot(snapshot es.Snapshot) error {
	switch ss := snapshot.(type) {
	case *ProductV1:
		p.Name = ss.Name
		p.Description = ss.Description
		p.Price = ss.Price
	default:
		return errors.ErrInternal.Msgf("%T received the unexpected snapshot %T", p, snapshot)
	}

	return nil
}

func (p Product) ToSnapshot() es.Snapshot {
	return ProductV1{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}
}
