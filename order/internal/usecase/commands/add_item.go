package commands

import (
	"context"
	"shopping/order/internal/domain"
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

func NewAddItemHandler(orders domain.OrderRepository, products domain.ProductRepository) AddItemHandler {
	return AddItemHandler{
		orders:   orders,
		products: products,
	}
}

func (h AddItemHandler) AddItem(ctx context.Context, cmd AddItem) error {
	order, err := h.orders.Load(ctx, cmd.OrderID)
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

	if err = h.orders.Save(ctx, order); err != nil {
		return err
	}

	return nil
}
