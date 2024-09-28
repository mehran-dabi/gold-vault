package persistence

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"goldvault/wallet-service/internal/core/application/ports"
	"goldvault/wallet-service/internal/core/domain/entity"
	"goldvault/wallet-service/internal/infrastructure/persistence/models"
	"goldvault/wallet-service/internal/infrastructure/persistence/queries"
	"goldvault/wallet-service/pkg/serr"
)

type WalletPersistence struct {
	db *sql.DB
}

func NewWalletPersistence(db *sql.DB) ports.WalletPersistence {
	return &WalletPersistence{db: db}
}

// SaveWallet inserts a new wallet into the database
func (p *WalletPersistence) SaveWallet(ctx context.Context, wallet *entity.Wallet) error {
	dbModel, err := models.ToWalletDB(wallet)
	if err != nil {
		return serr.DBError("CreateWallet", "wallet", err)
	}

	err = p.db.QueryRowContext(ctx, queries.CreateWallet, dbModel.UserID, time.Now(), time.Now()).Scan(&dbModel.ID)
	if err != nil {
		return serr.DBError("CreateWallet", "wallet", err)
	}

	// Update the wallet ID
	wallet.ID = dbModel.ID

	return nil
}

// GetWalletByUserID retrieves a wallet by the associated user ID
func (p *WalletPersistence) GetWalletByUserID(ctx context.Context, userID int64) (*entity.Wallet, error) {
	row := p.db.QueryRowContext(ctx, queries.GetWalletByUserID, userID)

	var dbModel models.Wallet
	err := dbModel.Scan(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No wallet found
		}
		return nil, serr.DBError("GetWalletByUserID", "wallet", err)
	}

	return dbModel.ToEntity(), nil
}

func (p *WalletPersistence) GetWalletsWithPagination(ctx context.Context, limit, offset int) ([]*entity.Wallet, error) {
	rows, err := p.db.QueryContext(ctx, queries.GetWalletsWithPagination, limit, offset)
	if err != nil {
		return nil, serr.DBError("GetWalletsWithPagination", "wallet", err)
	}
	defer rows.Close()

	var wallets []*entity.Wallet
	for rows.Next() {
		var dbModel models.Wallet
		err := dbModel.Scan(rows)
		if err != nil {
			return nil, serr.DBError("GetWalletsWithPagination", "wallet", err)
		}

		wallets = append(wallets, dbModel.ToEntity())
	}

	return wallets, nil
}
