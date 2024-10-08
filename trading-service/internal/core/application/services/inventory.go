package services

import (
	"context"
	"database/sql"
	"net/http"

	"goldvault/trading-service/internal/core/application/ports"
	"goldvault/trading-service/internal/core/domain/entity"
	"goldvault/trading-service/internal/infrastructure/db"
	"goldvault/trading-service/pkg/serr"
)

type InventoryService struct {
	inventoryDomainService   ports.InventoryDomainService
	transactionDomainService ports.TransactionDomainService
	walletServiceClient      ports.WalletServiceClient
	ignoreInvLimitCache      ports.IgnoreInventoryLimitCache
}

func NewInventoryService(
	inventoryDomainService ports.InventoryDomainService,
	transactionDomainService ports.TransactionDomainService,
	walletServiceClient ports.WalletServiceClient,
	ignoreInventoryLimitCache ports.IgnoreInventoryLimitCache,
) *InventoryService {
	return &InventoryService{
		inventoryDomainService:   inventoryDomainService,
		transactionDomainService: transactionDomainService,
		walletServiceClient:      walletServiceClient,
		ignoreInvLimitCache:      ignoreInventoryLimitCache,
	}
}

func (i *InventoryService) BuyAsset(ctx context.Context, userID int64, assetType string, quantity, price float64) error {
	transaction := &entity.Transaction{
		UserID:          userID,
		AssetType:       assetType,
		Quantity:        quantity,
		Price:           price,
		TransactionType: entity.TransactionTypeBuy,
		Status:          entity.TransactionStatusPending,
	}

	// check if the user has enough IRR balance to buy the asset
	assetBalance, err := i.walletServiceClient.GetAssetBalance(ctx, userID, "IRR")
	if err != nil {
		return serr.ServiceErr("Inventory.BuyAsset", err.Error(), err, http.StatusInternalServerError)
	}
	if assetBalance < quantity*price {
		return serr.ServiceErr("Inventory.BuyAsset", "insufficient balance to buy asset", nil, http.StatusBadRequest)
	}

	// check if the ignore inventory flag is set.
	// this flag is used to ignore the inventory limit when buying an asset
	ignoreInv, err := i.ignoreInvLimitCache.Get(ctx)
	if err != nil {
		return serr.ServiceErr("Inventory.BuyAsset", err.Error(), err, http.StatusInternalServerError)
	}

	err = db.Transaction(ctx, sql.LevelReadCommitted, func(tx *sql.Tx) error {
		// log pending transaction
		err := i.transactionDomainService.LogTransaction(ctx, tx, transaction)
		if err != nil {
			return err
		}

		err = i.inventoryDomainService.Buy(ctx, tx, assetType, quantity, ignoreInv)
		if err != nil {
			// update transaction to failed
			if err := i.transactionDomainService.UpdateTransactionStatus(ctx, transaction.ID, entity.TransactionStatusFailed.String()); err != nil {
				return err
			}
			return err
		}
		return nil
	})
	if err != nil {
		return serr.ServiceErr("Inventory.BuyAsset", err.Error(), err, http.StatusInternalServerError)
	}

	// call wallets-service to update user's wallet
	err = i.walletServiceClient.Deposit(ctx, userID, assetType, quantity)
	if err != nil {
		// update transaction to balance pending
		if err := i.transactionDomainService.UpdateTransactionStatus(ctx, transaction.ID, entity.TransactionStatusBalancePending.String()); err != nil {
			return err
		}
		return serr.ServiceErr("Inventory.BuyAsset", err.Error(), err, http.StatusInternalServerError)
	}

	// update transaction to completed
	err = i.transactionDomainService.UpdateTransactionStatus(ctx, transaction.ID, entity.TransactionStatusCompleted.String())
	if err != nil {
		return err
	}

	return nil
}

func (i *InventoryService) SellAsset(ctx context.Context, userID int64, assetType string, quantity, price float64) error {
	transaction := &entity.Transaction{
		UserID:          userID,
		AssetType:       assetType,
		Quantity:        quantity,
		Price:           price,
		TransactionType: entity.TransactionTypeSell,
		Status:          entity.TransactionStatusPending,
	}

	// check if the user has enough asset balance to sell
	assetBalance, err := i.walletServiceClient.GetAssetBalance(ctx, userID, assetType)
	if err != nil {
		return serr.ServiceErr("Inventory.SellAsset", err.Error(), err, http.StatusInternalServerError)
	}

	if assetBalance < quantity {
		return serr.ServiceErr("Inventory.SellAsset", "insufficient balance to sell asset", nil, http.StatusBadRequest)
	}

	err = db.Transaction(ctx, sql.LevelReadCommitted, func(tx *sql.Tx) error {
		// log pending transaction
		err := i.transactionDomainService.LogTransaction(ctx, tx, transaction)
		if err != nil {
			return err
		}

		err = i.inventoryDomainService.Sell(ctx, tx, assetType, quantity)
		if err != nil {
			// update transaction to failed
			if err := i.transactionDomainService.UpdateTransactionStatus(ctx, transaction.ID, entity.TransactionStatusFailed.String()); err != nil {
				return err
			}
			return err
		}

		return nil
	})
	if err != nil {
		return serr.ServiceErr("Inventory.SellAsset", err.Error(), err, http.StatusInternalServerError)
	}

	// call wallets-service to update user's wallet
	err = i.walletServiceClient.Withdraw(ctx, userID, assetType, quantity)
	if err != nil {
		// update transaction to balance pending
		if err := i.transactionDomainService.UpdateTransactionStatus(ctx, transaction.ID, entity.TransactionStatusBalancePending.String()); err != nil {
			return err
		}
		return serr.ServiceErr("Inventory.BuyAsset", err.Error(), err, http.StatusInternalServerError)
	}

	// update transaction to completed
	err = i.transactionDomainService.UpdateTransactionStatus(ctx, transaction.ID, entity.TransactionStatusCompleted.String())
	if err != nil {
		return err
	}

	return nil
}

func (i *InventoryService) GetInventory(ctx context.Context) ([]*entity.Inventory, error) {
	return i.inventoryDomainService.GetInventory(ctx)
}

func (i *InventoryService) CreateInventory(ctx context.Context, assetType string, quantity float64) (int64, error) {
	inventory := &entity.Inventory{
		AssetType:     entity.AssetType(assetType),
		TotalQuantity: quantity,
	}
	return i.inventoryDomainService.CreateInventory(ctx, inventory)
}

func (i *InventoryService) UpdateInventoryQuantity(ctx context.Context, assetType string, quantity float64) error {
	return i.inventoryDomainService.UpdateInventoryQuantity(ctx, nil, assetType, quantity)
}

func (i *InventoryService) DeleteInventory(ctx context.Context, assetType string) error {
	return i.inventoryDomainService.DeleteInventory(ctx, assetType)
}
