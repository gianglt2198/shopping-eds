package usecase

import (
	"context"
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
		DeleteProduct(context.Context, commands.DeleteProduct) error
		IncreasePriceProduct(context.Context, commands.IncreasePrice) error
		DecreasePriceProduct(context.Context, commands.DecreasePrice) error
	}

	Queries interface {
		GetProduct(context.Context, queries.GetProduct) (*domain.ManagementProduct, error)
	}

	usecaseCommands struct {
		commands.CreateProductHandler
		commands.IncreaseProductPriceHandler
		commands.DecreaseProductPriceHandler
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

func NewService(products domain.ProductRepository, management domain.ManagementRepository) ServiceUsecase {
	return &serviceUsecase{
		usecaseCommands: usecaseCommands{
			CreateProductHandler:        commands.NewCreateProductHandler(products),
			DeleteProductHandler:        commands.NewDeleteProductHandler(products),
			IncreaseProductPriceHandler: commands.NewIncreaseProductPriceHandler(products),
			DecreaseProductPriceHandler: commands.NewDecreaseProductPriceHandler(products),
		},
		usecaseQueries: usecaseQueries{
			GetProductHandler: queries.NewGetProductHandler(management),
		},
	}
}
