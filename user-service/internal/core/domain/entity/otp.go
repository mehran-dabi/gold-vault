package entity

import (
	"time"

	"goldvault/user-service/pkg/serr"
)

type OTP struct {
	PhoneNumber string
	Code        string
	ExpiresAt   time.Time
}

func (o *OTP) Validate() error {
	if o.PhoneNumber == "" {
		return serr.ValidationErr(
			"OTP.Validate",
			"phone number cannot be empty",
			serr.ErrInvalidOTP,
		)
	}

	if o.Code == "" {
		return serr.ValidationErr(
			"OTP.Validate",
			"code cannot be empty",
			serr.ErrInvalidOTP,
		)
	}

	if len(o.Code) != 5 {
		return serr.ValidationErr(
			"OTP.Validate",
			"code must be 5 characters long",
			serr.ErrInvalidOTP,
		)
	}

	return nil
}

type SimpleSMS struct {
	Receptor string
	Message  string
	Sender   string
}
