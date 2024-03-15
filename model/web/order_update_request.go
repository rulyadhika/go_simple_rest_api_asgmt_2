package web

import "time"

// tambah validate
type itemUpdate struct {
	ItemId      uint   `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"order_id"`
}

type OrderUpdateRequest struct {
	OrderId      uint         `json:"order_id"`
	CustomerName string       `json:"customer_name"`
	OrderedAt    time.Time    `json:"ordered_at"`
	Items        []itemUpdate `json:"items"`
}
