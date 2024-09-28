package ports

import (
	"context"

	"goldvault/user-service/internal/core/domain/entity"
)

type (
	// OTPCachePorts defines the methods for interacting with OTP data
	OTPCachePorts interface {
		// SaveOTP saves an OTP
		SaveOTP(ctx context.Context, phoneNumber, otp string) error

		// GetOTP retrieves an OTP
		GetOTP(ctx context.Context, phoneNumber string) (string, error)
	}

	// OTPDomainServicePorts defines the methods for interacting with OTP services
	OTPDomainServicePorts interface {
		// GenerateOTP generates an OTP
		GenerateOTP(ctx context.Context) (string, error)

		// ValidateOTP validates an OTP
		ValidateOTP(ctx context.Context, otp entity.OTP) error
	}
)
