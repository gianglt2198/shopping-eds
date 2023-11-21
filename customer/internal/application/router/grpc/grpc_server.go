package router

import (
	"context"
	"shopping/customer/internal/domain"
	"shopping/customer/internal/logging"
	"shopping/customer/internal/usecase"
	"shopping/customer/pb"

	"github.com/google/uuid"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

type server struct {
	app usecase.ServiceUsecase
	pb.UnimplementedCustomersServiceServer
}

var _ pb.CustomersServiceServer = (*server)(nil)

var CustomerGRPCServerSet = wire.NewSet(RegisterServer)

func RegisterServer(app logging.Usecase, registrar *grpc.Server) error {
	pb.RegisterCustomersServiceServer(registrar, server{app: app})
	return nil
}

func (s server) RegisterCustomer(ctx context.Context, request *pb.RegisterCustomerRequest) (*pb.RegisterCustomerResponse, error) {
	id := uuid.New().String()
	err := s.app.RegisterCustomer(ctx, usecase.RegisterCustomer{
		ID:        id,
		Name:      request.GetName(),
		SmsNumber: request.GetSmsNumber(),
		Email:     request.GetEmail(),
	})
	return &pb.RegisterCustomerResponse{Id: id}, err
}

func (s server) GetCustomer(ctx context.Context, request *pb.GetCustomerRequest) (*pb.GetCustomerResponse, error) {
	customer, err := s.app.GetCustomer(ctx, usecase.GetCustomer{
		ID: request.GetId(),
	})
	return &pb.GetCustomerResponse{Customer: s.customerFromDomain(customer)}, err
}

func (s server) customerFromDomain(customer *domain.Customer) *pb.Customer {
	return &pb.Customer{
		Id:        customer.ID,
		Name:      customer.Name,
		SmsNumber: customer.SmsNumber,
		Email:     customer.Email,
		Active:    customer.Active,
	}
}
