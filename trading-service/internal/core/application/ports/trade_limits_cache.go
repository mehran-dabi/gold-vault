package ports

import "context"

type (
	TradeLimitsCache interface {
		SetGlobalLimits(ctx context.Context, assetType string, minBuy, minSell, maxBuy, maxSell, dailyBuyLimit, dailySellLimit float64) error
		GetGlobalLimits(ctx context.Context, assetType string) (map[string]float64, error)
		GetUserDailyLimits(ctx context.Context, userID int64, assetType string) (map[string]float64, error)
		SetUserDailyLimits(ctx context.Context, userID int64, assetType string, dailySellLimit, dailyBuyLimit float64) error
	}
)
