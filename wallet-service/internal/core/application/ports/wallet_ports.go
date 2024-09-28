package ports

import (
	"context"

	"goldvault/wallet-service/internal/core/domain/entity"
)

type (
	WalletPersistence interface {
		SaveWallet(ctx context.Context, wallet *entity.Wallet) error
		GetWalletByUserID(ctx context.Context, userID int64) (*entity.Wallet, error)
		GetWalletsWithPagination(ctx context.Context, limit, offset int) ([]*entity.Wallet, error)
	}

	WalletDomainService interface {
		CreateWallet(ctx context.Context, userID int64) (*entity.Wallet, error)
		GetUserWallet(ctx context.Context, userID int64) (*entity.Wallet, error)
		GetWalletsWithPagination(ctx context.Context, limit, offset int) ([]*entity.Wallet, error)
	}
)
