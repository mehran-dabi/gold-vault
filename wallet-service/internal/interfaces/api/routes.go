package api

import (
	"goldvault/wallet-service/internal/server"
	"goldvault/wallet-service/internal/server/middlewares"
)

func SetupWalletRoutes(s *server.Server, w *WalletHandler, a *AdminWalletHandler) {
	wallets := s.External.Group("/wallets")
	wallets.Use(middlewares.JWTMiddleware(), middlewares.RoleMiddleware("customer"))
	{
		wallets.GET("/", w.GetUserWallet)
		wallets.GET("/transactions", w.GetWalletTransactions)
	}

	admin := s.External.Group("/admin")
	admin.Use(middlewares.JWTMiddleware(), middlewares.RoleMiddleware("admin"))
	{
		admin.GET("/wallets", a.GetWallets)
		admin.GET("/wallets/users/:id", a.GetWallet)
	}
}

func SetupTransactionRoutes(s *server.Server, a *AdminTransactionHandler) {
	admin := s.External.Group("/admin")
	admin.Use(middlewares.JWTMiddleware(), middlewares.RoleMiddleware("admin"))
	{
		admin.GET("/wallets/transactions", a.GetTransactions)
	}
}
