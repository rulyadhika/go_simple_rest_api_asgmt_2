package handler

import (
	"net/http"
	"strconv"

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

func (o *OrderHandlerImpl) FindAll(ctx *gin.Context) {
	result := o.OrderService.FindAll(ctx)

	response := &web.WebResponse{
		Status: http.StatusText(http.StatusOK),
		Code:   http.StatusOK,
		Data:   result,
	}

	ctx.JSON(http.StatusOK, response)
}

func (o *OrderHandlerImpl) Update(ctx *gin.Context) {
	orderId := ctx.Param("orderId")
	id, err := strconv.Atoi(orderId)

	helper.PanicIfErr(err)

	orderUpdateRequest := &web.OrderUpdateRequest{}

	ctx.ShouldBindJSON(orderUpdateRequest)

	orderUpdateRequest.OrderId = uint(id)

	result := o.OrderService.Update(ctx, orderUpdateRequest)

	response := &web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   result,
	}

	ctx.JSON(http.StatusOK, response)
}

func (o *OrderHandlerImpl) Delete(ctx *gin.Context) {
	orderId := ctx.Param("orderId")

	id, err := strconv.Atoi(orderId)

	helper.PanicIfErr(err)

	o.OrderService.Delete(ctx, uint(id))

	response := &web.WebResponse{
		Status: http.StatusText(http.StatusOK),
		Code:   http.StatusOK,
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, response)
}
