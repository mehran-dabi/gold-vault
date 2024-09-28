package ports

import (
	"context"

	"goldvault/trading-service/internal/core/domain/entity"
)

type (
	OrderPersistence interface {
		CreateOrder(ctx context.Context, order *entity.Order) error
		GetOrderByUserIDAndStatus(ctx context.Context, userID int64, status entity.OrderStatus) ([]*entity.Order, error)
		GetOrderByStatus(ctx context.Context, status entity.OrderStatus) ([]*entity.Order, error)
	}
)
