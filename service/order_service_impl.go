package service

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/fga_digitalent_assignment_2/helper"
	"github.com/rulyadhika/fga_digitalent_assignment_2/model/domain"
	"github.com/rulyadhika/fga_digitalent_assignment_2/model/web"
	"github.com/rulyadhika/fga_digitalent_assignment_2/repository"
)

// tambah validasi
type OrderServiceImpl struct {
	DB              *sql.DB
	OrderRepository repository.OrderRepository
}

func NewOrderServiceImpl(db *sql.DB, orderRepository repository.OrderRepository) *OrderServiceImpl {
	return &OrderServiceImpl{
		DB:              db,
		OrderRepository: orderRepository,
	}
}
func (o *OrderServiceImpl) Create(ctx *gin.Context, request *web.OrderCreateRequest) *web.OrderResponse {
	tx, err := o.DB.Begin()

	helper.PanicIfErr(err)
	defer helper.CommitOrRollbackTx(tx)

	order := domain.Order{
		CustomerName: request.CustomerName,
		OrderedAt:    request.OrderedAt,
	}

	items := []domain.Item{}

	for _, item := range request.Items {
		item := domain.Item{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
		}

		items = append(items, item)
	}

	orderResult, itemsResult := o.OrderRepository.Create(ctx, tx, order, items)

	return helper.ToOrderReponse(&orderResult, &itemsResult)
}

func (o *OrderServiceImpl) FindAll(ctx *gin.Context) *[]web.OrderResponse {
	tx, err := o.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollbackTx(tx)

	order, items, err := o.OrderRepository.FindAll(ctx, tx)

	helper.PanicIfErr(err)

	return helper.ToOrdersReponse(&order, &items)
}

// func (o *OrderServiceImpl) Update(ctx *gin.Context, request *web.OrderUpdateRequest) *web.OrderResponse {
// 	panic("not implemented") // TODO: Implement
// }

// func (o *OrderServiceImpl) Delete(ctx *gin.Context, orderId uint) {
// 	panic("not implemented") // TODO: Implement
// }
