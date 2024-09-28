package persistence

import (
	"context"
	"database/sql"

	"goldvault/trading-service/internal/core/application/ports"
	"goldvault/trading-service/internal/core/domain/entity"
	"goldvault/trading-service/internal/infrastructure/persistence/models"
	"goldvault/trading-service/internal/infrastructure/persistence/queries"
)

type OrderPersistence struct {
	db *sql.DB
}

func NewOrderPersistence(db *sql.DB) ports.OrderPersistence {
	return &OrderPersistence{db: db}
}

func (o *OrderPersistence) CreateOrder(ctx context.Context, order *entity.Order) error {
	orderDB := models.ToOrderDB(order)
	_, err := o.db.ExecContext(
		ctx,
		queries.CreateOrder,
		orderDB.UserID,
		orderDB.AssetType,
		orderDB.OrderType,
		orderDB.Price,
		orderDB.Quantity,
		orderDB.Status,
		orderDB.CreatedAt,
	)
	return err
}

func (o *OrderPersistence) GetOrderByUserIDAndStatus(ctx context.Context, userID int64, status entity.OrderStatus) ([]*entity.Order, error) {
	rows, err := o.db.QueryContext(ctx, queries.GetOrderByUserIDAndStatus, userID, status.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*entity.Order
	for rows.Next() {
		var orderDB models.Order
		err := orderDB.Scan(rows)
		if err != nil {
			return nil, err
		}
		orders = append(orders, orderDB.ToEntity())
	}
	return orders, nil
}

func (o *OrderPersistence) GetOrderByStatus(ctx context.Context, orderStatus entity.OrderStatus) ([]*entity.Order, error) {
	rows, err := o.db.QueryContext(ctx, queries.GetOrderByStatus, orderStatus.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*entity.Order
	for rows.Next() {
		var orderDB models.Order
		err := orderDB.Scan(rows)
		if err != nil {
			return nil, err
		}
		orders = append(orders, orderDB.ToEntity())
	}
	return orders, nil
}
