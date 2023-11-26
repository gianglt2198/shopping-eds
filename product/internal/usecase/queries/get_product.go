package queries

import (
	"context"
	"shopping/product/internal/domain"

	"github.com/google/wire"
)

type (
	GetProduct struct {
		ID string
	}

	GetProductHandler struct {
		products domain.ProductRepository
	}
)

var GetProductUsecaseSet = wire.NewSet(NewGetProductHandler)

func NewGetProductHandler(products domain.ProductRepository) GetProductHandler {
	return GetProductHandler{
		products: products,
	}
}

func (h GetProductHandler) GetProduct(ctx context.Context, query GetProduct) (*domain.Product, error) {
	return h.products.Find(ctx, query.ID)
}
