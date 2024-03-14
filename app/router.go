package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rulyadhika/fga_digitalent_assignment_2/handler"
)

func NewOrderRouter(orderHandler handler.OrderHandler) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/orders", orderHandler.Create)
	router.GET("/api/orders/:orderId", orderHandler.FindOne)
	router.PUT("/api/orders/:orderId", orderHandler.Update)
	router.DELETE("/api/orders/:orderId", orderHandler.Delete)

	return router
}
