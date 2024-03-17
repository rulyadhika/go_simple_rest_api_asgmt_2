package service

import (
	"database/sql"
	"errors"

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
func (o *OrderServiceImpl) Create(ctx *gin.Context, request *web.OrderCreateRequest) (*web.OrderResponse, error) {
	tx, err := o.DB.Begin()

	if err != nil {
		return &web.OrderResponse{}, errors.New("something went wrong")
	}

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

	orderResult, itemsResult, err := o.OrderRepository.Create(ctx, tx, order, items)

	if err != nil {
		errRollback := tx.Rollback()

		if errRollback != nil {
			err = errors.New("something went wrong")
		}
	} else {
		errCommit := tx.Commit()

		if errCommit != nil {
			err = errors.New("something went wrong")
		}
	}

	return helper.ToOrderReponse(&orderResult, &itemsResult), err
}

func (o *OrderServiceImpl) FindAll(ctx *gin.Context) (*[]web.OrderResponse, error) {
	order, items, err := o.OrderRepository.FindAll(ctx, o.DB)

	return helper.ToOrdersReponse(&order, &items), err
}

func (o *OrderServiceImpl) Update(ctx *gin.Context, request *web.OrderUpdateRequest) (*web.OrderResponse, error) {
	tx, err := o.DB.Begin()
	if err != nil {
		return &web.OrderResponse{}, errors.New("something went wrong")
	}

	order := domain.Order{
		OrderId:      request.OrderId,
		CustomerName: request.CustomerName,
		OrderedAt:    request.OrderedAt,
	}

	items := []domain.Item{}

	for _, item := range request.Items {
		item := domain.Item{
			ItemId:      item.ItemId,
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
		}

		items = append(items, item)
	}

	orderResult, itemResult, err := o.OrderRepository.Update(ctx, tx, order, items)

	if err != nil {
		errRollback := tx.Rollback()

		if errRollback != nil {
			err = errors.New("something went wrong")
		}
	} else {
		errCommit := tx.Commit()

		if errCommit != nil {
			err = errors.New("something went wrong")
		}
	}

	return helper.ToOrderReponse(&orderResult, &itemResult), err
}

func (o *OrderServiceImpl) Delete(ctx *gin.Context, orderId uint) error {
	err := o.OrderRepository.Delete(ctx, o.DB, orderId)

	return err
}
