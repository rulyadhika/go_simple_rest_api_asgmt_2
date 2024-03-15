package repository

import (
	"context"
	"database/sql"

	"github.com/rulyadhika/fga_digitalent_assignment_2/helper"
	"github.com/rulyadhika/fga_digitalent_assignment_2/model/domain"
)

type OrderRepositoryImpl struct{}

func NewOrderRepositoryImpl() *OrderRepositoryImpl {
	return &OrderRepositoryImpl{}
}

func (o *OrderRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, order domain.Order, items []domain.Item) (domain.Order, []domain.Item) {
	orderQuery := `INSERT INTO "orders" (customer_name, ordered_at) VALUES ($1, $2) RETURNING order_id`

	err := tx.QueryRowContext(ctx, orderQuery, order.CustomerName, order.OrderedAt).Scan(&order.OrderId)
	helper.PanicIfErr(err)

	itemQuery := `INSERT INTO "items" (item_code, description, quantity, order_id) VALUES($1, $2, $3, $4) RETURNING item_id`
	itemQueryStatement, err := tx.PrepareContext(ctx, itemQuery)

	helper.PanicIfErr(err)
	defer itemQueryStatement.Close()

	for index, item := range items {
		var itemId int
		err := itemQueryStatement.QueryRowContext(ctx, item.ItemCode, item.Description, item.Quantity, order.OrderId).Scan(&itemId)

		items[index].ItemId = uint(itemId)
		items[index].OrderID = uint(order.OrderId)

		helper.PanicIfErr(err)
	}

	return order, items
}

// func (o *OrderRepositoryImpl) FindOne(ctx context.Context, tx *sql.Tx, orderId uint) (*domain.Order, *domain.Item) {
// 	panic("not implemented") // TODO: Implement
// }

// func (o *OrderRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, order *domain.Order, items *[]domain.Item) (*domain.Order, *[]domain.Item) {
// 	panic("not implemented") // TODO: Implement
// }

// func (o *OrderRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, orderId uint) {
// 	panic("not implemented") // TODO: Implement
// }
