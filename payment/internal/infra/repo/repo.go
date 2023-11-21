package repo

import (
	"context"
	"database/sql"
	"fmt"
	"shopping/payment/internal/domain"

	"github.com/google/wire"
)

type PaymentRepository struct {
	tableName string
	db        *sql.DB
}

var _ domain.PaymentRepository = (*PaymentRepository)(nil)

var RepositorySet = wire.NewSet(NewPaymentRepository)

func NewPaymentRepository(tableName string, db *sql.DB) domain.PaymentRepository {
	return PaymentRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r PaymentRepository) Save(ctx context.Context, invoice *domain.Invoice) error {
	const query = "INSERT INTO %s (id, order_id, customer_id, amount, status) VALUES ($1, $2, $3, $4, $5)"

	_, err := r.db.ExecContext(ctx, r.table(query), invoice.ID, invoice.OrderID, invoice.CustomerID, invoice.Amount, invoice.Status)

	return err
}

func (r PaymentRepository) Find(ctx context.Context, invoiceID string) (*domain.Invoice, error) {
	const query = "SELECT order_id, customer_id, amount, status FROM %s WHERE id = $1 LIMIT 1"

	invoice := &domain.Invoice{
		ID: invoiceID,
	}

	var status string
	err := r.db.QueryRowContext(ctx, r.table(query), invoiceID).Scan(&invoice.OrderID, &invoice.CustomerID, &invoice.Amount, &status)
	invoice.Status = domain.ToPaymentStatus(status)
	return invoice, err
}

func (r PaymentRepository) Update(ctx context.Context, id, status string) error {
	const query = "UPDATE %s SET status = $1 WHERE id = $2"

	_, err := r.db.ExecContext(ctx, r.table(query), status, id)

	return err
}

func (r PaymentRepository) Delete(ctx context.Context, invoiceID string) error {
	const query = "DELETE FROM %s WHERE id = $1"

	_, err := r.db.ExecContext(ctx, r.table(query), invoiceID)

	return err
}

func (r PaymentRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
