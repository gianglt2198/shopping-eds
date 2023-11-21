package domain

type PaymentStatus string

const (
	PaymentUnknown   PaymentStatus = ""
	PaymentPending   PaymentStatus = "pending"
	PaymentPaid      PaymentStatus = "paid"
	PaymentCancelled PaymentStatus = "cancelled"
)

func (s PaymentStatus) String() string {
	switch s {
	case PaymentPending, PaymentPaid, PaymentCancelled:
		return string(s)
	default:
		return ""
	}
}

func ToPaymentStatus(status string) PaymentStatus {
	switch status {
	case PaymentPending.String():
		return PaymentPending
	case PaymentPaid.String():
		return PaymentPaid
	case PaymentCancelled.String():
		return PaymentCancelled
	default:
		return PaymentUnknown
	}
}
