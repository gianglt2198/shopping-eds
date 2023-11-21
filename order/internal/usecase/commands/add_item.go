package commands

import (
	"context"
	"shopping/order/internal/domain"

	"github.com/google/wire"
)

type AddItem struct {
	OrderID  string
	ItemID   string
	Quantity int32
}

type AddItemHandler struct {
	orders   domain.OrderRepository
	products domain.ProductRepository
}

var AddItemUseCaseSet = wire.NewSet(NewAddItemHandler)

func NewAddItemHandler(orders domain.OrderRepository, products domain.ProductRepository) AddItemHandler {
	return AddItemHandler{
		orders:   orders,
		products: products,
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

	order.AddItem(product, cmd.Quantity)

	return h.orders.Update(ctx, order)
}
