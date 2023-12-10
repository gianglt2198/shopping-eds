package commands

import (
	"context"
	"shopping/product/internal/domain"
)

type (
	IncreasePrice struct {
		ID    string
		Price float64
	}

	IncreaseProductPriceHandler struct {
		products domain.ProductRepository
	}
)

func NewIncreaseProductPriceHandler(products domain.ProductRepository) IncreaseProductPriceHandler {
	return IncreaseProductPriceHandler{
		products: products,
	}
}

func (h IncreaseProductPriceHandler) IncreasePriceProduct(ctx context.Context, cmd IncreasePrice) error {
	product, err := h.products.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = product.IncreasePrice(cmd.Price); err != nil {
		return err
	}

	return h.products.Save(ctx, product)
}
