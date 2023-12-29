package models

import (
	"time"
)

type Order struct {
	OrderID      string
	CustomerID   string
	CustomerName string
	Items        []Item
	Total        float64
	Status       string
	CreatedAt    time.Time
}

type Item struct {
	ProductID   string
	ProductName string
	Price       float64
	Quantity    int
}
