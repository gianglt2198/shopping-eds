package grpc_router

import (
	"context"
	"shopping/order/internal/domain"
	paymentspb "shopping/payment/pb"

	"github.com/google/wire"
	"google.golang.org/grpc"
)

type PaymentRepository struct {
	client paymentspb.PaymentsServiceClient
}

var _ domain.PaymentRepository = (*PaymentRepository)(nil)

var PaymentClientSet = wire.NewSet(NewPaymentRepository)

func NewPaymentRepository(conn *grpc.ClientConn) domain.PaymentRepository {
	return PaymentRepository{client: paymentspb.NewPaymentsServiceClient(conn)}
}

func (r PaymentRepository) GetInvoice(ctx context.Context, paymentID string) error {
	_, err := r.client.GetInvoice(ctx, &paymentspb.GetInvoiceRequest{Id: paymentID})
	return err
}

func (r PaymentRepository) CancelInvoice(ctx context.Context, paymentID string) error {
	_, err := r.client.CancelInvoice(ctx, &paymentspb.CancelInvoiceRequest{Id: paymentID})
	return err
}

func (r PaymentRepository) CreateInvoice(ctx context.Context, orderID, customerID string, amount float64) (string, error) {
	resp, err := r.client.CreateInvoice(ctx, &paymentspb.CreateInvoiceRequest{OrderId: orderID, CustomerId: customerID, Amount: amount})
	if err != nil {
		return "", err
	}

	return resp.Id, nil
}
