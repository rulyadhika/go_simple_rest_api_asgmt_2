package web

import "time"

type itemResponse struct {
	ItemId      uint   `json:"item_id"`
	ItemCode    uint   `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"order_id"`
}

type OrderResponse struct {
	OrderId      uint           `json:"order_id"`
	CustomerName string         `json:"customer_name"`
	OrderedAt    time.Time      `json:"ordered_at"`
	Items        []itemResponse `json:"items"`
}
