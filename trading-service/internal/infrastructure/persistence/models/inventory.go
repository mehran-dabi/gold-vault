package models

import (
	"time"

	"goldvault/trading-service/internal/core/domain/entity"
)

type Inventory struct {
	ID            int64
	AssetType     string
	TotalQuantity float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

const (
	InventoryTableName   = "inventory"
	InventoryColumns     = "id, asset_type, total_quantity, created_at, updated_at"
	InventoryColumnsNoID = "asset_type, total_quantity, created_at, updated_at"
)

func (i *Inventory) Scan(scanner Scanner) error {
	return scanner.Scan(
		&i.ID,
		&i.AssetType,
		&i.TotalQuantity,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
}

func (i *Inventory) ToEntity() *entity.Inventory {
	return &entity.Inventory{
		ID:            i.ID,
		AssetType:     entity.AssetType(i.AssetType),
		TotalQuantity: i.TotalQuantity,
		CreatedAt:     i.CreatedAt,
		UpdatedAt:     i.UpdatedAt,
	}
}

func ToInventoryDB(entity *entity.Inventory) *Inventory {
	return &Inventory{
		ID:            entity.ID,
		AssetType:     entity.AssetType.String(),
		TotalQuantity: entity.TotalQuantity,
		CreatedAt:     entity.CreatedAt,
		UpdatedAt:     entity.UpdatedAt,
	}
}
