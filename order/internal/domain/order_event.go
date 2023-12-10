package domain

const (
	OrderCreatedEvent    = "ordering.OrderCreated"
	OrderCanceledEvent   = "ordering.OrderCanceled"
	OrderReadiedEvent    = "ordering.OrderReadied"
	OrderCompletedEvent  = "ordering.OrderCompleted"
	OrderAddedItemEvent  = "ordering.OrderAddedItem"
	OrderCheckedOutEvent = "ordering.OrderCheckedOut"
)

type OrderCreated struct {
	CustomerID string
	Items      []Item
}

func (OrderCreated) Key() string { return OrderCreatedEvent }

type OrderCheckedout struct {
	CustomerID string
	Total      float64
}

func (OrderCheckedout) Key() string { return OrderCheckedOutEvent }

type OrderReadied struct {
	PaymenID string
}

func (OrderReadied) Key() string { return OrderReadiedEvent }

type OrderCancelled struct {
	PaymentID string
}

func (OrderCancelled) Key() string { return OrderCanceledEvent }

type OrderAddedItem struct {
	Item Item
}

func (OrderAddedItem) Key() string { return OrderAddedItemEvent }

type OrderCompleted struct{}

func (OrderCompleted) Key() string { return OrderCompletedEvent }
