package commands

import (
	"context"
	"shopping/order/internal/domain"

	"github.com/google/wire"
)

type CancelOrder struct {
	ID string
}

type CancelOrderHandler struct {
	orders   domain.OrderRepository
	payments domain.PaymentRepository
}

var CancelOrderUseCaseSet = wire.NewSet(NewCancelOrderHandler)

func NewCancelOrderHandler(orders domain.OrderRepository, payments domain.PaymentRepository) CancelOrderHandler {
	return CancelOrderHandler{
		orders:   orders,
		payments: payments,
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

	if order.PaymentID != "" {
		if err = h.payments.CancelInvoice(ctx, order.PaymentID); err != nil {
			return err
		}
	}

	return h.orders.Update(ctx, order)
}
