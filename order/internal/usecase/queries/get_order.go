package queries

import (
	"context"
	"shopping/order/internal/domain"

	"github.com/google/wire"
	"github.com/stackus/errors"
)

type GetOrder struct {
	ID string
}

type GetOrderHandler struct {
	repo domain.OrderRepository
}

var GetOrderUseCaseSet = wire.NewSet(NewGetOrderHandler)

func NewGetOrderHandler(repo domain.OrderRepository) GetOrderHandler {
	return GetOrderHandler{repo: repo}
}

func (h GetOrderHandler) GetOrder(ctx context.Context, query GetOrder) (*domain.Order, error) {
	order, err := h.repo.Load(ctx, query.ID)

	return order, errors.Wrap(err, "get order query")
}
