package domain

const (
	InvoiceCreatedEvent = "payments.InvoiceCreated"
	InvoicePaidEvent    = "payments.InvoicePaid"
)

type InvoiceCreated struct {
	ID      string
	OrderID string
}

type InvoicePaid struct {
	ID      string
	OrderID string
}
