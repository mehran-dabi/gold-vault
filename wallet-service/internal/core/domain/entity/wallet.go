package entity

import "time"

type Wallet struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Assets    []Asset   `json:"assets"` // Holds multiple types of assets
}
