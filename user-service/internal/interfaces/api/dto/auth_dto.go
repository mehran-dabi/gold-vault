package dto

type (
	GenerateOTPRequest struct {
		PhoneNumber string `json:"phone_number" validate:"required,iranphone"`
	}

	ValidateOTPRequest struct {
		PhoneNumber string `json:"phone_number" validate:"required,iranphone"`
		OTP         string `json:"otp" validate:"required,len=5"`
	}
)

func (g *GenerateOTPRequest) Validate() error {
	return GetValidator().Struct(g)
}

func (v *ValidateOTPRequest) Validate() error {
	return GetValidator().Struct(v)
}
