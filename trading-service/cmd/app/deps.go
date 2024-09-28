package main

import (
	"database/sql"
	"log"

	"goldvault/trading-service/cmd/clients"
	"goldvault/trading-service/internal/config"
	"goldvault/trading-service/internal/core/application/ports"
	"goldvault/trading-service/internal/infrastructure/cache"
	"goldvault/trading-service/internal/infrastructure/db"
	"goldvault/trading-service/internal/server"
)

func postgresDB() *sql.DB {
	psql, err := db.NewPostgres(
		config.ServiceConfig.Database.Postgres.Name,
		config.ServiceConfig.Database.Postgres.User,
		config.ServiceConfig.Database.Postgres.Password,
		config.ServiceConfig.Database.Postgres.Host,
		config.ServiceConfig.Database.Postgres.Port,
		config.ServiceConfig.Database.Postgres.MaxOpenConn,
		config.ServiceConfig.Database.Postgres.MaxIdleConn,
	)
	if err != nil {
		log.Fatalf("failed to initalize db: %v", err)
	}
	return psql
}

func redisCache() *config.CacheRedisConfig {
	redisClient, err := cache.NewRedisClient(
		config.ServiceConfig.Cache.Redis.Host,
		config.ServiceConfig.Cache.Redis.Port,
		config.ServiceConfig.Cache.Redis.Password,
		config.ServiceConfig.Cache.Redis.CacheDB,
	)
	if err != nil {
		log.Fatalf("failed to initalize cache redis: %v", err)
	}
	return &config.CacheRedisConfig{RedisClient: redisClient}
}

func redisRateLimit() *config.RateLimitConfig {
	redisClient, err := cache.NewRedisClient(
		config.ServiceConfig.Cache.Redis.Host,
		config.ServiceConfig.Cache.Redis.Port,
		config.ServiceConfig.Cache.Redis.Password,
		config.ServiceConfig.Cache.Redis.RateLimitDB,
	)
	if err != nil {
		log.Fatalf("failed to initalize rate limit redis: %v", err)
	}

	return &config.RateLimitConfig{
		RedisClient: redisClient,
		Limit:       config.ServiceConfig.Cache.Redis.RateLimit.Limit,
		Window:      config.ServiceConfig.Cache.Redis.RateLimit.Window,
		KeyPrefix:   config.ServiceConfig.Cache.Redis.RateLimit.KeyPrefix,
	}
}

func assetService() ports.AssetServiceClient {
	client, err := clients.NewAssetServiceClient(config.ServiceConfig.API.AssetService.GRPC)
	if err != nil {
		log.Fatalf("failed to initalize asset service client: %v", err)
	}
	return client
}

func walletService() ports.WalletServiceClient {
	client, err := clients.NewWalletServiceClient(config.ServiceConfig.API.WalletService.GRPC)
	if err != nil {
		log.Fatalf("failed to initalize wallet service client: %v", err)
	}
	return client
}

func setupServer(s *server.Server, psql *sql.DB) {
	s.SetHealthFunc(healthFunc(psql)).SetupRoutes()
}
