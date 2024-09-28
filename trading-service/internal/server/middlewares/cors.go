package middlewares

import (
	"time"

	"goldvault/trading-service/internal/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     config.ServiceConfig.App.Cors.AllowOrigins,
		AllowMethods:     config.ServiceConfig.App.Cors.AllowMethods,
		AllowHeaders:     config.ServiceConfig.App.Cors.AllowHeaders,
		AllowCredentials: config.ServiceConfig.App.Cors.AllowCredentials,
		MaxAge:           12 * time.Hour,
	})
}
