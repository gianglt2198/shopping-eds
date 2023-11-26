package commands

import (
	"context"
	"shopping/internal/ddd"
	"shopping/order/internal/domain"

	"github.com/google/wire"
)

type AddItem struct {
	OrderID  string
	ItemID   string
	Quantity int32
}

type AddItemHandler struct {
	orders          domain.OrderRepository
	products        domain.ProductRepository
	domainPublisher ddd.EventPublisher
}

var AddItemUseCaseSet = wire.NewSet(NewAddItemHandler)

func NewAddItemHandler(orders domain.OrderRepository, products domain.ProductRepository, domainPublisher ddd.EventPublisher) AddItemHandler {
	return AddItemHandler{
		orders:          orders,
		products:        products,
		domainPublisher: domainPublisher,
	}
}

func (h AddItemHandler) AddItem(ctx context.Context, cmd AddItem) error {
	order, err := h.orders.Find(ctx, cmd.OrderID)
	if err != nil {
		return err
	}

	product, err := h.products.GetProduct(ctx, cmd.ItemID)
	if err != nil {
		return err
	}

	if err = order.AddItem(product, cmd.Quantity); err != nil {
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
