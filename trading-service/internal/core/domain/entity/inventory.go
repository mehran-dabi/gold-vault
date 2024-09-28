package entity

import (
	"fmt"
	"time"
)

type Inventory struct {
	ID            int64
	AssetType     AssetType
	TotalQuantity float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (i *Inventory) Validate() error {
	if i.TotalQuantity < 0 {
		return fmt.Errorf("total quantity must be greater than or equal to 0")
	}
	return nil
}
