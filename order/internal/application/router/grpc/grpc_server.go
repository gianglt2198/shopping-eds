package grpc_router

import (
	"context"
	"shopping/order/internal/domain"
	"shopping/order/internal/logging"
	"shopping/order/internal/usecase"
	"shopping/order/internal/usecase/commands"
	"shopping/order/internal/usecase/queries"
	"shopping/order/pb"

	"github.com/google/uuid"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

type server struct {
	app usecase.ServiceUsecase
	pb.UnimplementedOrderingServiceServer
}

var _ pb.OrderingServiceServer = (*server)(nil)

var OrderGRPCServerSet = wire.NewSet(RegisterServer)

func RegisterServer(app logging.Usecase, registrar *grpc.Server) error {
	pb.RegisterOrderingServiceServer(registrar, server{app: app})
	return nil
}

func (s server) CreateOrder(ctx context.Context, request *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	data := commands.CreateOrder{
		ID:         uuid.New().String(),
		CustomerID: request.GetCustomerId(),
	}

	err := s.app.CreateOrder(ctx, data)

	return &pb.CreateOrderResponse{Id: data.ID}, err
}

func (s server) GetOrder(ctx context.Context, request *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	order, err := s.app.GetOrder(ctx, queries.GetOrder{
		ID: request.GetId(),
	})

	return &pb.GetOrderResponse{Order: s.orderFromDomain(order)}, err
}

func (s server) orderFromDomain(order *domain.Order) *pb.Order {
	if order == nil {
		return nil
	}
	items := make([]*pb.Item, 0)
	for _, item := range order.Items {
		items = append(items, &pb.Item{
			ProductId:   item.ProductID,
			ProductName: item.ProductName,
			Price:       item.Price,
			Quantity:    item.Quantity,
		})
	}
	return &pb.Order{
		Id:         order.ID(),
		CustomerId: order.CustomerID,
		Items:      items,
		Status:     order.Status.String(),
		PaymentId:  order.PaymentID,
	}
}

func (s server) AddItem(ctx context.Context, request *pb.AddItemRequest) (*pb.AddItemResponse, error) {
	err := s.app.AddItem(ctx, commands.AddItem{
		OrderID:  request.GetOrderId(),
		ItemID:   request.GetProductId(),
		Quantity: request.GetQuantity(),
	})
	return &pb.AddItemResponse{}, err
}

func (s server) CancelOrder(ctx context.Context, request *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	err := s.app.CancelOrder(ctx, commands.CancelOrder{
		ID: request.GetId(),
	})
	return &pb.CancelOrderResponse{}, err
}

func (s server) CheckoutOrder(ctx context.Context, request *pb.CheckoutOrderRequest) (*pb.CheckoutOrderResponse, error) {
	err := s.app.CheckoutOrder(ctx, commands.CheckoutOrder{
		ID: request.GetId(),
	})
	return &pb.CheckoutOrderResponse{}, err
}

func (s server) ReadyOrder(ctx context.Context, request *pb.ReadyOrderRequest) (*pb.ReadyOrderResponse, error) {
	err := s.app.ReadyOrder(ctx, commands.ReadyOrder{
		ID:        request.GetId(),
		PaymentID: request.GetPaymentId(),
	})
	return &pb.ReadyOrderResponse{}, err
}

func (s server) CompleteOrder(ctx context.Context, request *pb.CompleteOrderRequest) (*pb.CompleteOrderResponse, error) {
	err := s.app.CompleteOrder(ctx, commands.CompleteOrder{
		ID: request.GetId(),
	})
	return &pb.CompleteOrderResponse{}, err
}
