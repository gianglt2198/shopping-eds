package domain

type OrderV1 struct {
	CustomerID string
	PaymentID  string
	Status     OrderStatus
	Items      []*Item
}

func (OrderV1) SnapshotName() string { return "ordering.OrderV1" }
