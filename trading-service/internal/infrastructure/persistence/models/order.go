package models

import (
	"time"

	"goldvault/trading-service/internal/core/domain/entity"
)

type Order struct {
	ID        int64
	UserID    int64
	AssetType string
	Quantity  float64
	Price     float64
	OrderType string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

const (
	OrderTableName   = "orders"
	OrderColumns     = "id, user_id, asset_type, quantity, price, order_type, status, created_at, updated_at"
	OrderColumnsNoID = "user_id, asset_type, quantity, price, order_type, status, created_at, updated_at"
)

func (o *Order) Scan(scanner Scanner) error {
	return scanner.Scan(
		&o.ID,
		&o.UserID,
		&o.AssetType,
		&o.Quantity,
		&o.Price,
		&o.OrderType,
		&o.Status,
		&o.CreatedAt,
		&o.UpdatedAt,
	)
}

func (o *Order) ToEntity() *entity.Order {
	return &entity.Order{
		ID:        o.ID,
		UserID:    o.UserID,
		AssetType: entity.AssetType(o.AssetType),
		Quantity:  o.Quantity,
		Price:     o.Price,
		OrderType: entity.OrderType(o.OrderType),
		Status:    entity.OrderStatus(o.Status),
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
	}
}

func ToOrderDB(order *entity.Order) *Order {
	return &Order{
		ID:        order.ID,
		UserID:    order.UserID,
		AssetType: order.AssetType.String(),
		Quantity:  order.Quantity,
		Price:     order.Price,
		OrderType: order.OrderType.String(),
		Status:    order.Status.String(),
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}
}
