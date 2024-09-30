package services

import (
	"context"
	"fmt"

	"goldvault/user-service/internal/core/application/ports"
	"goldvault/user-service/internal/interfaces/api/dto"
	"goldvault/user-service/pkg/serr"
)

type AuthService struct {
	otpService        *OTPApplicationService
	jwtService        *JWTService
	userDomainService ports.UserDomainPorts
	userPersistence   ports.UserPersistencePorts
	walletClient      ports.WalletClientPorts
}

func NewAuthService(otpService *OTPApplicationService,
	jwtService *JWTService,
	userDomainService ports.UserDomainPorts,
	userPersistence ports.UserPersistencePorts,
	walletClient ports.WalletClientPorts) *AuthService {
	return &AuthService{
		otpService:        otpService,
		jwtService:        jwtService,
		userDomainService: userDomainService,
		userPersistence:   userPersistence,
		walletClient:      walletClient,
	}
}

// RequestOTP handles generating and sending an OTP for authentication
func (s *AuthService) RequestOTP(ctx context.Context, request *dto.GenerateOTPRequest) error {
	// validate the request
	if err := request.Validate(); err != nil {
		return serr.ValidationErr("AuthService.RequestOTP", err.Error(), serr.ErrInvalidInput)
	}

	err := s.otpService.GenerateAndSendOTP(ctx, request)
	if err != nil {
		return fmt.Errorf("failed to generate and send OTP: %w", err)
	}
	return nil
}

// VerifyOTPAndIssueToken validates the OTP and issues a JWT token
func (s *AuthService) VerifyOTPAndIssueToken(ctx context.Context, request *dto.ValidateOTPRequest) (string, error) {
	// Validate the request
	if err := request.Validate(); err != nil {
		return "", serr.ValidationErr("AuthService.VerifyOTPAndIssueToken", err.Error(), serr.ErrInvalidInput)
	}

	// Validate the OTP
	//isValid, err := s.otpService.ValidateOTP(ctx, request)
	//if err != nil {
	//	return "", fmt.Errorf("failed to validate OTP: %w", err)
	//}
	//if !isValid {
	//	return "", fmt.Errorf("invalid or expired OTP")
	//}

	// Retrieve the user from the database
	user, err := s.userPersistence.FindUserByPhone(ctx, request.PhoneNumber)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve user: %w", err)
	}

	if user == nil {
		newUser, err := s.userDomainService.CreateUser(ctx, request.PhoneNumber)
		if err != nil {
			return "", fmt.Errorf("failed to create user: %w", err)
		}

		// Create a wallet for the user
		_, err = s.walletClient.CreateWallet(ctx, newUser.ID)
		if err != nil {
			return "", fmt.Errorf("failed to create wallet: %w", err)
		}

		user = newUser
	}

	// Generate a JWT token for the user
	token, err := s.jwtService.GenerateToken(user.ID, user.Role.String())
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT: %w", err)
	}

	return token, nil
}
