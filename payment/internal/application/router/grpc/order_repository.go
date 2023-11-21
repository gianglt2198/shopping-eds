package router

import (
	"context"

	orderspb "shopping/order/pb"
	"shopping/payment/internal/domain"

	"github.com/google/wire"
	"google.golang.org/grpc"
)

type OrderRepository struct {
	client orderspb.OrderingServiceClient
}

var _ domain.OrderRepository = (*OrderRepository)(nil)

var OrderClientSet = wire.NewSet(NewOrderRepository)

func NewOrderRepository(conn *grpc.ClientConn) domain.OrderRepository {
	return OrderRepository{client: orderspb.NewOrderingServiceClient(conn)}
}

func (r OrderRepository) CompleteOrder(ctx context.Context, orderID string) error {
	_, err := r.client.CompleteOrder(ctx, &orderspb.CompleteOrderRequest{Id: orderID})
	return err
}
