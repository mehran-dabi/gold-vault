package entity

import "time"

type BankCard struct {
	ID         int64
	UserID     int64
	CardNumber string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
