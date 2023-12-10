package repo

import (
	"context"
	"database/sql"
	"fmt"
	"shopping/product/internal/domain"

	"github.com/pkg/errors"
)

type ManagementRepository struct {
	tableName string
	db        *sql.DB
}

var _ domain.ManagementRepository = (*ManagementRepository)(nil)

func NewManagementRepository(tableName string, db *sql.DB) ManagementRepository {
	return ManagementRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r ManagementRepository) CreateProduct(ctx context.Context, productID, name, description string, price float64) error {
	const query = `INSERT INTO %s (id, name, description, price) VALUES ($1, $2, $3, $4)`

	_, err := r.db.ExecContext(ctx, r.table(query), productID, name, description, price)

	return err
}
func (r ManagementRepository) UpdatePrice(ctx context.Context, productID string, delta float64) error {
	const query = `UPDATE %s SET price = price + $2 WHERE id = $1`

	_, err := r.db.ExecContext(ctx, r.table(query), productID, delta)

	return err
}
func (r ManagementRepository) DeleteProduct(ctx context.Context, productID string) error {
	const query = `DELETE FROM %s WHERE id = $1`

	_, err := r.db.ExecContext(ctx, r.table(query), productID)

	return err
}
func (r ManagementRepository) Find(ctx context.Context, productID string) (*domain.ManagementProduct, error) {
	const query = `SELECT name, description, price FROM %s WHERE id = $1 LIMIT 1`

	product := &domain.ManagementProduct{
		ID: productID,
	}

	err := r.db.QueryRowContext(ctx, r.table(query), productID).Scan(&product.Name, &product.Description, &product.Price)
	if err != nil {
		return nil, errors.Wrap(err, "scanning product")
	}

	return product, nil
}

func (r ManagementRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
