package service

import (
	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/fga_digitalent_assignment_2/model/web"
)

type OrderService interface {
	Create(ctx *gin.Context, request *web.OrderCreateRequest) *web.OrderResponse
	FindAll(ctx *gin.Context) *[]web.OrderResponse
	// Update(ctx *gin.Context, request *web.OrderUpdateRequest) *web.OrderResponse
	// Delete(ctx *gin.Context, orderId uint)
}
