package grpc

import (
	"context"
	"database/sql"
	"fmt"

	"goldvault/wallet-service/internal/core/application/services"
	"goldvault/wallet-service/internal/infrastructure/db"
	"goldvault/wallet-service/proto" // Import the generated protobuf code for WalletService
)

// WalletGRPCHandler implements the WalletService gRPC server
type WalletGRPCHandler struct {
	proto.UnimplementedWalletServiceServer // Embedding for forward compatibility
	walletAppService                       *services.WalletService
	assetAppService                        *services.AssetService
}

// NewWalletGRPCHandler creates a new WalletGRPCHandler
func NewWalletGRPCHandler(
	walletAppService *services.WalletService,
	assetAppService *services.AssetService,
) *WalletGRPCHandler {
	return &WalletGRPCHandler{
		walletAppService: walletAppService,
		assetAppService:  assetAppService,
	}
}

// CreateWallet creates a new wallet for a user
func (w *WalletGRPCHandler) CreateWallet(ctx context.Context, req *proto.CreateWalletRequest) (*proto.CreateWalletResponse, error) {
	// Extract the user ID from the request
	userID := req.UserId

	// Validate user ID
	if userID <= 0 {
		return nil, fmt.Errorf("invalid user ID")
	}

	// Persist the wallet using the persistence layer
	wallet, err := w.walletAppService.CreateWallet(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to create wallet: %w", err)
	}

	// Return the response with the new wallet ID
	return &proto.CreateWalletResponse{
		WalletId: wallet.ID,
	}, nil
}

// UpdateAssetBalance updates the asset balance for a user
func (w *WalletGRPCHandler) UpdateAssetBalance(ctx context.Context, req *proto.UpdateAssetBalanceRequest) (*proto.UpdateAssetBalanceResponse, error) {
	// Extract user id from the request
	userID := req.GetUserId()

	// Validate user ID
	if userID <= 0 {
		return nil, fmt.Errorf("invalid user ID")
	}

	// update the asset balance using the persistence layer
	err := w.assetAppService.UpdateAssetBalance(ctx, userID, req.GetAssetType(), req.GetAmount())
	if err != nil {
		return nil, fmt.Errorf("failed to update asset balance: %w", err)
	}

	// Return the response with the updated asset balance
	return &proto.UpdateAssetBalanceResponse{}, nil
}

// Withdraw withdraws an amount from the user's wallet
func (w *WalletGRPCHandler) Withdraw(ctx context.Context, req *proto.WithdrawRequest) (*proto.WithdrawResponse, error) {
	userID := req.GetUserId()
	amount := req.GetAmount()
	assetType := req.GetAssetType()

	if userID <= 0 {
		return nil, fmt.Errorf("invalid user ID")
	}

	err := db.Transaction(ctx, sql.LevelReadCommitted, func(tx *sql.Tx) error {
		// Update the asset and reduce the balance
		err := w.assetAppService.UpdateAssetBalance(ctx, userID, assetType, -amount)
		if err != nil {
			return fmt.Errorf("failed to withdraw amount: %w", err)
		}

		// Update the asset and add to the IRR balance
		err = w.assetAppService.UpdateAssetBalance(ctx, userID, "IRR", amount)
		if err != nil {
			return fmt.Errorf("failed to withdraw amount: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to withdraw amount: %w", err)
	}

	return &proto.WithdrawResponse{}, nil
}

// Deposit deposits an amount to the user's wallet
func (w *WalletGRPCHandler) Deposit(ctx context.Context, req *proto.DepositRequest) (*proto.DepositResponse, error) {
	userID := req.GetUserId()
	amount := req.GetAmount()
	assetType := req.GetAssetType()

	if userID <= 0 {
		return nil, fmt.Errorf("invalid user ID")
	}

	err := w.assetAppService.UpdateAssetBalance(ctx, userID, assetType, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to deposit amount: %w", err)
	}

	return &proto.DepositResponse{}, nil
}

func (w *WalletGRPCHandler) GetAssetBalance(ctx context.Context, req *proto.GetAssetBalanceRequest) (*proto.GetAssetBalanceResponse, error) {
	userID := req.GetUserId()
	assetType := req.GetAssetType()

	if userID <= 0 {
		return nil, fmt.Errorf("invalid user ID")
	}

	balance, err := w.assetAppService.GetAssetBalance(ctx, userID, assetType)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset balance: %w", err)
	}

	return &proto.GetAssetBalanceResponse{
		Balance: balance,
	}, nil
}
