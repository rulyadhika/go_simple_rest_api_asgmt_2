package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/go_simple_rest_api_asgmt_2/handler"
)

func NewOrderRouter(r *gin.Engine, orderHandler handler.OrderHandler) {
	r.POST("/api/orders", orderHandler.Create)
	r.GET("/api/orders", orderHandler.FindAll)
	r.PUT("/api/orders/:orderId", orderHandler.Update)
	r.DELETE("/api/orders/:orderId", orderHandler.Delete)
}
