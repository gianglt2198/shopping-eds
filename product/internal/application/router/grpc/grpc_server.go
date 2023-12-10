package grpc_router

import (
	"context"
	"shopping/product/internal/domain"
	"shopping/product/internal/logging"
	"shopping/product/internal/usecase"
	"shopping/product/internal/usecase/commands"
	"shopping/product/internal/usecase/queries"
	"shopping/product/pb"

	"github.com/google/uuid"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

type server struct {
	app usecase.ServiceUsecase
	pb.UnimplementedProductsServiceServer
}

var _ pb.ProductsServiceServer = (*server)(nil)

var ProductGRPCServerSet = wire.NewSet(RegisterServer)

func RegisterServer(app logging.Usecase, registrar *grpc.Server) error {
	pb.RegisterProductsServiceServer(registrar, server{app: app})
	return nil
}

func (s server) CreateProduct(ctx context.Context, request *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	id := uuid.New().String()
	err := s.app.CreateProduct(ctx, commands.CreateProduct{
		ID:          id,
		Name:        request.GetName(),
		Price:       request.GetPrice(),
		Description: request.GetDescription(),
	})
	return &pb.CreateProductResponse{Id: id}, err
}

func (s server) GetProduct(ctx context.Context, request *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	product, err := s.app.GetProduct(ctx, queries.GetProduct{
		ID: request.GetId(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.GetProductResponse{Product: s.productFromDomain(product)}, nil
}

func (s server) productFromDomain(product *domain.ManagementProduct) *pb.Product {
	return &pb.Product{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}
}

func (s server) DeleteProduct(ctx context.Context, request *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	err := s.app.DeleteProduct(ctx, commands.DeleteProduct{
		ID: request.GetId(),
	})
	return &pb.DeleteProductResponse{}, err
}

func (s server) IncreasePrice(ctx context.Context, request *pb.IncreasePriceRequest) (*pb.IncreasePriceResponse, error) {
	err := s.app.IncreasePriceProduct(ctx, commands.IncreasePrice{
		ID:    request.GetId(),
		Price: request.GetPrice(),
	})
	return &pb.IncreasePriceResponse{}, err
}

func (s server) DecreasePrice(ctx context.Context, request *pb.DecreasePriceRequest) (*pb.DecreasePriceResponse, error) {
	err := s.app.DecreasePriceProduct(ctx, commands.DecreasePrice{
		ID:    request.GetId(),
		Price: request.GetPrice(),
	})
	return &pb.DecreasePriceResponse{}, err
}
