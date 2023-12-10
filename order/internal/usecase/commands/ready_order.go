package commands

import (
	"context"
	"shopping/order/internal/domain"
)

type ReadyOrder struct {
	ID        string
	PaymentID string
}

type ReadyOrderHandler struct {
	orders   domain.OrderRepository
	payments domain.PaymentRepository
}

func NewReadyOrderHandler(orders domain.OrderRepository, payments domain.PaymentRepository) ReadyOrderHandler {
	return ReadyOrderHandler{
		orders:   orders,
		payments: payments,
	}
}

func (h ReadyOrderHandler) ReadyOrder(ctx context.Context, cmd ReadyOrder) error {
	order, err := h.orders.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err := h.payments.GetInvoice(ctx, cmd.PaymentID); err != nil {
		return err
	}

	if err = order.Ready(cmd.PaymentID); err != nil {
		return nil
	}

	if err = h.orders.Save(ctx, order); err != nil {
		return err
	}

	return nil
}
