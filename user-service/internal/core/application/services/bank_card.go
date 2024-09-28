package services

import (
	"context"
	"net/http"

	"goldvault/user-service/internal/core/application/ports"
	"goldvault/user-service/internal/core/domain/entity"
	"goldvault/user-service/pkg/serr"
)

type BankCardService struct {
	bankCardDomainService ports.BankCardDomainService
}

func NewBankCardService(bankCardDomainService ports.BankCardDomainService) *BankCardService {
	return &BankCardService{
		bankCardDomainService: bankCardDomainService,
	}
}

func (b *BankCardService) AddUserBankCard(ctx context.Context, userID int64, cardNumber string) error {
	err := b.bankCardDomainService.AddUserBankCard(ctx, userID, cardNumber)
	if err != nil {
		return serr.ServiceErr("AddUserBankCard", err.Error(), err, http.StatusInternalServerError)
	}

	return nil
}

func (b *BankCardService) GetUserBankCards(ctx context.Context, userID int64) ([]*entity.BankCard, error) {
	bankCards, err := b.bankCardDomainService.GetUserBankCards(ctx, userID)
	if err != nil {
		return nil, serr.ServiceErr("GetUserBankCards", err.Error(), err, http.StatusInternalServerError)
	}

	return bankCards, nil
}
