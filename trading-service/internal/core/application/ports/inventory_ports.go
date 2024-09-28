package ports

import (
	"context"
	"database/sql"

	"goldvault/trading-service/internal/core/domain/entity"
)

type (
	InventoryPersistence interface {
		GetInventoryForUpdate(ctx context.Context, tx *sql.Tx, assetType string) (*entity.Inventory, error)
		UpdateInventory(ctx context.Context, tx *sql.Tx, inventory *entity.Inventory) error
		GetInventory(ctx context.Context) ([]*entity.Inventory, error)
		CreateInventory(ctx context.Context, inventory *entity.Inventory) (int64, error)
		DeleteInventory(ctx context.Context, assetType string) error
	}

	InventoryDomainService interface {
		UpdateInventoryQuantity(ctx context.Context, tx *sql.Tx, assetType string, quantity float64) error
		Buy(ctx context.Context, tx *sql.Tx, assetType string, quantity float64) error
		Sell(ctx context.Context, tx *sql.Tx, assetType string, quantity float64) error
		GetInventory(ctx context.Context) ([]*entity.Inventory, error)
		CreateInventory(ctx context.Context, inventory *entity.Inventory) (int64, error)
		DeleteInventory(ctx context.Context, assetType string) error
	}
)
