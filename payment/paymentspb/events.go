package paymentspb

import (
	"shopping/internal/registry"
	"shopping/internal/registry/serdes"
)

const (
	PaymentAggregateChannel = "shopping.payments.events.Invoice"

	InvoiceCreatedEvent = "paymentspb.InvoiceCreated"
	InvoicePaidEvent    = "paymentspb.InvoicePaid"
)

func Registration(reg registry.Registry) error {
	serde := serdes.NewProtoSerde(reg)

	if err := serde.Register(&InvoiceCreated{}); err != nil {
		return err
	}
	if err := serde.Register(&InvoicePaid{}); err != nil {
		return err
	}

	return nil
}

func (*InvoiceCreated) Key() string { return InvoiceCreatedEvent }
func (*InvoicePaid) Key() string    { return InvoicePaidEvent }
