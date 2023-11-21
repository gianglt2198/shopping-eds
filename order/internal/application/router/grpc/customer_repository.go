package router

import (
	"context"
	customerspb "shopping/customer/pb"
	"shopping/order/internal/domain"

	"github.com/google/wire"
	"google.golang.org/grpc"
)

type CustomerRepository struct {
	client customerspb.CustomersServiceClient
}

var _ domain.CustomerRepository = (*CustomerRepository)(nil)

var CustomerClientSet = wire.NewSet(NewCustomerRepository)

func NewCustomerRepository(conn *grpc.ClientConn) domain.CustomerRepository {
	return CustomerRepository{client: customerspb.NewCustomersServiceClient(conn)}
}

func (r CustomerRepository) GetCustomer(ctx context.Context, customerID string) error {
	_, err := r.client.GetCustomer(ctx, &customerspb.GetCustomerRequest{Id: customerID})
	return err
}
