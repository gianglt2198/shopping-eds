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
}

type CreateOrderHandler struct {
	orders    domain.OrderRepository
	customers domain.CustomerRepository
}

var CreateOrderUseCaseSet = wire.NewSet(NewCreateOrderHandler)

func NewCreateOrderHandler(
	orders domain.OrderRepository,
	customers domain.CustomerRepository) CreateOrderHandler {
	return CreateOrderHandler{
		orders:    orders,
		customers: customers,
	}
}

func (h CreateOrderHandler) CreateOrder(ctx context.Context, cmd CreateOrder) error {
	order, err := domain.CreateOrder(cmd.ID, cmd.CustomerID)
	if err != nil {
		return errors.Wrap(err, "create order command")
	}

	// authorizeCustomer
	if _, err = h.customers.GetCustomer(ctx, cmd.CustomerID); err != nil {
		return errors.Wrap(err, "order customer authorization")
	}

	if err = h.orders.Save(ctx, order); err != nil {
		return errors.Wrap(err, "create order command")
	}

	return nil
}
