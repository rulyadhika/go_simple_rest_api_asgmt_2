package service

import (
	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/fga_digitalent_assignment_2/model/web"
)

type OrderService interface {
	Create(ctx *gin.Context, request *web.OrderCreateRequest) (*web.OrderResponse, error)
	FindAll(ctx *gin.Context) (*[]web.OrderResponse, error)
	Update(ctx *gin.Context, request *web.OrderUpdateRequest) (*web.OrderResponse, error)
	Delete(ctx *gin.Context, orderId uint) error
}
