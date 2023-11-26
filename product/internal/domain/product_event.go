package domain

type ProductCreated struct {
	Product *Product
}

func (ProductCreated) EventName() string { return "products.ProductCreated" }

type ProductUpdated struct {
	Product *Product
}

func (ProductUpdated) EventName() string { return "products.ProductUpdated" }

type ProductDeleted struct {
	Product *Product
}

func (ProductDeleted) EventName() string { return "products.ProductDeleted" }
