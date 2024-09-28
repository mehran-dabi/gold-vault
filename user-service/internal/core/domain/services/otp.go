package services

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"goldvault/user-service/internal/core/application/ports"
	"goldvault/user-service/internal/core/domain/entity"
)

// OTPService provides domain logic for OTP
type OTPService struct {
}

// NewOTPService creates a new instance of OTPService
func NewOTPService() ports.OTPDomainServicePorts {
	return &OTPService{}
}

// GenerateOTP generates a new OTP and saves it
func (o *OTPService) GenerateOTP(ctx context.Context) (string, error) {
	code, err := randomOTP()
	if err != nil {
		return "", err
	}

	return code, nil
}

// ValidateOTP validates an OTP
func (o *OTPService) ValidateOTP(ctx context.Context, otp entity.OTP) error {
	if err := otp.Validate(); err != nil {
		return err
	}

	return nil
}

// randomOTP generates a random 5-digit OTP
func randomOTP() (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(90000))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%05d", n.Int64()+10000), nil
}
