package middlewares

import (
	"context"
	"net/http"

	"goldvault/trading-service/internal/config"

	"github.com/gin-gonic/gin"
)

func RateLimitMiddleware(cfg *config.RateLimitConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Use IP address as the key
		key := cfg.KeyPrefix + c.ClientIP()

		// Increment the count and set expiration if it's the first request
		count, err := cfg.RedisClient.Incr(context.Background(), key).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		// Set expiration time on first request
		if count == 1 {
			cfg.RedisClient.Expire(context.Background(), key, cfg.Window)
		}

		// Check if the request count exceeds the limit
		if count > int64(cfg.Limit) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
			return
		}

		// Continue to the next handler
		c.Next()
	}
}
