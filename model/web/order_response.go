package web

import "time"

type ItemResponse struct {
	ItemId      uint      `json:"itemId"`
	ItemCode    string    `json:"itemCode"`
	Description string    `json:"description"`
	Quantity    uint      `json:"quantity"`
	OrderId     uint      `json:"orderId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type OrderResponse struct {
	OrderId      uint           `json:"orderId"`
	CustomerName string         `json:"customerName"`
	OrderedAt    time.Time      `json:"orderedAt"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	Items        []ItemResponse `json:"items"`
}
