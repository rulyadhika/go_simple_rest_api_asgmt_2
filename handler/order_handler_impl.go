package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/fga_digitalent_assignment_2/helper"
	"github.com/rulyadhika/fga_digitalent_assignment_2/model/web"
	"github.com/rulyadhika/fga_digitalent_assignment_2/service"
)

type OrderHandlerImpl struct {
	OrderService service.OrderService
}

func NewOrderHandlerImpl(orderService service.OrderService) *OrderHandlerImpl {
	return &OrderHandlerImpl{
		OrderService: orderService,
	}
}

func (o *OrderHandlerImpl) Create(ctx *gin.Context) {
	orderCreateRequest := &web.OrderCreateRequest{}

	err := ctx.ShouldBindJSON(orderCreateRequest)
	helper.PanicIfErr(err)

	result := o.OrderService.Create(ctx, orderCreateRequest)

	response := web.WebResponse{
		Status: http.StatusText(http.StatusCreated),
		Code:   http.StatusCreated,
		Data:   result,
	}

	ctx.JSON(http.StatusCreated, response)
}

// func (o *OrderHandlerImpl) FindOne(ctx *gin.Context) {
// 	panic("not implemented") // TODO: Implement
// }

// func (o *OrderHandlerImpl) Update(ctx *gin.Context) {
// 	panic("not implemented") // TODO: Implement
// }

// func (o *OrderHandlerImpl) Delete(ctx *gin.Context) {
// 	panic("not implemented") // TODO: Implement
// }
