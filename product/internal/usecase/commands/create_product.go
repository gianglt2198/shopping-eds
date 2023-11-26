package commands

import (
	"context"
	"shopping/internal/ddd"
	"shopping/product/internal/domain"

	"github.com/google/wire"
)

type (
	CreateProduct struct {
		ID          string
		Name        string
		Description string
		Price       float64
	}

	CreateProductHandler struct {
		products        domain.ProductRepository
		domainPublisher ddd.EventPublisher
	}
)

var CreateProductUseCaseSet = wire.NewSet(NewCreateProductHandler)

func NewCreateProductHandler(products domain.ProductRepository, domainPublisher ddd.EventPublisher) CreateProductHandler {
	return CreateProductHandler{
		products:        products,
		domainPublisher: domainPublisher,
	}
}

func (h CreateProductHandler) CreateProduct(ctx context.Context, cmd CreateProduct) error {
	product, err := domain.CreateProduct(cmd.ID, cmd.Name, cmd.Description, cmd.Price)
	if err != nil {
		return err
	}

	if err = h.products.Save(ctx, product); err != nil {
		return err
	}

	if err = h.domainPublisher.Publish(ctx, product.GetEvents()...); err != nil {
		return err
	}

	return nil
}
