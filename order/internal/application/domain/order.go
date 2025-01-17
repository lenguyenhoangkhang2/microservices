package domain

import "time"

type OrderItem struct {
	ProductCode string  `json:"product_code"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    int32   `json:"quantity"`
}

type Order struct {
	Status     string      `json:"status"`
	OrderItems []OrderItem `json:"order_items"`
	ID         int64       `json:"id"`
	CustomerId int64       `json:"customer_id"`
	CreatedAt  int64       `json:"created_at"`
}

func NewOrder(customerId int64, orderItems []OrderItem) Order {
	return Order{
		CreatedAt:  time.Now().Unix(),
		Status:     "Pending",
		CustomerId: customerId,
		OrderItems: orderItems,
	}
}
