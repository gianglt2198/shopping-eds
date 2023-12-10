package commands

import (
	"context"
	"shopping/order/internal/domain"
)

type CancelOrder struct {
	ID string
}

type CancelOrderHandler struct {
	orders   domain.OrderRepository
	payments domain.PaymentRepository
}

func NewCancelOrderHandler(orders domain.OrderRepository, payments domain.PaymentRepository) CancelOrderHandler {
	return CancelOrderHandler{
		orders:   orders,
		payments: payments,
	}
}

func (h CancelOrderHandler) CancelOrder(ctx context.Context, cmd CancelOrder) error {
	order, err := h.orders.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = order.Cancel(); err != nil {
		return err
	}

	if err = h.orders.Save(ctx, order); err != nil {
		return err
	}

	return nil
}
