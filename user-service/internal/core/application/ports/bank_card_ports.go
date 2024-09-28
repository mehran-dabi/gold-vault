package ports

import (
	"context"

	"goldvault/user-service/internal/core/domain/entity"
)

type (
	BankCardPersistence interface {
		CreateBankCard(ctx context.Context, bankCard *entity.BankCard) error
		GetBankCardsByUserID(ctx context.Context, userID int64) ([]*entity.BankCard, error)
	}

	BankCardDomainService interface {
		AddUserBankCard(ctx context.Context, userID int64, cardNumber string) error
		GetUserBankCards(ctx context.Context, userID int64) ([]*entity.BankCard, error)
	}
)
