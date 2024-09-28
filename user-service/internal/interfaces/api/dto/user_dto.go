package dto

import "time"

type (
	UpdateUserRequest struct {
		ID           int64     `json:"-"`
		FirstName    string    `json:"first_name"`
		LastName     string    `json:"last_name"`
		NationalCode string    `json:"national_code" validate:"len=10"`
		Birthday     time.Time `json:"birthday"`
	}

	AdminUpdateUserRequest struct {
		ID           int64     `json:"-"`
		FirstName    string    `json:"first_name"`
		LastName     string    `json:"last_name"`
		NationalCode string    `json:"national_code" validate:"len=10"`
		Birthday     time.Time `json:"birthday"`
		Role         string    `json:"role"`
		Phone        string    `json:"phone"`
		IsVerified   bool      `json:"is_verified"`
	}
)

func (u *UpdateUserRequest) Validate() error {
	return GetValidator().Struct(u)
}

func (u *AdminUpdateUserRequest) Validate() error {
	return GetValidator().Struct(u)
}
