package domain

import (
	"shopping/internal/ddd"

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
		ddd.AggregateBase
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
		AggregateBase: ddd.AggregateBase{
			ID: id,
		},
		CustomerID: customerID,
		PaymentID:  paymentID,
		Items:      orderItems,
		Status:     OrderStatusPending,
	}

	order.AddEvents(&OrderCreated{
		Order: order,
	})

	return order, nil
}

func (o *Order) Cancel() error {
	if o.Status != OrderStatusPending && o.Status != OrderStatusCheckedout && o.Status != OrderStatusReady {
		return ErrOrderCannotBeCancelled
	}

	o.Status = OrderStatusCancelled

	o.AddEvents(&OrderCancelled{
		Order: o,
	})

	return nil
}

func (o *Order) Checkout() error {
	o.Status = OrderStatusCheckedout

	o.AddEvents(&OrderCheckedout{
		Order: o,
	})
	return nil
}

func (o *Order) Ready(paymentID string) error {
	o.Status = OrderStatusReady
	o.PaymentID = paymentID

	o.AddEvents(&OrderReadied{
		Order: o,
	})
	return nil
}

func (o *Order) Complete() error {
	o.Status = OrderStatusCompleted
	o.AddEvents(&OrderCompleted{
		Order: o,
	})
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

	o.AddEvents(&OrderAddedItem{
		Order: o,
	})

	return nil
}
