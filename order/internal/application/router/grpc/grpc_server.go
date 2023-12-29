package grpc_router

import (
	"context"
	"shopping/order/internal/domain"
	"shopping/order/internal/logging"
	"shopping/order/internal/usecase"
	"shopping/order/internal/usecase/commands"
	"shopping/order/internal/usecase/queries"
	"shopping/order/orderspb"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	app usecase.ServiceUsecase
	orderspb.UnimplementedOrderingServiceServer
}

var _ orderspb.OrderingServiceServer = (*server)(nil)

func RegisterServer(app logging.Usecase, registrar *grpc.Server) error {
	orderspb.RegisterOrderingServiceServer(registrar, server{app: app})
	return nil
}

func (s server) CreateOrder(ctx context.Context, request *orderspb.CreateOrderRequest) (*orderspb.CreateOrderResponse, error) {
	data := commands.CreateOrder{
		ID:         uuid.New().String(),
		CustomerID: request.GetCustomerId(),
	}

	err := s.app.CreateOrder(ctx, data)

	return &orderspb.CreateOrderResponse{Id: data.ID}, err
}

func (s server) GetOrder(ctx context.Context, request *orderspb.GetOrderRequest) (*orderspb.GetOrderResponse, error) {
	order, err := s.app.GetOrder(ctx, queries.GetOrder{
		ID: request.GetId(),
	})

	return &orderspb.GetOrderResponse{Order: s.orderFromDomain(order)}, err
}

func (s server) orderFromDomain(order *domain.Order) *orderspb.Order {
	if order == nil {
		return nil
	}
	items := make([]*orderspb.Item, 0)
	for _, item := range order.Items {
		items = append(items, &orderspb.Item{
			ProductId:   item.ProductID,
			ProductName: item.ProductName,
			Price:       item.Price,
			Quantity:    item.Quantity,
		})
	}
	return &orderspb.Order{
		Id:         order.ID(),
		CustomerId: order.CustomerID,
		Items:      items,
		Status:     order.Status.String(),
		PaymentId:  order.PaymentID,
	}
}

func (s server) AddItem(ctx context.Context, request *orderspb.AddItemRequest) (*orderspb.AddItemResponse, error) {
	err := s.app.AddItem(ctx, commands.AddItem{
		OrderID:  request.GetOrderId(),
		ItemID:   request.GetProductId(),
		Quantity: request.GetQuantity(),
	})
	return &orderspb.AddItemResponse{}, err
}

func (s server) CancelOrder(ctx context.Context, request *orderspb.CancelOrderRequest) (*orderspb.CancelOrderResponse, error) {
	err := s.app.CancelOrder(ctx, commands.CancelOrder{
		ID: request.GetId(),
	})
	return &orderspb.CancelOrderResponse{}, err
}

func (s server) CheckoutOrder(ctx context.Context, request *orderspb.CheckoutOrderRequest) (*orderspb.CheckoutOrderResponse, error) {
	err := s.app.CheckoutOrder(ctx, commands.CheckoutOrder{
		ID: request.GetId(),
	})
	return &orderspb.CheckoutOrderResponse{}, err
}

func (s server) ReadyOrder(ctx context.Context, request *orderspb.ReadyOrderRequest) (*orderspb.ReadyOrderResponse, error) {
	err := s.app.ReadyOrder(ctx, commands.ReadyOrder{
		ID:        request.GetId(),
		PaymentID: request.GetPaymentId(),
	})
	return &orderspb.ReadyOrderResponse{}, err
}

func (s server) CompleteOrder(ctx context.Context, request *orderspb.CompleteOrderRequest) (*orderspb.CompleteOrderResponse, error) {
	err := s.app.CompleteOrder(ctx, commands.CompleteOrder{
		ID: request.GetId(),
	})
	return &orderspb.CompleteOrderResponse{}, err
}
