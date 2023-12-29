package queries

import (
	"context"
	"shopping/order/internal/domain"
	"shopping/order/internal/models"
)

type (
	SearchingOrderHanlder struct {
		repo domain.SearchingRepository
	}
)

func NewSearchingOrderHandler(repo domain.SearchingRepository) *SearchingOrderHanlder {
	return &SearchingOrderHanlder{
		repo: repo,
	}
}

func (h SearchingOrderHanlder) SearchOrders(ctx context.Context, search domain.SearchOrders) ([]*models.Order, error) {
	// TODO implement me
	panic("implement me")
}
