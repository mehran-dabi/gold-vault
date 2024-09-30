package main

import (
	"goldvault/user-service/internal/config"
	appService "goldvault/user-service/internal/core/application/services"
	"goldvault/user-service/internal/core/domain/services"
	"goldvault/user-service/internal/infrastructure/cache"
	"goldvault/user-service/internal/infrastructure/db"
	"goldvault/user-service/internal/infrastructure/persistence"
	"goldvault/user-service/internal/infrastructure/storage"
	"goldvault/user-service/internal/interfaces/api"
	"goldvault/user-service/internal/server"
	"goldvault/user-service/pkg/logger"

	"go.uber.org/fx"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	fx.New(
		fx.Provide(
			// external clients
			kavenegarClient,
			walletServiceClient,

			// postgres
			postgresDB,

			// redis
			redisCache,
			redisRateLimit,

			// minio
			minioStorage,

			// storage
			storage.NewFileStorage,

			// persistence
			persistence.NewUserPersistence,
			persistence.NewBankCardPersistence,

			// cache
			cache.NewOTPCacheAdapter,

			// domain services
			services.NewOTPService,
			services.NewUserService,
			services.NewBankCardDomainService,

			// application services
			appService.NewJWTService,
			appService.NewOTPApplicationService,
			appService.NewAuthService,
			appService.NewUserService,
			appService.NewBankCardService,

			// handlers
			api.NewAuthHandler,
			api.NewUserHandler,
			api.NewAdminUserHandler,
			api.NewBankCardHandler,

			// server
			server.NewServer,
		),

		fx.Supply(),

		fx.Invoke(
			config.InitConfig,
			logger.SetupLogger,
			setupServer,
			db.Migrate,
			api.SetupAuthRoutes,
			api.SetupUserRoutes,
			server.Run,
		),
	).Run()

}
