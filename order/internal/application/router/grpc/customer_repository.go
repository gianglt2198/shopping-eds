package grpc_router

import (
	"context"
	customerspb "shopping/customer/customerspb"
	"shopping/order/internal/domain"
	"shopping/order/internal/models"

	"google.golang.org/grpc"
)

type CustomerRepository struct {
	client customerspb.CustomersServiceClient
}

var _ domain.CustomerRepository = (*CustomerRepository)(nil)

func NewCustomerRepository(conn *grpc.ClientConn) domain.CustomerRepository {
	return CustomerRepository{client: customerspb.NewCustomersServiceClient(conn)}
}

func (r CustomerRepository) GetCustomer(ctx context.Context, customerID string) (*models.Customer, error) {
	c, err := r.client.GetCustomer(ctx, &customerspb.GetCustomerRequest{Id: customerID})
	if err != nil {
		return nil, err
	}
	return &models.Customer{
		ID:   c.Customer.GetId(),
		Name: c.Customer.GetName(),
	}, nil
}
