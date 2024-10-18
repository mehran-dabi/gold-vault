package entity

import (
	"fmt"
	"time"
)

type Transaction struct {
	ID              int64
	UserID          int64
	AssetType       string
	Quantity        float64
	Price           float64
	TransactionType TransactionType
	Status          TransactionStatus
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type TransactionType string

const (
	TransactionTypeBuy  TransactionType = "Buy"
	TransactionTypeSell TransactionType = "Sell"
)

func (t TransactionType) String() string {
	return string(t)
}

func (t TransactionType) IsValid() bool {
	switch t {
	case TransactionTypeBuy, TransactionTypeSell:
		return true
	}
	return false
}

type TransactionStatus string

const (
	TransactionStatusPending        TransactionStatus = "Pending"
	TransactionStatusCompleted      TransactionStatus = "Completed"
	TransactionStatusCancelled      TransactionStatus = "Cancelled"
	TransactionStatusBalancePending TransactionStatus = "BalancePending"
	TransactionStatusFailed         TransactionStatus = "Failed"
)

func (t TransactionStatus) String() string {
	return string(t)
}

func (t TransactionStatus) IsValid() bool {
	switch t {
	case TransactionStatusPending, TransactionStatusCompleted, TransactionStatusCancelled, TransactionStatusFailed, TransactionStatusBalancePending:
		return true
	}
	return false
}

func (t *Transaction) Validate() error {
	if t.UserID == 0 {
		return fmt.Errorf("invalid user id")
	}

	if t.AssetType == "" {
		return fmt.Errorf("invalid asset type")
	}

	if t.Quantity <= 0 {
		return fmt.Errorf("invalid quantity")
	}

	if t.Price <= 0 {
		return fmt.Errorf("invalid price")
	}

	if !t.TransactionType.IsValid() {
		return fmt.Errorf("invalid transaction type")
	}

	if !t.Status.IsValid() {
		return fmt.Errorf("invalid transaction status")
	}

	return nil
}

type TransactionsSummary struct {
	StartDate  time.Time
	EndDate    time.Time
	TotalSells float64
	TotalBuys  float64
}
