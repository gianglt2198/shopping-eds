package commands

import (
	"context"
	"shopping/internal/ddd"
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
	orders          domain.OrderRepository
	customers       domain.CustomerRepository
	domainPublisher ddd.EventPublisher
}

var CreateOrderUseCaseSet = wire.NewSet(NewCreateOrderHandler)

func NewCreateOrderHandler(
	orders domain.OrderRepository,
	customers domain.CustomerRepository,
	domainPublisher ddd.EventPublisher) CreateOrderHandler {
	return CreateOrderHandler{
		orders:          orders,
		customers:       customers,
		domainPublisher: domainPublisher,
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

	if err = h.orders.Save(ctx, order); err != nil {
		return errors.Wrap(err, "create order command")
	}

	if err = h.domainPublisher.Publish(ctx, order.GetEvents()...); err != nil {
		return err
	}

	return nil
}
