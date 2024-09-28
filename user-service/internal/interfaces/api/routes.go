package api

import (
	"goldvault/user-service/internal/config"
	"goldvault/user-service/internal/server"
	"goldvault/user-service/internal/server/middlewares"
)

func SetupAuthRoutes(s *server.Server, o *AuthHandler, rateLimitConfig *config.RateLimitConfig) {
	auth := s.External.Group("/auth")
	{
		auth.POST("/otp", middlewares.RateLimitMiddleware(rateLimitConfig), o.GenerateOTP)
		auth.POST("/otp/validate", o.ValidateOTP)
	}
}

func SetupUserRoutes(s *server.Server, u *UserHandler, au *AdminUserHandler, b *BankCardHandler) {
	users := s.External.Group("/users")
	users.Use(middlewares.JWTMiddleware(), middlewares.RoleMiddleware("customer"))
	{
		users.GET("/me", u.GetProfile)
		users.PATCH("/me", u.UpdateProfile)
		bankCard := users.Group("/bank-cards")
		{
			bankCard.POST("", b.AddUserBankCard)
			bankCard.GET("", b.GetUserBankCards)
		}
	}

	admin := s.External.Group("/admin")
	admin.Use(middlewares.JWTMiddleware(), middlewares.RoleMiddleware("admin"))
	{
		admin.GET("/users", au.GetUsers)
		admin.GET("/users/:id", au.GetProfile)
		admin.PATCH("/users/:id", au.UpdateProfile)
	}
}
