package dto

type (
	AddUserBankCardRequest struct {
		CardNumber string `json:"card_number" validate:"required,len=16"`
	}
)

func (a *AddUserBankCardRequest) Validate() error {
	return validate.Struct(a)
}
