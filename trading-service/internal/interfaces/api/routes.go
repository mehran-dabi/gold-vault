package api

import (
	"goldvault/trading-service/internal/server"
	"goldvault/trading-service/internal/server/middlewares"
)

func SetupTradingRoutes(s *server.Server, t *TransactionsHandler, i *InventoryHandler) {
	trading := s.External.Group("/trades")
	trading.Use(middlewares.JWTMiddleware(), middlewares.RoleMiddleware("customer"))
	{
		trading.GET("/transactions", t.GetUserTransactions)
		trading.POST("/buy", i.BuyAsset)
		trading.POST("/sell", i.SellAsset)
	}
}

func SetupAdminRoutes(s *server.Server, i *InventoryAdminHandler, t *TransactionsAdminHandler) {
	admin := s.External.Group("/admin/trades")
	admin.Use(middlewares.JWTMiddleware(), middlewares.RoleMiddleware("admin"))
	{
		admin.PATCH("/ignore-inventory-limit", i.UpdateIgnoreInventoryLimit)
		inventory := admin.Group("/inventory")
		{
			inventory.GET("", i.GetInventory)
			inventory.POST("", i.CreateInventory)
			inventory.DELETE("/:assetType", i.DeleteInventory)
			inventory.POST("/:assetType", i.SetGlobalTradeLimits)
		}

		transactions := admin.Group("/transactions")
		{
			transactions.GET("", t.GetTransactions)
			transactions.GET("/:userID", t.GetUserTransactions)
			transactions.GET("/summary/single-day", t.GetSingleDaySummary)
			transactions.GET("/summary/total", t.GetTotalSummary)
		}

	}
}
