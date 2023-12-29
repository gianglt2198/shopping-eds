package grpc_router

import (
	"context"
	"shopping/product/internal/domain"
	"shopping/product/internal/logging"
	"shopping/product/internal/usecase"
	"shopping/product/internal/usecase/commands"
	"shopping/product/internal/usecase/queries"
	"shopping/product/productspb"

	"github.com/google/uuid"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

type server struct {
	app usecase.ServiceUsecase
	productspb.UnimplementedProductsServiceServer
}

var _ productspb.ProductsServiceServer = (*server)(nil)

var ProductGRPCServerSet = wire.NewSet(RegisterServer)

func RegisterServer(app logging.Usecase, registrar *grpc.Server) error {
	productspb.RegisterProductsServiceServer(registrar, server{app: app})
	return nil
}

func (s server) CreateProduct(ctx context.Context, request *productspb.CreateProductRequest) (*productspb.CreateProductResponse, error) {
	id := uuid.New().String()
	err := s.app.CreateProduct(ctx, commands.CreateProduct{
		ID:          id,
		Name:        request.GetName(),
		Price:       request.GetPrice(),
		Description: request.GetDescription(),
	})
	return &productspb.CreateProductResponse{Id: id}, err
}

func (s server) GetProduct(ctx context.Context, request *productspb.GetProductRequest) (*productspb.GetProductResponse, error) {
	product, err := s.app.GetProduct(ctx, queries.GetProduct{
		ID: request.GetId(),
	})
	if err != nil {
		return nil, err
	}
	return &productspb.GetProductResponse{Product: s.productFromDomain(product)}, nil
}

func (s server) productFromDomain(product *domain.ManagementProduct) *productspb.Product {
	return &productspb.Product{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}
}

func (s server) DeleteProduct(ctx context.Context, request *productspb.DeleteProductRequest) (*productspb.DeleteProductResponse, error) {
	err := s.app.DeleteProduct(ctx, commands.DeleteProduct{
		ID: request.GetId(),
	})
	return &productspb.DeleteProductResponse{}, err
}

func (s server) IncreasePrice(ctx context.Context, request *productspb.IncreasePriceRequest) (*productspb.IncreasePriceResponse, error) {
	err := s.app.IncreasePriceProduct(ctx, commands.IncreasePrice{
		ID:    request.GetId(),
		Price: request.GetPrice(),
	})
	return &productspb.IncreasePriceResponse{}, err
}

func (s server) DecreasePrice(ctx context.Context, request *productspb.DecreasePriceRequest) (*productspb.DecreasePriceResponse, error) {
	err := s.app.DecreasePriceProduct(ctx, commands.DecreasePrice{
		ID:    request.GetId(),
		Price: request.GetPrice(),
	})
	return &productspb.DecreasePriceResponse{}, err
}
