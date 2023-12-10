package domain

const (
	ProductCreatedEvent         = "products.ProductCreated"
	ProductDeletedEvent         = "products.ProductDeleted"
	ProductInscreasedPriceEvent = "products.ProductInscreasedPrice"
	ProductDescreasedPriceEvent = "products.ProductDescreasedPrice"
)

type ProductCreated struct {
	Name        string
	Description string
	Price       float64
}

func (ProductCreated) Key() string { return ProductCreatedEvent }

type ProductPriceChanged struct {
	Delta float64
}

type ProductDeleted struct{}

func (ProductDeleted) Key() string { return ProductDeletedEvent }
