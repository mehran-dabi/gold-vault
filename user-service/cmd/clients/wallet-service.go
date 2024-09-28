package clients

import (
	"context"
	"fmt"
	"log"

	"goldvault/user-service/internal/core/application/ports"
	"goldvault/user-service/proto" // Import the generated protobuf code for WalletService

	"google.golang.org/grpc"
)

type WalletClient struct {
	client proto.WalletServiceClient
}

func NewWalletClient(address string) (ports.WalletClientPorts, error) {
	// Set up a connection to the Wallet Service
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}

	client := proto.NewWalletServiceClient(conn)
	return &WalletClient{client: client}, nil
}

// CreateWallet calls the CreateWallet gRPC method on the Wallet Service
func (w *WalletClient) CreateWallet(ctx context.Context, userID int64) (int64, error) {
	// Create the request for the gRPC call
	req := &proto.CreateWalletRequest{
		UserId: userID,
	}

	// Call the CreateWallet method
	resp, err := w.client.CreateWallet(ctx, req)
	if err != nil {
		return 0, fmt.Errorf("failed to create wallet: %w", err)
	}

	// Return the wallet ID from the response
	return resp.WalletId, nil
}
