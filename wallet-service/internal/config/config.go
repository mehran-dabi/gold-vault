package config

import (
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type Environment string

const (
	LOCAL Environment = "dev"
	BETA  Environment = "beta"
	PROD  Environment = "prod"
)

// Config holds all configuration for the application
type Config struct {
	Env      string         `mapstructure:"env"`
	Name     string         `mapstructure:"name"`
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"db"`
	Cache    CacheConfig    `mapstructure:"cache"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	App      AppConfig      `mapstructure:"app"`
	API      APIConfig      `mapstructure:"api"`
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Ports PortsConfig `mapstructure:"ports"`
	Debug bool        `mapstructure:"debug"`
}

// PortsConfig holds server port configurations
type PortsConfig struct {
	External string `mapstructure:"external"`
	Internal string `mapstructure:"internal"`
	GRPC     string `mapstructure:"grpc"`
}

// DatabaseConfig holds database-related configuration
type DatabaseConfig struct {
	Postgres PostgresConfig `mapstructure:"postgres"`
}

// PostgresConfig holds PostgreSQL-related configuration
type PostgresConfig struct {
	Name           string `mapstructure:"name"`
	Host           string `mapstructure:"host"`
	Port           string `mapstructure:"port"`
	User           string `mapstructure:"user"`
	Password       string `mapstructure:"password"`
	Debug          bool   `mapstructure:"debug"`
	MaxIdleConn    int    `mapstructure:"maxIdleConn"`
	MaxOpenConn    int    `mapstructure:"maxOpenConn"`
	MigrationsPath string `mapstructure:"migrationsPath"`
}

// CacheConfig holds cache-related configuration
type CacheConfig struct {
	Redis RedisConfig `mapstructure:"redis"`
}

// RedisConfig holds Redis-related configuration
type RedisConfig struct {
	Host            string          `mapstructure:"host"`
	Port            string          `mapstructure:"port"`
	Password        string          `mapstructure:"password"`
	CacheDB         int             `mapstructure:"cacheDB"`
	RateLimitDB     int             `mapstructure:"rateLimitDB"`
	PoolSize        int             `mapstructure:"poolSize"`
	MinIdleConns    int             `mapstructure:"minIdleConns"`
	MaxConnAge      time.Duration   `mapstructure:"maxConnAge"`
	PoolTimeout     time.Duration   `mapstructure:"poolTimeout"`
	IdleTimeout     time.Duration   `mapstructure:"idleTimeout"`
	ReadTimeout     time.Duration   `mapstructure:"readTimeout"`
	WriteTimeout    time.Duration   `mapstructure:"writeTimeout"`
	MaxRetries      int             `mapstructure:"maxRetries"`
	MinRetryBackoff time.Duration   `mapstructure:"minRetryBackoff"`
	MaxRetryBackoff time.Duration   `mapstructure:"maxRetryBackoff"`
	DialTimeout     time.Duration   `mapstructure:"dialTimeout"`
	Debug           bool            `mapstructure:"debug"`
	RateLimit       RateLimitConfig `mapstructure:"rateLimit"`
}

// RateLimitConfig holds rate limit configurations
type RateLimitConfig struct {
	RedisClient redis.UniversalClient
	Limit       int           `mapstructure:"limit"`
	Window      time.Duration `mapstructure:"window"`
	KeyPrefix   string        `mapstructure:"keyPrefix"`
}

// CacheRedisConfig holds the configuration for Redis cache.
type CacheRedisConfig struct {
	RedisClient redis.UniversalClient
}

// JWTConfig holds JWT-related configuration
type JWTConfig struct {
	SecretKey          string        `mapstructure:"secret-key"`
	AccessTokenExpiry  time.Duration `mapstructure:"access-token-expiry"`
	RefreshTokenExpiry time.Duration `mapstructure:"refresh-token-expiry"`
}

// AppConfig holds application-specific configuration
type AppConfig struct {
	Cors CorsConfig `mapstructure:"cors"`
}

// CorsConfig holds CORS configuration
type CorsConfig struct {
	AllowOrigins     []string `mapstructure:"allow-origins"`
	AllowMethods     []string `mapstructure:"allow-methods"`
	AllowHeaders     []string `mapstructure:"allow-headers"`
	AllowCredentials bool     `mapstructure:"allow-credentials"`
}

// APIConfig holds API-related configuration
type APIConfig struct {
	AssetServiceConfig AssetServiceConfig `mapstructure:"asset-service"`
}

// AssetServiceConfig holds asset service configuration
type AssetServiceConfig struct {
	GRPCAddress string `mapstructure:"grpc"`
}

var ServiceConfig Config

// InitConfig loads the configuration from file and environment variables
func InitConfig() {
	viper.SetConfigName(getEnv("CONFIG_NAME", "dev"))
	viper.AddConfigPath("/app/internal/config/env")
	if err := viper.MergeInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %w \n", err)
	}

	if err := viper.Unmarshal(&ServiceConfig); err != nil {
		log.Fatalf("Unable to unmarshal config: %v", err)
	}
}

func getEnv(key, fallback string) string {
	log.Println("getting environment")
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
