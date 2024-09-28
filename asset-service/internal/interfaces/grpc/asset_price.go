package grpc

import (
	"context"
	"fmt"

	"goldvault/asset-service/internal/core/application/services"
	"goldvault/asset-service/internal/core/domain/entity"
	"goldvault/asset-service/proto"
)

// AssetPriceGRPCHandler implements the AssetService gRPC server
type AssetPriceGRPCHandler struct {
	proto.UnimplementedAssetServiceServer                             // Embedding for forward compatibility
	assetAppService                       *services.AssetPriceService // Interface for persistence operations
}

// NewAssetPriceGRPCHandler creates a new AssetPriceGRPCHandler
func NewAssetPriceGRPCHandler(assetAppService *services.AssetPriceService) *AssetPriceGRPCHandler {
	return &AssetPriceGRPCHandler{
		assetAppService: assetAppService,
	}
}

func (s *AssetPriceGRPCHandler) GetAssetPrice(ctx context.Context, req *proto.AssetPriceRequest) (*proto.AssetPriceResponse, error) {
	// Extract the assetType from the request
	assetTypes := req.GetAssetTypes()

	// Validate assetType
	if len(assetTypes) == 0 {
		return nil, fmt.Errorf("asset type is required")
	}
	for _, assetType := range assetTypes {
		if !entity.IsValidAssetType(assetType) {
			return nil, fmt.Errorf("invalid asset type: %s", assetType)
		}
	}

	// Get the asset price using the persistence layer
	assetPrices, err := s.assetAppService.GetLatestPrices(ctx, assetTypes)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset price: %w", err)
	}

	// convert entity.PriceDetails to proto.PriceDetails
	priceDetailsProto := make(map[string]*proto.PriceDetails)
	for assetType, priceDetails := range assetPrices {
		priceDetailsProto[assetType] = &proto.PriceDetails{
			Buy:  priceDetails.BuyPrice,
			Sell: priceDetails.SellPrice,
		}
	}

	// Return the response with the asset price
	return &proto.AssetPriceResponse{
		Prices: priceDetailsProto,
	}, nil
}
