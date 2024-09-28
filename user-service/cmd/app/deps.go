package main

import (
	"database/sql"
	"log"

	"goldvault/user-service/cmd/clients"
	"goldvault/user-service/internal/config"
	"goldvault/user-service/internal/core/application/ports"
	"goldvault/user-service/internal/infrastructure/cache"
	"goldvault/user-service/internal/infrastructure/db"
	"goldvault/user-service/internal/server"
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

func kavenegarClient() ports.KavenegarSMSProviderPort {
	client := clients.NewKavenegarSMSProvider(
		config.ServiceConfig.API.Kavenegar.ApiKey,
		config.ServiceConfig.API.Kavenegar.Sender,
		config.ServiceConfig.API.Kavenegar.Host,
	)

	return client
}

func walletServiceClient() ports.WalletClientPorts {
	client, err := clients.NewWalletClient(config.ServiceConfig.API.Wallet.GRPC)
	if err != nil {
		log.Fatalf("failed to initalize wallet service client: %v", err)
	}

	return client
}

func setupServer(s *server.Server, psql *sql.DB) {
	s.SetHealthFunc(healthFunc(psql)).SetupRoutes()
}
