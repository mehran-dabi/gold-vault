package api

import (
	"goldvault/trading-service/internal/server"
	"goldvault/trading-service/internal/server/middlewares"
)

func SetupTradingRoutes(s *server.Server, t *TransactionsHandler, i *InventoryHandler) {
	trading := s.External.Group("/trades")
	trading.Use(middlewares.JWTMiddleware(), middlewares.RoleMiddleware("customer"))
	{
		trading.GET("transactions", t.GetUserTransactions)
		trading.POST("/buy", i.BuyAsset)
		trading.POST("/sell", i.SellAsset)
	}
}

func SetupAdminRoutes(s *server.Server, i *InventoryAdminHandler, t *TransactionsAdminHandler) {
	admin := s.External.Group("/admin")
	admin.Use(middlewares.JWTMiddleware(), middlewares.RoleMiddleware("admin"))
	{
		inventory := admin.Group("/inventory")
		{
			inventory.GET("", i.GetInventory)
			inventory.POST("", i.CreateInventory)
			inventory.DELETE("/:assetType", i.DeleteInventory)
		}

		transactions := admin.Group("/transactions")
		{
			transactions.GET("", t.GetTransactions)
			transactions.GET("/:userID", t.GetUserTransactions)
		}

	}
}
