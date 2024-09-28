package dto

type (
	DepositRequest struct {
		Amount float64 `json:"amount" validate:"required,gt=0"`
		Type   string  `json:"type" validate:"required"`
	}

	WithdrawRequest struct {
		Amount float64 `json:"amount" validate:"required,gt=0"`
		Type   string  `json:"type" validate:"required"`
	}
)

func (d *DepositRequest) Validate() error {
	return GetValidator().Struct(d)
}

func (w *WithdrawRequest) Validate() error {
	return GetValidator().Struct(w)
}
