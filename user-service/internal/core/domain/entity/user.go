package entity

import (
	"errors"
	"regexp"
	"time"
)

// User represents the user entity
type User struct {
	ID                int64
	Phone             string
	FirstName         string
	LastName          string
	NationalCode      string
	NationalCardImage string
	Birthday          time.Time
	Role              Roles
	IsVerified        bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

// Validate validates the User entity fields
func (u *User) Validate() error {
	if !isValidPhoneNumber(u.Phone) {
		return errors.New("invalid phone number format")
	}
	if !isValidNationalCode(u.NationalCode) {
		return errors.New("invalid national code format")
	}
	if !u.Role.IsValid() {
		return errors.New("invalid role")
	}

	return nil
}

// IsValidPhoneNumber checks if the phone number is in a valid format
func isValidPhoneNumber(phone string) bool {
	re := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`) // Basic E.164 format check
	return re.MatchString(phone)
}

// IsValidNationalCode checks if the national code is in a valid format
func isValidNationalCode(nationalCode string) bool {
	re := regexp.MustCompile(`^\d{10}$`) // Ensure national code is exactly 10 digits
	return re.MatchString(nationalCode)
}
