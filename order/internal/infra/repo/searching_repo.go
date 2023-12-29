package repo

import (
	"bytes"
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"shopping/order/internal/domain"
	"shopping/order/internal/models"
	"strings"

	"github.com/stackus/errors"
)

type SearchingRepository struct {
	tableName string
	db        *sql.DB
}

var _ domain.SearchingRepository = (*SearchingRepository)(nil)

func NewSearchingRepository(tableName string, db *sql.DB) *SearchingRepository {
	return &SearchingRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r *SearchingRepository) Add(ctx context.Context, order *models.Order) error {
	const query = `INSERT INTO %s (
		order_id, customer_id, customer_name, 
		status, created_at) VALUES (
		$1, $2, $3,
		$4, $5)`

	_, err := r.db.ExecContext(ctx, r.table(query),
		order.OrderID, order.CustomerID, order.CustomerName,
		order.Status, order.CreatedAt,
	)
	return err
}
func (r *SearchingRepository) UpdateStatus(ctx context.Context, orderID, status string) error {
	const query = `UPDATE %s SET status = $2 WHERE order_id = $1`

	_, err := r.db.ExecContext(ctx, r.table(query), orderID, status)
	return err
}

func (r *SearchingRepository) UpdateItem(ctx context.Context, orderID, productID string, quantity int, price float64) error {
	// const query = `UPDATE %s SET status = $2 WHERE order_id = $1`

	// _, err := r.db.ExecContext(ctx, r.table(query), orderID, status)

	const query = `SELECT items FROM %s WHERE order_id = $1`
	var (
		itemData []byte
	)
	err := r.db.QueryRowContext(ctx, r.table(query), orderID).Scan(&itemData)
	if err != nil {
		return err
	}
	var items []models.Item
	if len(items) > 0 {
		err = json.Unmarshal(itemData, &items)
		if err != nil {
			return err
		}
	}

	num := -1
	productIDs := make(IDArray, len(items))
	for i, item := range items {
		if item.ProductID == productID {
			num = i
		}
		productIDs = append(productIDs, productID)
	}

	if num == -1 {
		items = append(items, models.Item{
			ProductID: productID,
			Quantity:  quantity,
			Price:     price,
		})
		productIDs = append(productIDs, productID)
	} else {
		items[num].Quantity += quantity
		items[num].Price = price
	}

	itemsData, err := json.Marshal(items)
	if err != nil {
		return err
	}

	const state = `UPDATE %s SET items = $2, product_ids = $3 WHERE order_id = $1`

	_, err = r.db.ExecContext(ctx, r.table(state), orderID, itemsData, productIDs)

	return err
}

func (r *SearchingRepository) Search(ctx context.Context, search domain.SearchOrders) ([]*models.Order, error) {
	// TODO implement me
	panic("implement me")
}
func (r *SearchingRepository) Get(ctx context.Context, orderID string) (*models.Order, error) {
	const query = `SELECT customer_id, customer_name, items, status, created_at FROM %s WHERE order_id = $1`

	order := &models.Order{
		OrderID: orderID,
	}

	var itemData []byte
	err := r.db.QueryRowContext(ctx, r.table(query)).Scan(&order.CustomerID, &order.CustomerName, &itemData, &order.Status, &order.CreatedAt)
	if err != nil {
		return nil, err
	}

	var items []models.Item
	err = json.Unmarshal(itemData, &items)
	if err != nil {
		return nil, err
	}
	order.Items = items

	return order, nil
}

func (r SearchingRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}

type IDArray []string

func (a *IDArray) Scan(src any) error {
	var sep = []byte(",")

	var data []byte
	switch v := src.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	default:
		return errors.ErrInvalidArgument.Msgf("IDArray: unsupported type: %T", src)
	}

	ids := make([]string, bytes.Count(data, sep))
	for i, id := range bytes.Split(bytes.Trim(data, "{}"), sep) {
		ids[i] = string(id)
	}

	*a = ids

	return nil
}

func (a IDArray) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}
	if len(a) == 0 {
		return "{}", nil
	}
	// unsafe way to do this; assumption is all ids are UUIDs
	return fmt.Sprintf("{%s}", strings.Join(a, ",")), nil
}
