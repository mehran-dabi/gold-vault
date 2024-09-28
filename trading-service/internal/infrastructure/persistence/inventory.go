package persistence

import (
	"context"
	"database/sql"

	"goldvault/trading-service/internal/core/application/ports"
	"goldvault/trading-service/internal/core/domain/entity"
	"goldvault/trading-service/internal/infrastructure/persistence/models"
	"goldvault/trading-service/internal/infrastructure/persistence/queries"
	"goldvault/trading-service/pkg/serr"
)

type InventoryPersistence struct {
	db *sql.DB
}

func NewInventoryPersistence(db *sql.DB) ports.InventoryPersistence {
	return &InventoryPersistence{db: db}
}

func (i *InventoryPersistence) GetInventoryForUpdate(ctx context.Context, tx *sql.Tx, assetType string) (*entity.Inventory, error) {
	row := tx.QueryRowContext(
		ctx,
		queries.GetInventoryForUpdate,
		assetType,
	)
	inventory := &models.Inventory{}
	if err := inventory.Scan(row); err != nil {
		return nil, serr.DBError("GetInventoryForUpdate", "inventory", err)
	}
	return inventory.ToEntity(), nil
}

func (i *InventoryPersistence) UpdateInventory(ctx context.Context, tx *sql.Tx, inventory *entity.Inventory) error {
	inventoryDB := models.ToInventoryDB(inventory)
	_, err := tx.ExecContext(
		ctx,
		queries.UpdateInventoryQuantity,
		inventoryDB.AssetType,
		inventoryDB.TotalQuantity,
	)
	if err != nil {
		return serr.DBError("UpdateInventory", "inventory", err)
	}

	return nil
}

func (i *InventoryPersistence) CreateInventory(ctx context.Context, inventory *entity.Inventory) (int64, error) {
	inventoryDB := models.ToInventoryDB(inventory)
	var id int64
	err := i.db.QueryRowContext(
		ctx,
		queries.CreateInventory,
		inventoryDB.AssetType,
		inventoryDB.TotalQuantity,
	).Scan(&id)
	if err != nil {
		return 0, serr.DBError("CreateInventory", "inventory", err)
	}
	return id, nil
}

func (i *InventoryPersistence) GetInventory(ctx context.Context) ([]*entity.Inventory, error) {
	rows, err := i.db.QueryContext(ctx, queries.GetInventory)
	if err != nil {
		return nil, serr.DBError("GetInventory", "inventory", err)
	}
	defer rows.Close()
	inventories := make([]*entity.Inventory, 0)
	for rows.Next() {
		inventory := &models.Inventory{}
		if err := inventory.Scan(rows); err != nil {
			return nil, serr.DBError("GetInventory", "inventory", err)
		}
		inventories = append(inventories, inventory.ToEntity())
	}
	return inventories, nil
}

func (i *InventoryPersistence) DeleteInventory(ctx context.Context, assetType string) error {
	_, err := i.db.ExecContext(ctx, queries.DeleteInventory, assetType)
	if err != nil {
		return serr.DBError("DeleteInventory", "inventory", err)
	}
	return nil
}
