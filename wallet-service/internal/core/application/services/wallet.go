package services

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"goldvault/wallet-service/internal/core/application/ports"
	"goldvault/wallet-service/internal/core/domain/entity"
	"goldvault/wallet-service/internal/infrastructure/db"
	"goldvault/wallet-service/pkg/serr"
)

type WalletService struct {
	walletDomainService ports.WalletDomainService
	assetDomainService  ports.AssetDomainService
	assetServiceClient  ports.AssetServiceClientPorts
}

func NewWalletService(
	walletDomainService ports.WalletDomainService,
	assetDomainService ports.AssetDomainService,
	assetServiceClient ports.AssetServiceClientPorts,
) *WalletService {
	return &WalletService{
		walletDomainService: walletDomainService,
		assetDomainService:  assetDomainService,
		assetServiceClient:  assetServiceClient,
	}
}

func (w *WalletService) CreateWallet(ctx context.Context, userID int64) (*entity.Wallet, error) {
	wallet, err := w.walletDomainService.CreateWallet(ctx, userID)
	if err != nil {
		return nil, err
	}

	// create default IRR asset
	err = w.assetDomainService.AddAsset(ctx, wallet.ID, entity.AssetTypeIRR, 0)

	return wallet, nil
}

func (w *WalletService) GetUserWallet(ctx context.Context, userID int64) (*entity.Wallet, error) {
	wallet, err := w.walletDomainService.GetUserWallet(ctx, userID)
	if err != nil {
		return nil, serr.ServiceErr("WalletApplicationService.GetUserWallet", err.Error(), err, http.StatusInternalServerError)
	}

	if wallet == nil {
		return nil, serr.ServiceErr("WalletApplicationService.GetUserWallet", "wallet not found", fmt.Errorf("wallet not found"), http.StatusNotFound)
	}

	// get wallet assets
	assets, err := w.assetDomainService.GetAssetsByWalletID(ctx, wallet.ID)
	if err != nil {
		return nil, serr.ServiceErr("WalletApplicationService.GetUserWallet", err.Error(), err, http.StatusInternalServerError)
	}

	// get asset prices
	assetTypes := make([]string, 0)
	for _, asset := range assets {
		assetTypes = append(assetTypes, asset.Type)
	}
	assetPrices, err := w.assetServiceClient.GetAssetPrice(ctx, assetTypes)
	if err != nil {
		return nil, serr.ServiceErr("WalletApplicationService.GetUserWallet", err.Error(), err, http.StatusInternalServerError)
	}

	// set prices in the assets
	for i, asset := range assets {
		assets[i].TotalPrice = asset.Balance * assetPrices[asset.Type].SellPrice
	}

	wallet.Assets = assets

	return wallet, nil
}

func (w *WalletService) GetWalletsWithPagination(ctx context.Context, limit, offset int) ([]*entity.Wallet, error) {
	wallets, err := w.walletDomainService.GetWalletsWithPagination(ctx, limit, offset)
	if err != nil {
		return nil, serr.ServiceErr("WalletApplicationService.GetWalletsWithPagination", err.Error(), err, http.StatusInternalServerError)
	}

	return wallets, nil
}

func (w *WalletService) Withdraw(ctx context.Context, userID int64, assetType string, amount float64) error {
	// validate user ID
	if userID <= 0 {
		return serr.ServiceErr("WalletApplicationService.Withdraw", "invalid user ID", fmt.Errorf("invalid user ID"), http.StatusBadRequest)
	}

	// validate amount
	if amount <= 0 {
		return serr.ServiceErr("WalletApplicationService.Withdraw", "invalid amount", fmt.Errorf("invalid amount"), http.StatusBadRequest)
	}

	// validate asset type
	if assetType == "" {
		return serr.ServiceErr("WalletApplicationService.Withdraw", "invalid asset type", fmt.Errorf("invalid asset type"), http.StatusBadRequest)
	}

	err := db.Transaction(ctx, sql.LevelReadCommitted, func(tx *sql.Tx) error {
		// update asset balance and reduce the balance
		err := w.assetDomainService.UpdateAssetBalance(ctx, tx, userID, assetType, -amount)
		if err != nil {
			return serr.ServiceErr("WalletApplicationService.Withdraw", err.Error(), err, http.StatusInternalServerError)
		}

		// update asset balance and add to the IRR balance
		err = w.assetDomainService.UpdateAssetBalance(ctx, tx, userID, entity.AssetTypeIRR, amount)
		if err != nil {
			return serr.ServiceErr("WalletApplicationService.Withdraw", err.Error(), err, http.StatusInternalServerError)
		}
		return nil
	})
	if err != nil {
		return serr.ServiceErr("WalletApplicationService.Withdraw", err.Error(), err, http.StatusInternalServerError)
	}

	return nil
}

func (w *WalletService) Deposit(ctx context.Context, userID int64, assetType string, amount float64) error {
	// validate user ID
	if userID <= 0 {
		return serr.ServiceErr("WalletApplicationService.Deposit", "invalid user ID", fmt.Errorf("invalid user ID"), http.StatusBadRequest)
	}

	// validate amount
	if amount <= 0 {
		return serr.ServiceErr("WalletApplicationService.Deposit", "invalid amount", fmt.Errorf("invalid amount"), http.StatusBadRequest)
	}

	// validate asset type
	if assetType == "" {
		return serr.ServiceErr("WalletApplicationService.Deposit", "invalid asset type", fmt.Errorf("invalid asset type"), http.StatusBadRequest)
	}

	err := db.Transaction(ctx, sql.LevelReadCommitted, func(tx *sql.Tx) error {
		// update IRR balance and reduce the balance
		err := w.assetDomainService.UpdateAssetBalance(ctx, tx, userID, entity.AssetTypeIRR, -amount)
		if err != nil {
			return serr.ServiceErr("WalletApplicationService.Deposit", err.Error(), err, http.StatusInternalServerError)
		}

		// update asset balance and add to the asset balance
		err = w.assetDomainService.UpdateAssetBalance(ctx, tx, userID, assetType, amount)
		if err != nil {
			return serr.ServiceErr("WalletApplicationService.Deposit", err.Error(), err, http.StatusInternalServerError)
		}
		return nil
	})
	if err != nil {
		return serr.ServiceErr("WalletApplicationService.Deposit", err.Error(), err, http.StatusInternalServerError)
	}

	return nil
}
