package commands

import (
	"context"
	"shopping/internal/ddd"
	"shopping/order/internal/domain"

	"github.com/google/wire"
)

type CompleteOrder struct {
	ID string
}

type CompleteOrderHandler struct {
	orders          domain.OrderRepository
	domainPublisher ddd.EventPublisher
}

var CompleteOrderUseCaseSet = wire.NewSet(NewCompleteOrderHandler)

func NewCompleteOrderHandler(orders domain.OrderRepository, domainPublisher ddd.EventPublisher) CompleteOrderHandler {
	return CompleteOrderHandler{
		orders:          orders,
		domainPublisher: domainPublisher,
	}
}

func (h CompleteOrderHandler) CompleteOrder(ctx context.Context, cmd CompleteOrder) error {
	order, err := h.orders.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = order.Complete(); err != nil {
		return err
	}

	if err = h.orders.Update(ctx, order); err != nil {
		return err
	}

	if err = h.domainPublisher.Publish(ctx, order.GetEvents()...); err != nil {
		return err
	}

	return nil
}
