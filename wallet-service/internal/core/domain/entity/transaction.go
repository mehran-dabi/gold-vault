package entity

import "time"

type Transaction struct {
	ID          int64     `json:"id"`
	WalletID    int64     `json:"wallet_id"`
	AssetID     int64     `json:"asset_id"`
	Type        TxType    `json:"type"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Status      TxStatus  `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type TxStatus string

const (
	TxPending TxStatus = "pending"
	TxSuccess TxStatus = "success"
	TxFailed  TxStatus = "failed"
)

func (t TxStatus) IsValid() bool {
	switch t {
	case TxPending, TxSuccess, TxFailed:
		return true
	}
	return false
}

type TxType string

const (
	TxTypeCredit TxType = "credit"
	TxTypeDebit  TxType = "debit"
)

func (t TxType) IsValid() bool {
	switch t {
	case TxTypeCredit, TxTypeDebit:
		return true
	}
	return false
}
