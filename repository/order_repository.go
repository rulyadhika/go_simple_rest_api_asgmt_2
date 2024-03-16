package repository

import (
	"context"
	"database/sql"

	"github.com/rulyadhika/fga_digitalent_assignment_2/model/domain"
)

type OrderRepository interface {
	Create(ctx context.Context, tx *sql.Tx, order domain.Order, items []domain.Item) (domain.Order, []domain.Item)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Order, []domain.Item, error)
	Update(ctx context.Context, tx *sql.Tx, order domain.Order, items []domain.Item) (domain.Order, []domain.Item)
	Delete(ctx context.Context, tx *sql.Tx, orderId uint)
}
