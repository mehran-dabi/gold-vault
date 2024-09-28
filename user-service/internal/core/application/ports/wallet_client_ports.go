package ports

import "context"

type WalletClientPorts interface {
	CreateWallet(ctx context.Context, userID int64) (int64, error)
}
