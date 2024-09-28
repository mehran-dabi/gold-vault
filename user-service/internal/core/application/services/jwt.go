package services

import (
	"time"

	"goldvault/user-service/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	SecretKey          string
	AccessTokenExpiry  time.Duration
	RefreshTokenExpire time.Duration
}

func NewJWTService() *JWTService {
	return &JWTService{
		SecretKey:          config.ServiceConfig.JWT.SecretKey,
		AccessTokenExpiry:  config.ServiceConfig.JWT.AccessTokenExpiry,
		RefreshTokenExpire: config.ServiceConfig.JWT.RefreshTokenExpiry,
	}
}

type Claims struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func (s *JWTService) GenerateToken(userID int64, role string) (string, error) {
	claims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.AccessTokenExpiry)),
			Issuer:    "user-service",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.SecretKey))
}

func (s *JWTService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
