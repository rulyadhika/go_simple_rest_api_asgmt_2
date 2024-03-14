package repository

import (
	"context"
	"database/sql"

	"github.com/rulyadhika/fga_digitalent_assignment_2/model/domain"
)

type OrderRepositoryImpl struct{}

func NewOrderRepositoryImpl() *OrderRepositoryImpl {
	return &OrderRepositoryImpl{}
}

func (o *OrderRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, order *domain.Order, items *[]domain.Item) (*domain.Order, *domain.Item) {
	panic("not implemented") // TODO: Implement
}

func (o *OrderRepositoryImpl) FindOne(ctx context.Context, tx *sql.Tx, orderId uint) (*domain.Order, *domain.Item) {
	panic("not implemented") // TODO: Implement
}

func (o *OrderRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, order *domain.Order, items *[]domain.Item) (*domain.Order, *domain.Item) {
	panic("not implemented") // TODO: Implement
}

func (o *OrderRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, orderId uint) {
	panic("not implemented") // TODO: Implement
}
