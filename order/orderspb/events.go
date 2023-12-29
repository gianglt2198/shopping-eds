package orderspb

import (
	"shopping/internal/registry"
	"shopping/internal/registry/serdes"
)

const (
	OrderAggregateChannel = "shopping.ordering.events.Order"

	OrderCreatedEvent    = "orderspb.OrderCreated"
	OrderAddedItemEvent  = "orderspb.OrderAddedItem"
	OrderCheckedOutEvent = "orderspb.OrderCheckedOut"
	OrderReadiedEvent    = "orderspb.OrderReadied"
	OrderCanceledEvent   = "orderspb.OrderCanceled"
	OrderCompletedEvent  = "orderspb.OrderCompleted"
)

func Registrations(reg registry.Registry) (err error) {
	serde := serdes.NewProtoSerde(reg)

	if err = serde.Register(&OrderCreated{}); err != nil {
		return err
	}
	if err = serde.Register(&OrderAddedItem{}); err != nil {
		return err
	}
	if err = serde.Register(&OrderCheckedOut{}); err != nil {
		return err
	}
	if err = serde.Register(&OrderReadied{}); err != nil {
		return err
	}
	if err = serde.Register(&OrderCanceled{}); err != nil {
		return err
	}
	if err = serde.Register(&OrderCompleted{}); err != nil {
		return err
	}
	return nil
}

func (*OrderCreated) Key() string    { return OrderCreatedEvent }
func (*OrderAddedItem) Key() string  { return OrderAddedItemEvent }
func (*OrderCheckedOut) Key() string { return OrderCheckedOutEvent }
func (*OrderReadied) Key() string    { return OrderReadiedEvent }
func (*OrderCanceled) Key() string   { return OrderCanceledEvent }
func (*OrderCompleted) Key() string  { return OrderCompletedEvent }
