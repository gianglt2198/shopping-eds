package commands

import (
	"context"
	"shopping/product/internal/domain"
)

type (
	DecreasePrice struct {
		ID    string
		Price float64
	}

	DecreaseProductPriceHandler struct {
		products domain.ProductRepository
	}
)

func NewDecreaseProductPriceHandler(products domain.ProductRepository) DecreaseProductPriceHandler {
	return DecreaseProductPriceHandler{
		products: products,
	}
}

func (h DecreaseProductPriceHandler) DecreasePriceProduct(ctx context.Context, cmd DecreasePrice) error {
	product, err := h.products.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = product.DecreasePrice(cmd.Price); err != nil {
		return err
	}

	return h.products.Save(ctx, product)
}
