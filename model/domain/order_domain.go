package domain

import "time"

type Order struct {
	OrderId      uint
	CustomerName string
	OrderedAt    time.Time
}
