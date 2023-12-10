package commands

import (
	"context"
	"shopping/order/internal/domain"

	"github.com/google/wire"
)

type CompleteOrder struct {
	ID string
}

type CompleteOrderHandler struct {
	orders domain.OrderRepository
}

var CompleteOrderUseCaseSet = wire.NewSet(NewCompleteOrderHandler)

func NewCompleteOrderHandler(orders domain.OrderRepository) CompleteOrderHandler {
	return CompleteOrderHandler{
		orders: orders,
	}
}

func (h CompleteOrderHandler) CompleteOrder(ctx context.Context, cmd CompleteOrder) error {
	order, err := h.orders.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = order.Complete(); err != nil {
		return err
	}

	if err = h.orders.Save(ctx, order); err != nil {
		return err
	}

	return nil
}
