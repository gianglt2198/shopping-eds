package commands

import (
	"context"
	"shopping/internal/ddd"
	"shopping/order/internal/domain"

	"github.com/google/wire"
)

type CheckoutOrder struct {
	ID string
}

type CheckoutOrderHandler struct {
	orders          domain.OrderRepository
	payments        domain.PaymentRepository
	domainPublisher ddd.EventPublisher
}

var CheckoutOrderUseCaseSet = wire.NewSet(NewCheckoutOrderHandler)

func NewCheckoutOrderHandler(orders domain.OrderRepository, payments domain.PaymentRepository, domainPublisher ddd.EventPublisher) CheckoutOrderHandler {
	return CheckoutOrderHandler{
		orders:          orders,
		payments:        payments,
		domainPublisher: domainPublisher,
	}
}

func (h CheckoutOrderHandler) CheckoutOrder(ctx context.Context, cmd CheckoutOrder) error {
	order, err := h.orders.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = order.Checkout(); err != nil {
		return nil
	}

	if err = h.orders.Update(ctx, order); err != nil {
		return err
	}

	if err = h.domainPublisher.Publish(ctx, order.GetEvents()...); err != nil {
		return err
	}

	return nil
}
