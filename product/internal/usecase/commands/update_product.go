package commands

import (
	"context"
	"shopping/internal/ddd"
	"shopping/product/internal/domain"

	"github.com/google/wire"
)

type (
	UpdateProduct struct {
		ID          string
		Name        string
		Description string
		Price       float64
	}

	UpdateProductHandler struct {
		products        domain.ProductRepository
		domainPublisher ddd.EventPublisher
	}
)

var UpdateProductUseCaseSet = wire.NewSet(NewUpdateProductHandler)

func NewUpdateProductHandler(products domain.ProductRepository, domainPublisher ddd.EventPublisher) UpdateProductHandler {
	return UpdateProductHandler{
		products:        products,
		domainPublisher: domainPublisher,
	}
}

func (h UpdateProductHandler) UpdateProduct(ctx context.Context, cmd UpdateProduct) error {
	product, err := h.products.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = product.Update(); err != nil {
		return err
	}

	if err = h.products.Update(ctx, product); err != nil {
		return err
	}

	if err = h.domainPublisher.Publish(ctx, product.GetEvents()...); err != nil {
		return err
	}

	return nil
}
