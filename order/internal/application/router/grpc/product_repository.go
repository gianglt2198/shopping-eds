package grpc_router

import (
	"context"
	"shopping/order/internal/domain"
	productspb "shopping/product/pb"

	"github.com/google/wire"
	"google.golang.org/grpc"
)

type ProductRepository struct {
	client productspb.ProductsServiceClient
}

var _ domain.ProductRepository = (*ProductRepository)(nil)

var ProductClientSet = wire.NewSet(NewProductRepository)

func NewProductRepository(conn *grpc.ClientConn) domain.ProductRepository {
	return ProductRepository{client: productspb.NewProductsServiceClient(conn)}
}

func (r ProductRepository) GetProduct(ctx context.Context, productID string) (*domain.Product, error) {
	product, err := r.client.GetProduct(ctx, &productspb.GetProductRequest{Id: productID})
	if err != nil {
		return nil, err
	}

	return r.productToDomain(product.Product), nil
}

func (r ProductRepository) productToDomain(product *productspb.Product) *domain.Product {
	return &domain.Product{
		ID:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}
}
