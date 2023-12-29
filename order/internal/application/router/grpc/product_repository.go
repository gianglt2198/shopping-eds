package grpc_router

import (
	"context"
	"shopping/order/internal/domain"
	"shopping/order/internal/models"
	productspb "shopping/product/productspb"

	"google.golang.org/grpc"
)

type ProductRepository struct {
	client productspb.ProductsServiceClient
}

var _ domain.ProductRepository = (*ProductRepository)(nil)

func NewProductRepository(conn *grpc.ClientConn) domain.ProductRepository {
	return ProductRepository{client: productspb.NewProductsServiceClient(conn)}
}

func (r ProductRepository) GetProduct(ctx context.Context, productID string) (*models.Product, error) {
	product, err := r.client.GetProduct(ctx, &productspb.GetProductRequest{Id: productID})
	if err != nil {
		return nil, err
	}

	return r.productToDomain(product.Product), nil
}

func (r ProductRepository) productToDomain(product *productspb.Product) *models.Product {
	return &models.Product{
		ID:    product.Id,
		Name:  product.Name,
		Price: product.Price,
	}
}
