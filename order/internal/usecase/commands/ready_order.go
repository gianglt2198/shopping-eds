package commands

import (
	"context"
	"shopping/order/internal/domain"

	"github.com/google/wire"
)

type ReadyOrder struct {
	ID string
}

type ReadyOrderHandler struct {
	orders   domain.OrderRepository
	payments domain.PaymentRepository
}

var ReadyOrderUseCaseSet = wire.NewSet(NewReadyOrderHandler)

func NewReadyOrderHandler(orders domain.OrderRepository, payments domain.PaymentRepository) ReadyOrderHandler {
	return ReadyOrderHandler{
		orders:   orders,
		payments: payments,
	}
}

func (h ReadyOrderHandler) ReadyOrder(ctx context.Context, cmd ReadyOrder) (string, error) {
	order, err := h.orders.Find(ctx, cmd.ID)
	if err != nil {
		return "", err
	}

	if order.PaymentID, err = h.payments.CreateInvoice(ctx, order.ID, order.CustomerID, order.GetTotal()); err != nil {
		return "", err
	}

	if err = order.Ready(); err != nil {
		return "", nil
	}

	return order.PaymentID, h.orders.Update(ctx, order)
}
