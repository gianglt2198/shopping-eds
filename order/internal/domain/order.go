package domain

import (
	"shopping/internal/ddd"
	"shopping/internal/es"
	"shopping/order/internal/models"

	"github.com/stackus/errors"
)

const OrderAggregate = "ordering.Order"

var (
	ErrOrderHasNoItems          = errors.Wrap(errors.ErrBadRequest, "the order has no items")
	ErrOrderCannotBeCancelled   = errors.Wrap(errors.ErrBadRequest, "the order cannot be cancelled")
	ErrCustomerIDCannotBeBlank  = errors.Wrap(errors.ErrBadRequest, "the customer id cannot be blank")
	ErrPaymentIDCannotBeBlank   = errors.Wrap(errors.ErrBadRequest, "the payment id cannot be blank")
	ErrOrderCannotBeModified    = errors.Wrap(errors.ErrBadRequest, "the order cannot be modified")
	ErrQuantityCannotBeNegative = errors.Wrap(errors.ErrBadRequest, "the quantity cannot be nagative")
)

type (
	Order struct {
		es.Aggregate
		CustomerID string
		PaymentID  string
		Status     OrderStatus
		Items      []*Item
	}

	Item struct {
		ProductID   string
		ProductName string
		Price       float64
		Quantity    int32
	}
)

var _ interface {
	es.Snapshotter
	es.EventApplier
} = (*Order)(nil)

func NewOrder(id string) *Order {
	return &Order{
		Aggregate: es.NewAggregate(id, OrderAggregate),
	}
}

func (Order) Key() string { return OrderAggregate }

func (o Order) IsPending() bool {
	return o.Status == OrderStatusPending
}

func (o Order) IsCheckOut() bool {
	return o.Status == OrderStatusCheckedout
}

func (o Order) IsCanceled() bool {
	return o.Status != OrderStatusCheckedout && o.Status != OrderStatusReady
}

func CreateOrder(id, customerID string) (*Order, error) {
	if customerID == "" {
		return nil, ErrCustomerIDCannotBeBlank
	}

	order := NewOrder(id)

	order.AddEvent(OrderCreatedEvent, &OrderCreated{
		CustomerID: customerID,
	})

	return order, nil
}

func (o *Order) Cancel() error {
	if !o.IsPending() && o.IsCanceled() {
		return ErrOrderCannotBeCancelled
	}

	o.AddEvent(OrderCanceledEvent, &OrderCancelled{
		PaymentID: o.PaymentID,
	})

	return nil
}

func (o *Order) Checkout() error {
	if !o.IsPending() {
		return ErrOrderCannotBeModified
	}

	o.AddEvent(OrderCheckedOutEvent, &OrderCheckedout{
		CustomerID: o.CustomerID,
		Total:      o.GetTotal(),
	})

	return nil
}

func (o *Order) Ready(paymentID string) error {
	if !o.IsPending() && !o.IsCheckOut() {
		return ErrOrderCannotBeModified
	}

	if paymentID == "" {
		return ErrPaymentIDCannotBeBlank
	}

	o.PaymentID = paymentID

	o.AddEvent(OrderReadiedEvent, &OrderReadied{
		PaymenID: paymentID,
	})

	return nil
}

func (o *Order) Complete() error {
	o.AddEvent(OrderCompletedEvent, &OrderCompleted{})
	return nil
}

func (o Order) GetTotal() float64 {
	var total float64

	for _, item := range o.Items {
		total += item.Price * float64(item.Quantity)
	}

	return total
}

func (o *Order) AddItem(product *models.Product, quantity int32) error {
	if !o.IsPending() {
		return ErrOrderCannotBeModified
	}

	if quantity < 0 {
		return ErrQuantityCannotBeNegative
	}

	o.AddEvent(OrderAddedItemEvent, &OrderAddedItem{
		Item: Item{
			ProductID:   product.ID,
			ProductName: product.Name,
			Price:       product.Price,
			Quantity:    quantity,
		},
	})

	return nil
}

func (o *Order) ApplyEvent(event ddd.Event) error {
	switch payload := event.Payload().(type) {
	case *OrderCreated:
		o.CustomerID = payload.CustomerID
		o.Status = OrderStatusPending
	case *OrderCheckedout:
		o.Status = OrderStatusCheckedout
	case *OrderReadied:
		o.PaymentID = payload.PaymenID
		o.Status = OrderStatusReady
	case *OrderCancelled:
		o.Status = OrderStatusCancelled
	case *OrderCompleted:
		o.Status = OrderStatusCompleted
	case *OrderAddedItem:
		isFound := false
		item := &Item{
			ProductID:   payload.Item.ProductID,
			ProductName: payload.Item.ProductName,
			Price:       payload.Item.Price,
			Quantity:    payload.Item.Quantity,
		}
		for _, i := range o.Items {
			if i.ProductID == payload.Item.ProductID {
				i.Quantity += payload.Item.Quantity
				isFound = true
				break
			}
		}

		if !isFound {
			o.Items = append(o.Items, item)
		}
		o.Status = OrderStatusPending
	default:
		return errors.ErrInternal.Msgf("%T received the event %s with unexpected payload %T", o, event.EventName(), payload)
	}
	return nil
}

func (o *Order) ApplySnapshot(snapshot es.Snapshot) error {
	switch ss := snapshot.(type) {
	case *OrderV1:
		o.CustomerID = ss.CustomerID
		o.PaymentID = ss.PaymentID
		o.Items = ss.Items
		o.Status = ss.Status
	default:
		return errors.ErrInternal.Msgf("%T received the unexpected snapshot %T", o, snapshot)
	}
	return nil
}

func (o Order) ToSnapshot() es.Snapshot {
	return OrderV1{
		CustomerID: o.CustomerID,
		PaymentID:  o.PaymentID,
		Items:      o.Items,
		Status:     o.Status,
	}
}
