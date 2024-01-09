package router_grpc

import (
	"context"
	"shopping/payment/internal/domain"
	"shopping/payment/internal/logging"
	"shopping/payment/internal/usecase"
	pb "shopping/payment/paymentspb"

	"github.com/google/uuid"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

type server struct {
	app usecase.ServiceUsecase
	pb.UnimplementedPaymentsServiceServer
}

var _ pb.PaymentsServiceServer = (*server)(nil)

var PaymentGRPCServerSet = wire.NewSet(RegisterServer)

func RegisterServer(app logging.Usecase, registrar *grpc.Server) error {
	pb.RegisterPaymentsServiceServer(registrar, server{app: app})
	return nil
}

func (s server) CreateInvoice(ctx context.Context, request *pb.CreateInvoiceRequest) (*pb.CreateInvoiceResponse, error) {
	id := uuid.New().String()
	err := s.app.CreateInvoice(ctx, usecase.CreateInvoice{
		ID:         id,
		OrderID:    request.GetOrderId(),
		CustomerID: request.GetCustomerId(),
		Amount:     request.GetAmount(),
	})
	return &pb.CreateInvoiceResponse{Id: id}, err
}

func (s server) GetInvoice(ctx context.Context, request *pb.GetInvoiceRequest) (*pb.GetInvoiceResponse, error) {
	invoice, err := s.app.GetInvoice(ctx, usecase.GetInvoice{
		ID: request.GetId(),
	})
	return &pb.GetInvoiceResponse{Invoice: s.invoiceFromDomain(invoice)}, err
}

func (s server) invoiceFromDomain(invoice *domain.Invoice) *pb.Invoice {
	return &pb.Invoice{
		Id:         invoice.ID,
		OrderId:    invoice.OrderID,
		CustomerId: invoice.CustomerID,
		Amount:     invoice.Amount,
		Status:     invoice.Status.String(),
	}
}

func (s server) PayInvoice(ctx context.Context, request *pb.PayInvoiceRequest) (*pb.PayInvoiceResponse, error) {
	err := s.app.PayInvoice(ctx, usecase.PayInvoice{
		ID: request.GetId(),
	})
	return &pb.PayInvoiceResponse{}, err
}

func (s server) CancelInvoice(ctx context.Context, request *pb.CancelInvoiceRequest) (*pb.CancelInvoiceResponse, error) {
	err := s.app.CancelInvoice(ctx, usecase.CancelInvoice{
		ID: request.GetId(),
	})
	return &pb.CancelInvoiceResponse{}, err
}
