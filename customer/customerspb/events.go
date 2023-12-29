package customerspb

import (
	"shopping/internal/registry"
	"shopping/internal/registry/serdes"
)

const (
	CustomerAggregateChannel = "shopping.customers.events.Customer"

	CustomerRegisteredEvent = "customerspb.CustomerRegistered"
)

func Registrations(reg registry.Registry) error {
	serde := serdes.NewProtoSerde(reg)

	if err := serde.Register(&CustomerRegistered{}); err != nil {
		return err
	}

	return nil
}

func (*CustomerRegistered) Key() string { return CustomerRegisteredEvent }
