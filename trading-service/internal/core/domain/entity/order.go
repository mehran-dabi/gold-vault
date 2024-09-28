package entity

import "time"

type Order struct {
	ID        int64
	UserID    int64
	AssetType AssetType
	Quantity  float64
	Price     float64
	OrderType OrderType
	Status    OrderStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderType string

const (
	OrderTypeBuy  OrderType = "Buy"
	OrderTypeSell OrderType = "Sell"
)

func (o *OrderType) String() string {
	return string(*o)
}

func (o *OrderType) IsValid() bool {
	switch *o {
	case OrderTypeBuy, OrderTypeSell:
		return true
	}
	return false
}

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "Pending"
	OrderStatusCompleted OrderStatus = "Completed"
	OrderStatusCancelled OrderStatus = "Cancelled"
	OrderStatusFailed    OrderStatus = "Failed"
)

func (o *OrderStatus) String() string {
	return string(*o)
}

func (o *OrderStatus) IsValid() bool {
	switch *o {
	case OrderStatusPending, OrderStatusCompleted, OrderStatusCancelled, OrderStatusFailed:
		return true
	}
	return false
}
