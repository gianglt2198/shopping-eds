package pb

import (
	"shopping/internal/registry"
	"shopping/internal/registry/serdes"
)

const (
	ProductAggregateChannel = "shopping.products.events.Product"

	ProductCreatedEvent        = "productspb.ProductCreated"
	ProductPriceIncreasedEvent = "productspb.ProductPriceIncreased"
	ProductPriceDecreasedEvent = "productspb.ProductPriceDecreased"
	ProductDeletedEvent        = "productspb.ProductDeleted"
)

func Registration(reg registry.Registry) error {
	serde := serdes.NewProtoSerde(reg)

	if err := serde.Register(&ProductCreated{}); err != nil {
		return err
	}
	if err := serde.RegisterKey(ProductPriceIncreasedEvent, &ProductPriceChanged{}); err != nil {
		return err
	}
	if err := serde.RegisterKey(ProductPriceDecreasedEvent, &ProductPriceChanged{}); err != nil {
		return err
	}
	if err := serde.Register(&ProductDeleted{}); err != nil {
		return err
	}

	return nil
}

func (*ProductCreated) Key() string { return ProductCreatedEvent }
func (*ProductDeleted) Key() string { return ProductDeletedEvent }
