package helper

import (
	"github.com/rulyadhika/go_simple_rest_api_asgmt_2/model/domain"
	"github.com/rulyadhika/go_simple_rest_api_asgmt_2/model/web"
)

func toItemResponse(items *[]domain.Item) *[]web.ItemResponse {
	itemsResponse := []web.ItemResponse{}

	for _, item := range *items {
		itemResponse := web.ItemResponse{
			ItemId:      item.ItemId,
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
			OrderId:     item.OrderId,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		}

		itemsResponse = append(itemsResponse, itemResponse)
	}

	return &itemsResponse
}

func ToOrderReponse(order *domain.Order, items *[]domain.Item) *web.OrderResponse {
	orderResponse := &web.OrderResponse{
		OrderId:      order.OrderId,
		CustomerName: order.CustomerName,
		OrderedAt:    order.OrderedAt,
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
		Items:        *toItemResponse(items),
	}

	return orderResponse
}

func ToOrdersReponse(orders *[]domain.Order, items *[]domain.Item) *[]web.OrderResponse {
	ordersResponse := []web.OrderResponse{}

	for _, order := range *orders {
		orderItems := []domain.Item{}

		for _, item := range *items {
			if item.OrderId == order.OrderId {
				orderItems = append(orderItems, item)
			}
		}

		orderResponse := web.OrderResponse{
			OrderId:      order.OrderId,
			CustomerName: order.CustomerName,
			OrderedAt:    order.OrderedAt,
			CreatedAt:    order.CreatedAt,
			UpdatedAt:    order.UpdatedAt,
			Items:        *toItemResponse(&orderItems),
		}

		ordersResponse = append(ordersResponse, orderResponse)
	}

	return &ordersResponse
}
