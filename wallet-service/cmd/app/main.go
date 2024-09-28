package main

import (
	"goldvault/wallet-service/internal/config"
	appService "goldvault/wallet-service/internal/core/application/services"
	"goldvault/wallet-service/internal/core/domain/services"
	"goldvault/wallet-service/internal/infrastructure/db"
	"goldvault/wallet-service/internal/infrastructure/persistence"
	"goldvault/wallet-service/internal/interfaces/api"
	"goldvault/wallet-service/internal/interfaces/grpc"
	"goldvault/wallet-service/internal/server"
	"goldvault/wallet-service/pkg/logger"

	"go.uber.org/fx"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	fx.New(
		fx.Provide(
			// external clients
			newAssetServiceClient,

			// postgres
			postgresDB,

			// redis
			redisCache,
			redisRateLimit,

			// persistence
			persistence.NewAssetPersistence,
			persistence.NewTransactionPersistence,
			persistence.NewWalletPersistence,

			// cache
			//cache.NewOTPCacheAdapter,

			// domain services
			services.NewAssetDomainService,
			services.NewTransactionDomainService,
			services.NewWalletDomainService,

			// application services
			appService.NewAssetService,
			appService.NewTransactionService,
			appService.NewWalletService,

			// handlers
			api.NewWalletHandler,
			api.NewAdminWalletHandler,
			api.NewAdminTransactionHandler,
			grpc.NewWalletGRPCHandler,
			grpc.NewAssetGRPCHandler,

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
			api.SetupWalletRoutes,
			api.SetupTransactionRoutes,
			server.Run,
		),
	).Run()

}
