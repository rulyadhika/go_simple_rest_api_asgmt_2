package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/go_simple_rest_api_asgmt_2/exception"
	"github.com/rulyadhika/go_simple_rest_api_asgmt_2/model/web"
	"github.com/rulyadhika/go_simple_rest_api_asgmt_2/service"
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
	if err != nil {
		ctx.Error(exception.NewUnprocessableEntityError("invalid json request body"))
		return
	}

	result, err := o.OrderService.Create(ctx, orderCreateRequest)

	if err != nil {
		ctx.Error(err)
		return
	}

	response := web.WebResponse{
		Status: http.StatusText(http.StatusCreated),
		Code:   http.StatusCreated,
		Data:   result,
	}

	ctx.JSON(http.StatusCreated, response)
}

func (o *OrderHandlerImpl) FindAll(ctx *gin.Context) {
	result, err := o.OrderService.FindAll(ctx)

	if err != nil {
		ctx.Error(err)
		return
	}

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

	if err != nil {
		ctx.Error(exception.NewUnprocessableEntityError("invalid syntax for orderId param. Should be numeric type"))
		return
	}

	orderUpdateRequest := &web.OrderUpdateRequest{}

	err = ctx.ShouldBindJSON(orderUpdateRequest)
	if err != nil {
		ctx.Error(exception.NewUnprocessableEntityError("invalid json request body"))
		return
	}

	orderUpdateRequest.OrderId = uint(id)

	result, err := o.OrderService.Update(ctx, orderUpdateRequest)

	if err != nil {
		ctx.Error(err)
		return
	}

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

	if err != nil {
		ctx.Error(exception.NewUnprocessableEntityError("invalid syntax for orderId param. Should be numeric type"))
		return
	}

	err = o.OrderService.Delete(ctx, uint(id))

	if err != nil {
		ctx.Error(err)
		return
	}

	response := &web.WebResponse{
		Status: http.StatusText(http.StatusOK),
		Code:   http.StatusOK,
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, response)
}
