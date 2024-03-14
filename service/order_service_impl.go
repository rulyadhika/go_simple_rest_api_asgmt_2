package service

import (
	"context"
	"database/sql"

	"github.com/rulyadhika/fga_digitalent_assignment_2/model/web"
	"github.com/rulyadhika/fga_digitalent_assignment_2/repository"
)

// tambah validasi
type OrderServiceImpl struct {
	DB              *sql.DB
	OrderRepository *repository.OrderRepository
}

func NewOrderServiceImpl(db *sql.DB, orderRepository *repository.OrderRepository) *OrderServiceImpl {
	return &OrderServiceImpl{
		DB:              db,
		OrderRepository: orderRepository,
	}
}

func (o *OrderServiceImpl) Create(ctx context.Context, request *web.OrderCreateRequest) *web.OrderResponse {
	panic("not implemented") // TODO: Implement
}

func (o *OrderServiceImpl) FindOne(ctx context.Context, orderId uint) *web.OrderResponse {
	panic("not implemented") // TODO: Implement
}

func (o *OrderServiceImpl) Update(ctx context.Context, request *web.OrderUpdateRequest) *web.OrderResponse {
	panic("not implemented") // TODO: Implement
}

func (o *OrderServiceImpl) Delete(ctx context.Context, orderId uint) {
	panic("not implemented") // TODO: Implement
}
