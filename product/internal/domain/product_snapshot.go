package domain

type ProductV1 struct {
	Name        string
	Description string
	Price       float64
}

func (ProductV1) SnapshotName() string { return "products.ProductV1" }
