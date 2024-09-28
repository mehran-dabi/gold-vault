package main

import (
	"database/sql"
	"log"

	"goldvault/wallet-service/cmd/clients"
	"goldvault/wallet-service/internal/config"
	"goldvault/wallet-service/internal/core/application/ports"
	"goldvault/wallet-service/internal/infrastructure/cache"
	"goldvault/wallet-service/internal/infrastructure/db"
	"goldvault/wallet-service/internal/server"
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

func newAssetServiceClient() ports.AssetServiceClientPorts {
	client, err := clients.NewAssetServiceClient(config.ServiceConfig.API.AssetServiceConfig.GRPCAddress)
	if err != nil {
		log.Fatalf("failed to create asset service client: %v", err)
	}
	return client
}

func setupServer(s *server.Server, psql *sql.DB) {
	s.SetHealthFunc(healthFunc(psql)).SetupRoutes()
}
