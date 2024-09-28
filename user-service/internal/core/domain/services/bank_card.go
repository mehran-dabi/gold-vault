package services

import (
	"context"

	"goldvault/user-service/internal/core/application/ports"
	"goldvault/user-service/internal/core/domain/entity"
)

type BankCardDomainService struct {
	bankCardPersistence ports.BankCardPersistence
}

func NewBankCardDomainService(bankCardPersistence ports.BankCardPersistence) ports.BankCardDomainService {
	return &BankCardDomainService{
		bankCardPersistence: bankCardPersistence,
	}
}

func (b *BankCardDomainService) AddUserBankCard(ctx context.Context, userID int64, cardNumber string) error {
	bankCard := &entity.BankCard{
		UserID:     userID,
		CardNumber: cardNumber,
	}

	err := b.bankCardPersistence.CreateBankCard(ctx, bankCard)
	if err != nil {
		return err
	}

	return nil
}

func (b *BankCardDomainService) GetUserBankCards(ctx context.Context, userID int64) ([]*entity.BankCard, error) {
	bankCards, err := b.bankCardPersistence.GetBankCardsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return bankCards, nil
}
