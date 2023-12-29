package router_grpc

import (
	"context"
	"shopping/customer/customerspb"
	"shopping/customer/internal/domain"
	"shopping/customer/internal/logging"
	"shopping/customer/internal/usecase"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	app usecase.ServiceUsecase
	customerspb.UnimplementedCustomersServiceServer
}

var _ customerspb.CustomersServiceServer = (*server)(nil)

func RegisterServer(app logging.Usecase, registrar *grpc.Server) error {
	customerspb.RegisterCustomersServiceServer(registrar, server{app: app})
	return nil
}

func (s server) RegisterCustomer(ctx context.Context, request *customerspb.RegisterCustomerRequest) (*customerspb.RegisterCustomerResponse, error) {
	id := uuid.New().String()
	err := s.app.RegisterCustomer(ctx, usecase.RegisterCustomer{
		ID:        id,
		Name:      request.GetName(),
		SmsNumber: request.GetSmsNumber(),
		Email:     request.GetEmail(),
	})
	return &customerspb.RegisterCustomerResponse{Id: id}, err
}

func (s server) GetCustomer(ctx context.Context, request *customerspb.GetCustomerRequest) (*customerspb.GetCustomerResponse, error) {
	customer, err := s.app.GetCustomer(ctx, usecase.GetCustomer{
		ID: request.GetId(),
	})
	return &customerspb.GetCustomerResponse{Customer: s.customerFromDomain(customer)}, err
}

func (s server) customerFromDomain(customer *domain.Customer) *customerspb.Customer {
	return &customerspb.Customer{
		Id:        customer.ID(),
		Name:      customer.Name,
		SmsNumber: customer.SmsNumber,
		Email:     customer.Email,
		Active:    customer.Active,
	}
}
