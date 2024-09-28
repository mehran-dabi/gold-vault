package cache

import (
	"context"
	"time"

	"goldvault/user-service/internal/config"
	"goldvault/user-service/internal/core/application/ports"
)

// OTPCacheAdapter provides persistence operations for OTPs
type OTPCacheAdapter struct {
	cacheRedisClient *config.CacheRedisConfig
}

// NewOTPCacheAdapter creates a new instance of OTPCacheAdapter
func NewOTPCacheAdapter(cacheRedisClient *config.CacheRedisConfig) ports.OTPCachePorts {
	return &OTPCacheAdapter{cacheRedisClient: cacheRedisClient}
}

// SaveOTP saves an OTP
func (r *OTPCacheAdapter) SaveOTP(ctx context.Context, phoneNumber string, otp string) error {
	return r.cacheRedisClient.RedisClient.Set(ctx, phoneNumber, otp, 5*time.Minute).Err()
}

// GetOTP retrieves an OTP
func (r *OTPCacheAdapter) GetOTP(ctx context.Context, phoneNumber string) (string, error) {
	return r.cacheRedisClient.RedisClient.Get(ctx, phoneNumber).Result()
}
