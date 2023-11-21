package commands

import (
	"context"
	"shopping/order/internal/domain"

	"github.com/google/wire"
	"github.com/stackus/errors"
)

type CreateOrder struct {
	ID         string
	CustomerID string
	PaymentID  string
	Items      []*domain.Item
}

type CreateOrderHandler struct {
	orders    domain.OrderRepository
	customers domain.CustomerRepository
}

var CreateOrderUseCaseSet = wire.NewSet(NewCreateOrderHandler)

func NewCreateOrderHandler(orders domain.OrderRepository, customers domain.CustomerRepository) CreateOrderHandler {
	return CreateOrderHandler{
		orders:    orders,
		customers: customers,
	}
}

func (h CreateOrderHandler) CreateOrder(ctx context.Context, cmd CreateOrder) error {
	order, err := domain.CreateOrder(cmd.ID, cmd.CustomerID, cmd.PaymentID, cmd.Items)
	if err != nil {
		return errors.Wrap(err, "create order command")
	}

	// authorizeCustomer
	if err = h.customers.GetCustomer(ctx, order.CustomerID); err != nil {
		return errors.Wrap(err, "order customer authorization")
	}

	return errors.Wrap(h.orders.Save(ctx, order), "create order command")
}
