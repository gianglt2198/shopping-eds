package domain

type OrderCreated struct {
	Order *Order
}

func (OrderCreated) EventName() string { return "orders.OrderCreated" }

type OrderCheckedout struct {
	Order *Order
}

func (OrderCheckedout) EventName() string { return "orders.OrderCheckedout" }

type OrderReadied struct {
	Order *Order
}

func (OrderReadied) EventName() string { return "orders.OrderReadied" }

type OrderCancelled struct {
	Order *Order
}

func (OrderCancelled) EventName() string { return "orders.OrderCancelled" }

type OrderAddedItem struct {
	Order *Order
}

func (OrderAddedItem) EventName() string { return "orders.OrderAddedItem" }

type OrderCompleted struct {
	Order *Order
}

func (OrderCompleted) EventName() string { return "orders.OrderCompleted" }
