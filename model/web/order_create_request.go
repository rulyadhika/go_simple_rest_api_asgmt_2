package web

import "time"

// tambah validate
type itemCreate struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
}

type OrderCreateRequest struct {
	CustomerName string       `json:"customerName"`
	OrderedAt    time.Time    `json:"orderedAt"`
	Items        []itemCreate `json:"items"`
}
