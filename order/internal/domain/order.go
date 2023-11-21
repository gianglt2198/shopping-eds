package domain

import (
	"github.com/stackus/errors"
)

var (
	ErrOrderHasNoItems         = errors.Wrap(errors.ErrBadRequest, "the order has no items")
	ErrOrderCannotBeCancelled  = errors.Wrap(errors.ErrBadRequest, "the order cannot be cancelled")
	ErrCustomerIDCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the customer id cannot be blank")
	ErrPaymentIDCannotBeBlank  = errors.Wrap(errors.ErrBadRequest, "the payment id cannot be blank")
)

type (
	Order struct {
		ID         string
		CustomerID string
		PaymentID  string
		Status     OrderStatus
		Items      map[string]*Item
	}

	Item struct {
		ProductID   string
		ProductName string
		Price       float64
		Quantity    int32
	}
)

func CreateOrder(id, customerID, paymentID string, items []*Item) (*Order, error) {
	if len(items) == 0 {
		return nil, ErrOrderHasNoItems
	}

	if customerID == "" {
		return nil, ErrCustomerIDCannotBeBlank
	}

	orderItems := make(map[string]*Item)
	for _, item := range items {
		orderItems[item.ProductID] = item
	}

	order := &Order{
		ID:         id,
		CustomerID: customerID,
		PaymentID:  paymentID,
		Items:      orderItems,
		Status:     OrderPending,
	}

	return order, nil
}

func (o *Order) Cancel() error {
	if o.Status != OrderPending {
		return ErrOrderCannotBeCancelled
	}

	o.Status = OrderCancelled

	return nil
}

func (o *Order) Ready() error {
	o.Status = OrderReady

	return nil
}

func (o *Order) Complete() error {
	o.Status = OrderCompleted

	return nil
}

func (o Order) GetTotal() float64 {
	var total float64

	for _, item := range o.Items {
		total += item.Price * float64(item.Quantity)
	}

	return total
}

func (o *Order) AddItem(product *Product, quantity int32) error {
	if v, ok := o.Items[product.ID]; ok {
		v.Quantity = quantity
		return nil
	}

	o.Items[product.ID] = &Item{
		ProductID:   product.ID,
		ProductName: product.Name,
		Price:       product.Price,
		Quantity:    quantity,
	}

	return nil
}
