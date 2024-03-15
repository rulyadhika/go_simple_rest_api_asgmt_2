package helper

import (
	"github.com/rulyadhika/fga_digitalent_assignment_2/model/domain"
	"github.com/rulyadhika/fga_digitalent_assignment_2/model/web"
)

func toItemResponse(items *[]domain.Item) *[]web.ItemResponse {
	itemsResponse := []web.ItemResponse{}

	for _, item := range *items {
		itemResponse := web.ItemResponse{
			ItemId:      item.ItemId,
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
			OrderID:     item.OrderID,
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
		Items:        *toItemResponse(items),
	}

	return orderResponse
}
