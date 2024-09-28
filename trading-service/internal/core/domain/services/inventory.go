package services

import (
	"context"
	"database/sql"
	"fmt"

	"goldvault/trading-service/internal/core/application/ports"
	"goldvault/trading-service/internal/core/domain/entity"
)

type InventoryDomainService struct {
	inventoryPersistence ports.InventoryPersistence
}

func NewInventoryDomainService(inventoryPersistence ports.InventoryPersistence) ports.InventoryDomainService {
	return &InventoryDomainService{
		inventoryPersistence: inventoryPersistence,
	}
}

func (i *InventoryDomainService) Buy(ctx context.Context, tx *sql.Tx, assetType string, quantity float64) error {
	inventory, err := i.inventoryPersistence.GetInventoryForUpdate(ctx, tx, assetType)
	if err != nil {
		return err
	}

	if inventory.TotalQuantity < quantity {
		return fmt.Errorf("insufficient inventory for asset type %s", assetType)
	}

	inventory.TotalQuantity -= quantity

	if err := inventory.Validate(); err != nil {
		return err
	}

	err = i.inventoryPersistence.UpdateInventory(ctx, tx, inventory)
	if err != nil {
		return err
	}

	return nil
}

func (i *InventoryDomainService) Sell(ctx context.Context, tx *sql.Tx, assetType string, quantity float64) error {
	inventory, err := i.inventoryPersistence.GetInventoryForUpdate(ctx, tx, assetType)
	if err != nil {
		return err
	}

	inventory.TotalQuantity += quantity

	if err := inventory.Validate(); err != nil {
		return err
	}

	err = i.inventoryPersistence.UpdateInventory(ctx, tx, inventory)
	if err != nil {
		return err
	}

	return nil
}

func (i *InventoryDomainService) UpdateInventoryQuantity(ctx context.Context, tx *sql.Tx, assetType string, quantity float64) error {
	// lock the inventory row for update
	inventory, err := i.inventoryPersistence.GetInventoryForUpdate(ctx, tx, assetType)
	if err != nil {
		return err
	}

	// check sufficient inventory
	if inventory.TotalQuantity < quantity {
		return fmt.Errorf("insufficient inventory for asset type %s", assetType)
	}

	inventory.TotalQuantity += quantity

	// validate inventory
	if err := inventory.Validate(); err != nil {
		return err
	}

	// update inventory
	err = i.inventoryPersistence.UpdateInventory(ctx, tx, inventory)
	if err != nil {
		return err
	}

	return nil
}

func (i *InventoryDomainService) GetInventory(ctx context.Context) ([]*entity.Inventory, error) {
	return i.inventoryPersistence.GetInventory(ctx)
}

func (i *InventoryDomainService) CreateInventory(ctx context.Context, inventory *entity.Inventory) (int64, error) {
	// validate inventory
	if err := inventory.Validate(); err != nil {
		return 0, err
	}

	id, err := i.inventoryPersistence.CreateInventory(ctx, inventory)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (i *InventoryDomainService) DeleteInventory(ctx context.Context, assetType string) error {
	err := i.inventoryPersistence.DeleteInventory(ctx, assetType)
	if err != nil {
		return err
	}

	return nil
}
