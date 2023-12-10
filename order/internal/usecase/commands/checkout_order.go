package commands

import (
	"context"
	"shopping/order/internal/domain"
)

type CheckoutOrder struct {
	ID string
}

type CheckoutOrderHandler struct {
	orders   domain.OrderRepository
	payments domain.PaymentRepository
}

func NewCheckoutOrderHandler(orders domain.OrderRepository, payments domain.PaymentRepository) CheckoutOrderHandler {
	return CheckoutOrderHandler{
		orders:   orders,
		payments: payments,
	}
}

func (h CheckoutOrderHandler) CheckoutOrder(ctx context.Context, cmd CheckoutOrder) error {
	order, err := h.orders.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = order.Checkout(); err != nil {
		return err
	}

	if err = h.orders.Save(ctx, order); err != nil {
		return err
	}

	return nil
}
