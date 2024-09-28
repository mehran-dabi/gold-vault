package clients

import (
	"context"
	"fmt"
	"log"

	"goldvault/trading-service/internal/core/application/ports"
	"goldvault/trading-service/proto"

	"google.golang.org/grpc"
)

type WalletServiceClient struct {
	client proto.WalletServiceClient
}

func NewWalletServiceClient(address string) (ports.WalletServiceClient, error) {
	// Set up a connection to the Wallet Service
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}

	client := proto.NewWalletServiceClient(conn)
	return &WalletServiceClient{client: client}, nil
}

func (w *WalletServiceClient) UpdateAssetBalance(ctx context.Context, userID int64, assetType string, quantity float64) error {
	req := &proto.UpdateAssetBalanceRequest{
		UserId:    userID,
		AssetType: assetType,
		Amount:    quantity,
	}

	_, err := w.client.UpdateAssetBalance(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to update asset balance: %v", err)
	}

	return nil
}

func (w *WalletServiceClient) Withdraw(ctx context.Context, userID int64, assetType string, quantity float64) error {
	req := &proto.WithdrawRequest{
		UserId:    userID,
		AssetType: assetType,
		Amount:    quantity,
	}

	_, err := w.client.Withdraw(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to withdraw: %v", err)
	}

	return nil
}

func (w *WalletServiceClient) Deposit(ctx context.Context, userID int64, assetType string, quantity float64) error {
	req := &proto.DepositRequest{
		UserId:    userID,
		AssetType: assetType,
		Amount:    quantity,
	}

	_, err := w.client.Deposit(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to deposit: %v", err)
	}

	return nil
}

func (w *WalletServiceClient) GetAssetBalance(ctx context.Context, userID int64, assetType string) (float64, error) {
	req := &proto.GetAssetBalanceRequest{
		UserId:    userID,
		AssetType: assetType,
	}

	resp, err := w.client.GetAssetBalance(ctx, req)
	if err != nil {
		return 0, fmt.Errorf("failed to get asset balance: %v", err)
	}

	return resp.Balance, nil
}
