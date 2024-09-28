package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"goldvault/wallet-service/internal/core/application/ports"
	"goldvault/wallet-service/internal/core/domain/entity"
)

type AssetService struct {
	assetPersistence  ports.AssetPersistence
	walletPersistence ports.WalletPersistence
}

func NewAssetDomainService(
	assetPersistence ports.AssetPersistence,
	walletPersistence ports.WalletPersistence,
) ports.AssetDomainService {
	return &AssetService{
		assetPersistence:  assetPersistence,
		walletPersistence: walletPersistence,
	}
}

func (a *AssetService) AddAsset(ctx context.Context, walletID int64, assetType string, amount float64) error {
	assetEntity := &entity.Asset{
		WalletID:  walletID,
		Type:      assetType,
		Balance:   amount,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := a.assetPersistence.AddAsset(ctx, assetEntity)
	if err != nil {
		return err
	}

	return nil
}

func (a *AssetService) GetAssetsByWalletID(ctx context.Context, walletID int64) ([]entity.Asset, error) {
	assets, err := a.assetPersistence.GetAssetsByWalletID(ctx, walletID)
	if err != nil {
		return nil, err
	}

	return assets, nil
}

func (a *AssetService) UpdateAssetBalance(ctx context.Context, tx *sql.Tx, userID int64, assetType string, amount float64) error {
	//// Retrieve user_id from the context
	//userID, ok := ctx.Value("user_id").(int64)
	//if !ok {
	//	return fmt.Errorf("user_id not found in context")
	//}
	//
	//// Check if the wallet_id belongs to the user_id
	wallet, err := a.walletPersistence.GetWalletByUserID(ctx, userID)
	if err != nil {
		return err
	}
	//if wallet.ID != walletID {
	//	return fmt.Errorf("wallet_id does not match the user_id")
	//}

	// Lock the asset row for update
	asset, err := a.assetPersistence.GetAssetForUpdate(ctx, tx, wallet.ID, assetType)
	if err != nil && asset != nil { // if the asset does not exist, skip...
		return err
	}

	// Upsert the asset balance
	err = a.assetPersistence.UpsertAssetBalance(ctx, tx, wallet.ID, assetType, amount)
	if err != nil {
		return err
	}

	return nil
}

func (a *AssetService) GetAssetBalance(ctx context.Context, userID int64, assetType string) (float64, error) {
	wallet, err := a.walletPersistence.GetWalletByUserID(ctx, userID)
	if err != nil {
		return 0, err
	}

	if wallet == nil {
		return 0, fmt.Errorf("wallet not found")
	}

	balance, err := a.assetPersistence.GetAssetBalance(ctx, wallet.ID, assetType)
	if err != nil {
		return 0, err
	}

	return balance, nil
}
