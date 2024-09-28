package clients

import (
	"context"
	"fmt"
	"log"

	"goldvault/wallet-service/internal/core/domain/entity"
	"goldvault/wallet-service/proto"

	"google.golang.org/grpc"
)

type AssetServiceClient struct {
	client proto.AssetServiceClient
}

func NewAssetServiceClient(address string) (*AssetServiceClient, error) {
	// Set up a connection to the Wallet Service
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}

	client := proto.NewAssetServiceClient(conn)
	return &AssetServiceClient{client: client}, nil
}

func (a *AssetServiceClient) GetAssetPrice(ctx context.Context, assetTypes []string) (map[string]*entity.PriceDetails, error) {
	req := &proto.AssetPriceRequest{
		AssetTypes: assetTypes,
	}

	resp, err := a.client.GetAssetPrice(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset price: %v", err)
	}

	assetPrices := make(map[string]*entity.PriceDetails)
	for asset, price := range resp.Prices {
		assetPrices[asset] = &entity.PriceDetails{
			BuyPrice:  price.GetBuy(),
			SellPrice: price.GetSell(),
		}
	}

	return assetPrices, nil
}
