package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"

	"github.com/rulyadhika/fga_digitalent_assignment_2/exception"
	"github.com/rulyadhika/fga_digitalent_assignment_2/model/domain"
)

type OrderRepositoryImpl struct{}

func NewOrderRepositoryImpl() *OrderRepositoryImpl {
	return &OrderRepositoryImpl{}
}

func (o *OrderRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, order domain.Order, items []domain.Item) (domain.Order, []domain.Item, error) {
	orderQuery := `INSERT INTO "orders" (customer_name, ordered_at) VALUES ($1, $2) RETURNING order_id`

	err := tx.QueryRowContext(ctx, orderQuery, order.CustomerName, order.OrderedAt).Scan(&order.OrderId)
	if err != nil {
		return order, items, errors.New("something went wrong")
	}

	itemQuery := `INSERT INTO "items" (item_code, description, quantity, order_id) VALUES($1, $2, $3, $4) RETURNING item_id`
	itemQueryStatement, err := tx.PrepareContext(ctx, itemQuery)

	if err != nil {
		return order, items, errors.New("something went wrong")
	}

	defer itemQueryStatement.Close()

	for index, item := range items {
		var itemId int
		err := itemQueryStatement.QueryRowContext(ctx, item.ItemCode, item.Description, item.Quantity, order.OrderId).Scan(&itemId)

		items[index].ItemId = uint(itemId)
		items[index].OrderID = uint(order.OrderId)

		if err != nil {
			return order, items, errors.New("something went wrong")
		}
	}

	return order, items, nil
}

func (o *OrderRepositoryImpl) FindAll(ctx context.Context, db *sql.DB) ([]domain.Order, []domain.Item, error) {
	sqlQuery := `SELECT orders.order_id, orders.customer_name, orders.ordered_at, items.item_id, items.item_code, items.description, items.quantity, items.order_id 
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

		err := rows.Scan(&order.OrderId, &order.CustomerName, &order.OrderedAt, &item.ItemId, &item.ItemCode, &item.Description, &item.Quantity, &item.OrderID)
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
	orderQuery := `UPDATE orders SET customer_name=$1, ordered_at=$2 WHERE order_id=$3 RETURNING order_id`

	var affectedOrderRow int
	err := tx.QueryRowContext(ctx, orderQuery, order.CustomerName, order.OrderedAt, order.OrderId).Scan(&affectedOrderRow)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return order, items, exception.NewNotFoundError(fmt.Sprintf("data order dengan order_id:%v tidak ditemukan", order.OrderId))
		}

		return order, items, errors.New("something went wrong")
	}

	itemQuery := `UPDATE items SET item_code=$1, description=$2, quantity=$3 WHERE item_id=$4 AND order_id=$5 RETURNING item_id`
	itemQueryStatement, err := tx.PrepareContext(ctx, itemQuery)
	if err != nil {
		return order, items, errors.New("something went wrong")
	}

	for index, item := range items {
		var affectedItemRow int

		items[index].OrderID = order.OrderId
		err := itemQueryStatement.QueryRowContext(ctx, item.ItemCode, item.Description, item.Quantity, item.ItemId, order.OrderId).Scan(&affectedItemRow)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return order, items, exception.NewBadRequestError(fmt.Sprintf("data item dengan item_id:%v bukan merupakan data item milik order_id:%v", item.ItemId, order.OrderId))
			}

			return order, items, errors.New("something went wrong")
		}
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
