package repo

import (
	"context"
	"database/sql"
	"fmt"
	"shopping/internal/ddd"
	"shopping/product/internal/domain"
	"strings"

	"github.com/google/wire"
)

type ProductRepository struct {
	tableName string
	db        *sql.DB
}

var _ domain.ProductRepository = (*ProductRepository)(nil)

var RepositorySet = wire.NewSet(NewProductRepository)

func NewProductRepository(tableName string, db *sql.DB) domain.ProductRepository {
	return ProductRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r ProductRepository) Save(ctx context.Context, product *domain.Product) error {
	const query = "INSERT INTO %s (id, name, description, price) VALUES ($1, $2, $3, $4)"

	_, err := r.db.ExecContext(ctx, r.table(query), product.ID, product.Name, product.Description, product.Price)

	return err
}

func (r ProductRepository) Find(ctx context.Context, productID string) (*domain.Product, error) {
	const query = "SELECT name, description, price FROM %s WHERE id = $1 LIMIT 1"

	product := &domain.Product{
		AggregateBase: ddd.AggregateBase{
			ID: productID,
		},
	}

	err := r.db.QueryRowContext(ctx, r.table(query), productID).Scan(&product.Name, &product.Description, &product.Price)

	return product, err
}

func (r ProductRepository) Update(ctx context.Context, product *domain.Product) error {
	query := strings.Builder{}
	args := []interface{}{product.ID}
	query.WriteString("UPDATE %s SET ")
	i := 2
	if product.Name != "" {
		query.WriteString(fmt.Sprintf(" name = $%v ", i))
		args = append(args, product.Name)
		i++
	}
	if product.Description != "" {
		query.WriteString(fmt.Sprintf(", description = $%v ", i))
		args = append(args, product.Description)
		i++
	}
	if product.Price != 0 {
		query.WriteString(fmt.Sprintf(", price = $%v ", i))
		args = append(args, product.Price)
		i++
	}
	query.WriteString(" WHERE id = $1")

	_, err := r.db.ExecContext(ctx, r.table(query.String()), args...)

	return err
}

func (r ProductRepository) Delete(ctx context.Context, productID string) error {
	const query = "DELETE FROM %s WHERE id = $1"

	_, err := r.db.ExecContext(ctx, r.table(query), productID)

	return err
}

func (r ProductRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
