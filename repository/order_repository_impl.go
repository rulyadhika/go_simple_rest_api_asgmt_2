package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/rulyadhika/fga_digitalent_assignment_2/exception"
	"github.com/rulyadhika/fga_digitalent_assignment_2/model/domain"
)

type OrderRepositoryImpl struct{}

func NewOrderRepositoryImpl() *OrderRepositoryImpl {
	return &OrderRepositoryImpl{}
}

func (o *OrderRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, order domain.Order, items []domain.Item) (domain.Order, []domain.Item, error) {
	orderQuery := `INSERT INTO "orders" (customer_name, ordered_at) VALUES ($1, $2) RETURNING order_id, created_at, updated_at`

	err := tx.QueryRowContext(ctx, orderQuery, order.CustomerName, order.OrderedAt).Scan(&order.OrderId, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		return order, items, errors.New("something went wrong")
	}

	itemQuery := `INSERT INTO "items" (item_code, description, quantity, order_id) VALUES($1, $2, $3, $4) RETURNING item_id, created_at, updated_at`
	itemQueryStatement, err := tx.PrepareContext(ctx, itemQuery)

	if err != nil {
		return order, items, errors.New("something went wrong")
	}

	defer itemQueryStatement.Close()

	for index, item := range items {
		var itemId int
		var createdAt, updatedAt time.Time
		err := itemQueryStatement.QueryRowContext(ctx, item.ItemCode, item.Description, item.Quantity, order.OrderId).Scan(&itemId, &createdAt, &updatedAt)

		items[index].ItemId = uint(itemId)
		items[index].OrderId = uint(order.OrderId)
		items[index].CreatedAt = createdAt
		items[index].UpdatedAt = updatedAt

		if err != nil {
			return order, items, errors.New("something went wrong")
		}
	}

	return order, items, nil
}

func (o *OrderRepositoryImpl) FindAll(ctx context.Context, db *sql.DB) ([]domain.Order, []domain.Item, error) {
	sqlQuery := `SELECT orders.order_id, orders.customer_name, orders.ordered_at, orders.created_at, orders.updated_at, items.item_id,
	items.item_code, items.description, items.quantity, items.order_id, items.created_at, items.updated_at 
	FROM orders LEFT JOIN items ON orders.order_id=items.order_id`

	orders := []domain.Order{}
	items := []domain.Item{}

	rows, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		return orders, items, errors.New("something went wrong")
	}

	defer rows.Close()

	wg := &sync.WaitGroup{}
	mx := &sync.Mutex{}

	for rows.Next() {
		order := domain.Order{}
		item := domain.Item{}

		err := rows.Scan(&order.OrderId, &order.CustomerName, &order.OrderedAt, &order.CreatedAt, &order.UpdatedAt, &item.ItemId, &item.ItemCode, &item.Description, &item.Quantity, &item.OrderId, &item.CreatedAt, &item.UpdatedAt)
		items = append(items, item)

		wg.Add(1)

		go func(order *domain.Order) {
			defer wg.Done()

			mx.Lock()
			exist := false
			for _, o := range orders {
				if order.OrderId == o.OrderId {
					exist = true
					break
				}
			}

			if !exist {
				orders = append(orders, *order)
			}
			mx.Unlock()

		}(&order)

		if err != nil {
			return orders, items, errors.New("something went wrong")
		}
	}

	wg.Wait()

	return orders, items, nil
}

func (o *OrderRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, order domain.Order, items []domain.Item) (domain.Order, []domain.Item, error) {
	orderQuery := `UPDATE orders SET customer_name=$1, ordered_at=$2 WHERE order_id=$3 RETURNING order_id, created_at, updated_at`

	var affectedOrderRow int
	var createdAt, updatedAt time.Time

	err := tx.QueryRowContext(ctx, orderQuery, order.CustomerName, order.OrderedAt, order.OrderId).Scan(&affectedOrderRow, &createdAt, &updatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return order, items, exception.NewNotFoundError(fmt.Sprintf("data order dengan order_id:%v tidak ditemukan", order.OrderId))
		}

		return order, items, errors.New("something went wrong")
	}

	order.CreatedAt = createdAt
	order.UpdatedAt = updatedAt

	itemQuery := `UPDATE items SET item_code=$1, description=$2, quantity=$3 WHERE item_id=$4 AND order_id=$5 RETURNING item_id, created_at, updated_at`
	itemQueryStatement, err := tx.PrepareContext(ctx, itemQuery)
	if err != nil {
		return order, items, errors.New("something went wrong")
	}

	for index, item := range items {
		var affectedItemRow int
		var createdAt, updatedAt time.Time

		err := itemQueryStatement.QueryRowContext(ctx, item.ItemCode, item.Description, item.Quantity, item.ItemId, item.OrderId).Scan(&affectedItemRow, &createdAt, &updatedAt)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return order, items, exception.NewBadRequestError(fmt.Sprintf("data item dengan item_id:%v bukan merupakan data item milik order_id:%v", item.ItemId, order.OrderId))
			}

			return order, items, errors.New("something went wrong")
		}

		items[index].CreatedAt = createdAt
		items[index].UpdatedAt = updatedAt
	}

	return order, items, nil
}

func (o *OrderRepositoryImpl) Delete(ctx context.Context, db *sql.DB, orderId uint) error {
	sqlQuery := `DELETE FROM orders WHERE order_id=$1 RETURNING order_id`

	var afectedId int

	err := db.QueryRowContext(ctx, sqlQuery, orderId).Scan(&afectedId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return exception.NewNotFoundError(fmt.Sprintf("data order dengan order_id:%v tidak ditemukan", orderId))
		}

		return errors.New("something went wrong")
	}

	return nil
}
