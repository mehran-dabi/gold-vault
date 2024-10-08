package main

import (
	"goldvault/trading-service/internal/config"
	appService "goldvault/trading-service/internal/core/application/services"
	"goldvault/trading-service/internal/core/domain/services"
	"goldvault/trading-service/internal/infrastructure/cache"
	"goldvault/trading-service/internal/infrastructure/db"
	"goldvault/trading-service/internal/infrastructure/persistence"
	"goldvault/trading-service/internal/interfaces/api"
	"goldvault/trading-service/internal/server"
	"goldvault/trading-service/pkg/logger"

	"go.uber.org/fx"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	fx.New(
		fx.Provide(
			// external clients
			assetService,
			walletService,

			// postgres
			postgresDB,

			// redis
			redisCache,
			redisRateLimit,

			// persistence
			persistence.NewTransactionPersistence,
			persistence.NewOrderPersistence,
			persistence.NewInventoryPersistence,

			// cache
			cache.NewIgnoreInventoryLimitCache,

			// domain services
			services.NewInventoryDomainService,
			services.NewTransactionDomainService,

			// application services
			appService.NewTransactionService,
			appService.NewInventoryService,

			// handlers
			api.NewTransactionsHandler,
			api.NewInventoryHandler,
			api.NewTransactionAdminHandler,
			api.NewInventoryAdminHandler,

			// server
			server.NewServer,
		),

		fx.Supply(),

		fx.Invoke(
			config.InitConfig,
			logger.SetupLogger,
			setupServer,
			db.Migrate,
			api.SetupTradingRoutes,
			api.SetupAdminRoutes,
			server.Run,
		),
	).Run()

}
