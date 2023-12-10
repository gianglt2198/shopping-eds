package commands

import (
	"context"
	"shopping/product/internal/domain"

	"github.com/google/wire"
)

type (
	DeleteProduct struct {
		ID string
	}

	DeleteProductHandler struct {
		products domain.ProductRepository
	}
)

var DeleteProducUsecaseSet = wire.NewSet(NewDeleteProductHandler)

func NewDeleteProductHandler(products domain.ProductRepository) DeleteProductHandler {
	return DeleteProductHandler{
		products: products,
	}
}

func (h DeleteProductHandler) DeleteProduct(ctx context.Context, cmd DeleteProduct) error {
	product, err := h.products.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = product.Delete(); err != nil {
		return err
	}

	return h.products.Save(ctx, product)
}
