package commands

import (
	"context"
	"shopping/internal/ddd"
	"shopping/order/internal/domain"

	"github.com/google/wire"
)

type CancelOrder struct {
	ID string
}

type CancelOrderHandler struct {
	orders          domain.OrderRepository
	payments        domain.PaymentRepository
	domainPublisher ddd.EventPublisher
}

var CancelOrderUseCaseSet = wire.NewSet(NewCancelOrderHandler)

func NewCancelOrderHandler(orders domain.OrderRepository, payments domain.PaymentRepository, domainPublisher ddd.EventPublisher) CancelOrderHandler {
	return CancelOrderHandler{
		orders:          orders,
		payments:        payments,
		domainPublisher: domainPublisher,
	}
}

func (h CancelOrderHandler) CancelOrder(ctx context.Context, cmd CancelOrder) error {
	order, err := h.orders.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = order.Cancel(); err != nil {
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
