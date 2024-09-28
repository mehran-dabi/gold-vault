package ports

import "context"

type (
	WalletServiceClient interface {
		UpdateAssetBalance(ctx context.Context, userID int64, assetType string, quantity float64) error
		Withdraw(ctx context.Context, userID int64, assetType string, quantity float64) error
		Deposit(ctx context.Context, userID int64, assetType string, quantity float64) error
		GetAssetBalance(ctx context.Context, userID int64, assetType string) (float64, error)
	}
)
