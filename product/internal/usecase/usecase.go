package usecase

import (
	"context"
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

	GetProduct struct {
		ID string
	}

	UpdateProduct struct {
		ID          string
		Name        string
		Description string
		Price       float64
	}

	DeleteProduct struct {
		ID string
	}

	ServiceUsecase interface {
		CreateProduct(context.Context, CreateProduct) error
		GetProduct(context.Context, GetProduct) (*domain.Product, error)
		UpdateProduct(context.Context, UpdateProduct) error
		DeleteProduct(context.Context, DeleteProduct) error
	}

	serviceUsecase struct {
		product domain.ProductRepository
	}
)

var _ ServiceUsecase = (*serviceUsecase)(nil)

var UseCaseSet = wire.NewSet(NewService)

func NewService(repo domain.ProductRepository) ServiceUsecase {
	return &serviceUsecase{
		product: repo,
	}
}

func (a *serviceUsecase) CreateProduct(ctx context.Context, create CreateProduct) error {
	product, err := domain.CreateProduct(create.ID, create.Name, create.Description, create.Price)
	if err != nil {
		return err
	}

	return a.product.Save(ctx, product)
}

func (a *serviceUsecase) GetProduct(ctx context.Context, get GetProduct) (*domain.Product, error) {
	return a.product.Find(ctx, get.ID)
}

func (a *serviceUsecase) UpdateProduct(ctx context.Context, update UpdateProduct) error {
	product, err := domain.UpdateProduct(update.ID, update.Name, update.Description, update.Price)
	if err != nil {
		return err
	}

	return a.product.Update(ctx, product)
}

func (a *serviceUsecase) DeleteProduct(ctx context.Context, delete DeleteProduct) error {
	return a.product.Delete(ctx, delete.ID)
}
