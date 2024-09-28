package entity

import (
	"fmt"
	"time"
)

type AssetType string

const (
	AssetTypeMesghal        AssetType = "mesghal"
	AssetTypeSekke          AssetType = "sekke"
	AssetTypeSekkeBahar     AssetType = "sekke_bahar"
	AssetTypeNimSekke       AssetType = "nim_sekke"
	AssetTypeRob            AssetType = "rob"
	AssetTypeGerami         AssetType = "gerami"
	AssetTypeOther          AssetType = "other"
	AssetTypePose           AssetType = "pose"
	AssetTypeNimUnder80To85 AssetType = "nim_under_80_to_85"
	AssetTypeRobUnder80To85 AssetType = "rob_under_80_to_85"
	AssetTypeParsian100     AssetType = "parsian_100"
	AssetTypeParsian200     AssetType = "parsian_200"
	AssetTypeParsian300     AssetType = "parsian_300"
	AssetTypeParsian400     AssetType = "parsian_400"
	AssetTypeParsian500     AssetType = "parsian_500"
	AssetTypeParsian600     AssetType = "parsian_600"
	AssetTypeParsian700     AssetType = "parsian_700"
	AssetTypeParsian800     AssetType = "parsian_800"
	AssetTypeParsian900     AssetType = "parsian_900"
	AssetTypeParsian1000    AssetType = "parsian_1000"
	AssetTypeIRR            AssetType = "IRR"
)

var AssetTypes = []AssetType{
	AssetTypeMesghal,
	AssetTypeSekke,
	AssetTypeSekkeBahar,
	AssetTypeNimSekke,
	AssetTypeRob,
	AssetTypeGerami,
	AssetTypeOther,
	AssetTypePose,
	AssetTypeNimUnder80To85,
	AssetTypeRobUnder80To85,
	AssetTypeParsian100,
	AssetTypeParsian200,
	AssetTypeParsian300,
	AssetTypeParsian400,
	AssetTypeParsian500,
	AssetTypeParsian600,
	AssetTypeParsian700,
	AssetTypeParsian800,
	AssetTypeParsian900,
	AssetTypeParsian1000,
	AssetTypeIRR,
}

func (a AssetType) Validate() error {
	for _, at := range AssetTypes {
		if a == at {
			return nil
		}
	}

	return fmt.Errorf("invalid asset type")
}

func (a AssetType) String() string {
	return string(a)
}

type AssetPrice struct {
	ID        int64
	AssetType AssetType
	Prices    PriceDetails
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PriceDetails struct {
	BuyPrice  float64
	SellPrice float64
}

func (a *AssetPrice) Validate() error {
	if a.AssetType == "" {
		return fmt.Errorf("asset type must not be empty")
	}

	if a.AssetType.Validate() != nil {
		return fmt.Errorf("invalid asset type")
	}

	if a.Prices.BuyPrice < 0 {
		return fmt.Errorf("buy price must be greater than or equal to 0")
	}

	if a.Prices.SellPrice < 0 {
		return fmt.Errorf("sell price must be greater than or equal to 0")
	}

	return nil
}

func IsValidAssetType(assetType string) bool {
	for _, at := range AssetTypes {
		if at.String() == assetType {
			return true
		}
	}

	return false
}
