package web

import "time"

// tambah validate
type itemCreate struct {
	ItemCode    uint   `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"order_id"`
}

type OrderCreateRequest struct {
	CustomerName string       `json:"customer_name"`
	OrderedAt    time.Time    `json:"ordered_at"`
	Items        []itemCreate `json:"items"`
}
