package controllers

import "time"

type AddOrderRequest struct {
	CustomerName string         `json:"customer_name"`
	OrderAt      string         `json:"order_at"`
	Items        []ItemsRequest `json:"items"`
}

type ItemsRequest struct {
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
}

type OrderResponse struct {
	ID           uint            `json:"id"`
	CustomerName string          `json:"customer_name"`
	Items        []ItemsResponse `json:"items"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}

type ItemsResponse struct {
	ID          uint      `json:"item_id"`
	ItemCode    string    `json:"item_code"`
	OrderId     uint      `json:"order_id"`
	Description string    `json:"description"`
	Quantity    uint      `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
