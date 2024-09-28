package ports

import (
	"context"

	"goldvault/user-service/internal/core/domain/entity"
)

type KavenegarSMSProviderPort interface {
	SendSMS(ctx context.Context, sms entity.SimpleSMS) error
}
