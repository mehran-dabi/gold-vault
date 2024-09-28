package services

import (
	"context"
	"errors"

	"goldvault/user-service/internal/core/application/ports"
	"goldvault/user-service/internal/core/domain/entity"
	api "goldvault/user-service/internal/interfaces/api/dto"

	"github.com/redis/go-redis/v9"
)

// OTPApplicationService provides application logic for OTPs.
type OTPApplicationService struct {
	otpDomainService     ports.OTPDomainServicePorts
	otpPersistence       ports.OTPCachePorts
	kavenegarSMSProvider ports.KavenegarSMSProviderPort
}

// NewOTPApplicationService creates a new instance of OTPApplicationService.
func NewOTPApplicationService(
	otpDomainService ports.OTPDomainServicePorts,
	otpPersistence ports.OTPCachePorts,
	kavenegarSMSProvider ports.KavenegarSMSProviderPort) *OTPApplicationService {
	return &OTPApplicationService{
		otpDomainService:     otpDomainService,
		otpPersistence:       otpPersistence,
		kavenegarSMSProvider: kavenegarSMSProvider,
	}
}

// GenerateAndSendOTP generates an OTP and sends it to the user's phone number.
func (o *OTPApplicationService) GenerateAndSendOTP(ctx context.Context, request *api.GenerateOTPRequest) error {
	code, err := o.otpDomainService.GenerateOTP(ctx)
	if err != nil {
		return err
	}

	err = o.otpPersistence.SaveOTP(ctx, request.PhoneNumber, code)
	if err != nil {
		return err
	}

	sms := entity.SimpleSMS{
		Receptor: request.PhoneNumber,
		Message:  code,
	}

	if err := o.kavenegarSMSProvider.SendSMS(ctx, sms); err != nil {
		return err
	}

	return nil
}

// ValidateOTP validates an OTP code for a phone number with the stored code.
func (o *OTPApplicationService) ValidateOTP(ctx context.Context, request *api.ValidateOTPRequest) (bool, error) {
	otp := entity.OTP{
		PhoneNumber: request.PhoneNumber,
		Code:        request.OTP,
	}
	if err := o.otpDomainService.ValidateOTP(ctx, otp); err != nil {
		return false, err
	}

	storedCode, err := o.otpPersistence.GetOTP(ctx, request.PhoneNumber)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, nil
		}
		return false, err
	}

	if storedCode != request.OTP {
		return false, nil
	}

	return true, nil
}
