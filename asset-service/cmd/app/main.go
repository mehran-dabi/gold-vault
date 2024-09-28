package main

import (
	"goldvault/asset-service/internal/config"
	appService "goldvault/asset-service/internal/core/application/services"
	"goldvault/asset-service/internal/core/domain/services"
	"goldvault/asset-service/internal/infrastructure/cache"
	"goldvault/asset-service/internal/infrastructure/db"
	"goldvault/asset-service/internal/infrastructure/persistence"
	"goldvault/asset-service/internal/interfaces/api"
	"goldvault/asset-service/internal/interfaces/grpc"
	"goldvault/asset-service/internal/server"
	"goldvault/asset-service/pkg/logger"

	"go.uber.org/fx"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	fx.New(
		fx.Provide(
			// external clients
			//

			// postgres
			postgresDB,

			// redis
			redisCache,
			redisRateLimit,

			// persistence
			persistence.NewAssetPricePersistence,
			persistence.NewPriceHistoryPersistence,

			// cache
			cache.NewPriceCache,

			// domain services
			services.NewAssetPriceDomainService,
			services.NewPriceHistoryDomainService,

			// application services
			appService.NewAssetPriceService,
			appService.NewPriceHistoryService,

			// handlers
			api.NewAssetPriceHandler,
			api.NewAdminAssetPriceHandler,
			api.NewPriceHistoryHandler,
			grpc.NewAssetPriceGRPCHandler,

			// server
			server.NewGRPCServer,
			server.NewGRPCListener,
			server.NewServer,
		),

		fx.Supply(),

		fx.Invoke(
			config.InitConfig,
			logger.SetupLogger,
			setupServer,
			db.Migrate,
			api.SetupAssetPriceRoutes,
			api.SetupPriceHistoryRoutes,
			server.Run,
		),
	).Run()

}
