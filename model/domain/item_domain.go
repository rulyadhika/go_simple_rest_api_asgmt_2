package domain

import "time"

type Item struct {
	ItemId      uint
	ItemCode    string
	Description string
	Quantity    uint
	OrderId     uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
