package models

import (
	"time"

	"goldvault/wallet-service/internal/core/domain/entity"
)

type Asset struct {
	ID        int64
	WalletID  int64
	Type      string
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

const (
	AssetsTableName   = "assets"
	AssetsColumns     = "id, wallet_id, type, balance, created_at, updated_at"
	AssetsColumnsNoID = "wallet_id, type, balance, created_at, updated_at"
)

func ToAssetDB(entity *entity.Asset) (*Asset, error) {
	return &Asset{
		ID:        entity.ID,
		WalletID:  entity.WalletID,
		Type:      entity.Type,
		Balance:   entity.Balance,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}, nil
}

func (a *Asset) Scan(scanner Scanner) error {
	return scanner.Scan(
		&a.ID,
		&a.WalletID,
		&a.Type,
		&a.Balance,
		&a.CreatedAt,
		&a.UpdatedAt,
	)
}

func (a *Asset) ToEntity() *entity.Asset {
	return &entity.Asset{
		ID:        a.ID,
		WalletID:  a.WalletID,
		Type:      a.Type,
		Balance:   a.Balance,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}
