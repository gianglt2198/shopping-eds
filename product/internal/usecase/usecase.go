package usecase

import (
	"context"
	"shopping/internal/ddd"
	"shopping/product/internal/domain"
	"shopping/product/internal/usecase/commands"
	"shopping/product/internal/usecase/queries"

	"github.com/google/wire"
)

type (
	ServiceUsecase interface {
		Commands
		Queries
	}

	Commands interface {
		CreateProduct(context.Context, commands.CreateProduct) error
		UpdateProduct(context.Context, commands.UpdateProduct) error
		DeleteProduct(context.Context, commands.DeleteProduct) error
	}

	Queries interface {
		GetProduct(context.Context, queries.GetProduct) (*domain.Product, error)
	}

	usecaseCommands struct {
		commands.CreateProductHandler
		commands.UpdateProductHandler
		commands.DeleteProductHandler
	}

	usecaseQueries struct {
		queries.GetProductHandler
	}

	serviceUsecase struct {
		usecaseCommands
		usecaseQueries
	}
)

var _ ServiceUsecase = (*serviceUsecase)(nil)

var UseCaseSet = wire.NewSet(NewService)

func NewService(repo domain.ProductRepository, domainPublisher ddd.EventPublisher) ServiceUsecase {
	return &serviceUsecase{
		usecaseCommands: usecaseCommands{
			CreateProductHandler: commands.NewCreateProductHandler(repo, domainPublisher),
			UpdateProductHandler: commands.NewUpdateProductHandler(repo, domainPublisher),
			DeleteProductHandler: commands.NewDeleteProductHandler(repo, domainPublisher),
		},
		usecaseQueries: usecaseQueries{
			GetProductHandler: queries.NewGetProductHandler(repo),
		},
	}
}
