package api

import (
	"goldvault/asset-service/internal/server"
	"goldvault/asset-service/internal/server/middlewares"
)

func SetupAssetPriceRoutes(s *server.Server, a *AssetPriceHandler, ad *AdminAssetPriceHandler) {
	assetPrices := s.External.Group("/asset-prices")
	assetPrices.Use(middlewares.JWTMiddleware(), middlewares.RoleMiddleware("customer"))
	{
		assetPrices.GET("/:assetType", a.GetPrice)
		assetPrices.GET("/", a.GetAllAssetPrices)
	}

	admin := s.External.Group("/admin")
	admin.Use(middlewares.JWTMiddleware(), middlewares.RoleMiddleware("admin"))
	{
		admin.POST("/asset-prices", ad.UpsertPrice)
		admin.DELETE("/asset-prices/:assetType", ad.DeleteAssetPrice)
		admin.PUT("/asset-prices/:assetType/adjust-by-step", ad.UpdateAssetPriceByStep)
		admin.POST("/asset-prices/step", ad.SetPriceChangeStep)
		admin.GET("/asset-prices/step", ad.GetPriceChangeStep)
	}
}

func SetupPriceHistoryRoutes(s *server.Server, ph *PriceHistoryHandler) {
	priceHistory := s.External.Group("/price-history")
	priceHistory.Use(middlewares.JWTMiddleware(), middlewares.RoleMiddleware("customer"))
	{
		priceHistory.GET("/:assetType", ph.GetAssetPriceHistory)
	}
}
