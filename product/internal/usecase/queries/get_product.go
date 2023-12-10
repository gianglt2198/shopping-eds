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
		management domain.ManagementRepository
	}
)

var GetProductUsecaseSet = wire.NewSet(NewGetProductHandler)

func NewGetProductHandler(management domain.ManagementRepository) GetProductHandler {
	return GetProductHandler{
		management: management,
	}
}

func (h GetProductHandler) GetProduct(ctx context.Context, query GetProduct) (*domain.ManagementProduct, error) {
	return h.management.Find(ctx, query.ID)
}
