package domain

type OrderStatus string

const (
	OrderUnknown          OrderStatus = ""
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusCheckedout OrderStatus = "checkedout"
	OrderStatusReady      OrderStatus = "ready"
	OrderStatusCompleted  OrderStatus = "completed"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

func (s OrderStatus) String() string {
	switch s {
	case OrderStatusPending, OrderStatusCheckedout, OrderStatusReady, OrderStatusCompleted, OrderStatusCancelled:
		return string(s)
	default:
		return ""
	}
}

func ToOrderStatus(status string) OrderStatus {
	switch status {
	case OrderStatusPending.String():
		return OrderStatusPending
	case OrderStatusCheckedout.String():
		return OrderStatusCheckedout
	case OrderStatusReady.String():
		return OrderStatusReady
	case OrderStatusCancelled.String():
		return OrderStatusCancelled
	case OrderStatusCompleted.String():
		return OrderStatusCompleted
	default:
		return OrderUnknown
	}
}
