package models

import (
	"fmt"
	"time"

	"goldvault/user-service/internal/core/domain/entity"
)

type User struct {
	ID                int64
	Phone             string
	FirstName         string
	LastName          string
	Role              string
	IsVerified        bool
	NationalCode      string
	NationalCardImage string
	Birthday          time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

const (
	UsersTableName  = "users"
	UserColumns     = "id, phone, first_name, last_name, role, is_verified, national_code, national_card_image, birthday, created_at, updated_at"
	UserColumnsNoID = "phone, first_name, last_name, role, is_verified, national_code, national_card_image, birthday, created_at, updated_at"
)

func (u *User) Scan(scanner Scanner) error {
	return scanner.Scan(
		&u.ID,
		&u.Phone,
		&u.FirstName,
		&u.LastName,
		&u.Role,
		&u.IsVerified,
		&u.NationalCode,
		&u.NationalCardImage,
		&u.Birthday,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
}

func ToUserDB(entity *entity.User) (*User, error) {
	if !entity.Role.IsValid() {
		return nil, fmt.Errorf("invalid role: %s", entity.Role)
	}

	return &User{
		ID:                entity.ID,
		Phone:             entity.Phone,
		FirstName:         entity.FirstName,
		LastName:          entity.LastName,
		Birthday:          entity.Birthday,
		Role:              entity.Role.String(),
		NationalCode:      entity.NationalCode,
		NationalCardImage: entity.NationalCardImage,
		IsVerified:        entity.IsVerified,
		CreatedAt:         entity.CreatedAt,
		UpdatedAt:         entity.UpdatedAt,
	}, nil
}

func (u *User) ToUserEntity() (*entity.User, error) {
	role := entity.Roles(u.Role)
	if !role.IsValid() {
		return nil, fmt.Errorf("invalid role: %s", u.Role)
	}

	return &entity.User{
		ID:                u.ID,
		Phone:             u.Phone,
		FirstName:         u.FirstName,
		LastName:          u.LastName,
		Birthday:          u.Birthday,
		Role:              entity.Roles(u.Role),
		NationalCode:      u.NationalCode,
		NationalCardImage: u.NationalCardImage,
		IsVerified:        u.IsVerified,
		CreatedAt:         u.CreatedAt,
		UpdatedAt:         u.UpdatedAt,
	}, nil
}
