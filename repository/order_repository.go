package repository

import (
	"context"
	"database/sql"

	"github.com/rulyadhika/go_simple_rest_api_asgmt_2/model/domain"
)

type OrderRepository interface {
	Create(ctx context.Context, tx *sql.Tx, order domain.Order, items []domain.Item) (domain.Order, []domain.Item, error)
	FindAll(ctx context.Context, db *sql.DB) ([]domain.Order, []domain.Item, error)
	Update(ctx context.Context, tx *sql.Tx, order domain.Order, items []domain.Item) (domain.Order, []domain.Item, error)
	Delete(ctx context.Context, db *sql.DB, orderId uint) error
}
