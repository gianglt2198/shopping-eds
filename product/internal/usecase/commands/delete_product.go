package commands

import (
	"context"
	"shopping/internal/ddd"
	"shopping/product/internal/domain"

	"github.com/google/wire"
)

type (
	DeleteProduct struct {
		ID string
	}

	DeleteProductHandler struct {
		products        domain.ProductRepository
		domainPublisher ddd.EventPublisher
	}
)

var DeleteProducUsecaseSet = wire.NewSet(NewDeleteProductHandler)

func NewDeleteProductHandler(products domain.ProductRepository, domainPublisher ddd.EventPublisher) DeleteProductHandler {
	return DeleteProductHandler{
		products:        products,
		domainPublisher: domainPublisher,
	}
}

func (h DeleteProductHandler) DeleteProduct(ctx context.Context, cmd DeleteProduct) error {
	product, err := h.products.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = product.Delete(); err != nil {
		return err
	}

	if err = h.products.Delete(ctx, cmd.ID); err != nil {
		return err
	}

	return nil
}
