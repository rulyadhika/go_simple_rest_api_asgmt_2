package web

import "time"

// tambah validate
type itemUpdate struct {
	ItemId      uint   `json:"itemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
}

type OrderUpdateRequest struct {
	OrderId      uint         `json:"orderId"`
	CustomerName string       `json:"customerName"`
	OrderedAt    time.Time    `json:"orderedAt"`
	Items        []itemUpdate `json:"items"`
}
