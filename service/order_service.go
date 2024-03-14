package service

import (
	"context"

	"github.com/rulyadhika/fga_digitalent_assignment_2/model/web"
)

type OrderService interface {
	Create(ctx context.Context, request *web.OrderCreateRequest) *web.OrderResponse
	FindOne(ctx context.Context, orderId uint) *web.OrderResponse
	Update(ctx context.Context, request *web.OrderUpdateRequest) *web.OrderResponse
	Delete(ctx context.Context, orderId uint)
}
