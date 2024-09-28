package grpc

import (
	"context"
	"fmt"

	"goldvault/wallet-service/internal/core/application/services"
	"goldvault/wallet-service/proto"
)

type AssetGRPCHandler struct {
	proto.UnimplementedAssetServiceServer
	assetAppService *services.AssetService
}

func NewAssetGRPCHandler(assetAppService *services.AssetService) *AssetGRPCHandler {
	return &AssetGRPCHandler{
		assetAppService: assetAppService,
	}
}

// UpdateAssetBalance updates the asset balance for a user
func (a *AssetGRPCHandler) UpdateAssetBalance(ctx context.Context, req *proto.UpdateAssetBalanceRequest) (*proto.UpdateAssetBalanceResponse, error) {
	// Extract user id from the request
	userID := req.GetUserId()

	// Validate user ID
	if userID <= 0 {
		return nil, fmt.Errorf("invalid user ID")
	}

	// update the asset balance using the persistence layer
	err := a.assetAppService.UpdateAssetBalance(ctx, userID, req.GetAssetType(), req.GetAmount())
	if err != nil {
		return nil, fmt.Errorf("failed to update asset balance: %w", err)
	}

	// Return the response with the updated asset balance
	return &proto.UpdateAssetBalanceResponse{}, nil
}
