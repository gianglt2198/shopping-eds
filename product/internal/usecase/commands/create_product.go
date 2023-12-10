package commands

import (
	"context"
	"shopping/product/internal/domain"

	"github.com/google/wire"
	"github.com/pkg/errors"
)

type (
	CreateProduct struct {
		ID          string
		Name        string
		Description string
		Price       float64
	}

	CreateProductHandler struct {
		products domain.ProductRepository
	}
)

var CreateProductUseCaseSet = wire.NewSet(NewCreateProductHandler)

func NewCreateProductHandler(products domain.ProductRepository) CreateProductHandler {
	return CreateProductHandler{
		products: products,
	}
}

func (h CreateProductHandler) CreateProduct(ctx context.Context, cmd CreateProduct) error {
	product, err := domain.CreateProduct(cmd.ID, cmd.Name, cmd.Description, cmd.Price)
	if err != nil {
		return errors.Wrap(err, "error adding product")
	}

	return errors.Wrap(h.products.Save(ctx, product), "error adding product")
}
