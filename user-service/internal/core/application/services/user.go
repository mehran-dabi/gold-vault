package services

import (
	"context"
	"fmt"

	"goldvault/user-service/internal/core/application/ports"
	"goldvault/user-service/internal/core/domain/entity"
	"goldvault/user-service/internal/interfaces/api/dto"
	"goldvault/user-service/pkg/serr"
)

type UserService struct {
	userPersistence ports.UserPersistencePorts
	userDomain      ports.UserDomainPorts
}

func NewUserService(userPersistence ports.UserPersistencePorts,
	userDomain ports.UserDomainPorts) *UserService {
	return &UserService{
		userPersistence: userPersistence,
		userDomain:      userDomain,
	}
}

func (u *UserService) GetProfile(ctx context.Context, userID int64) (*entity.User, error) {
	user, err := u.userPersistence.FindUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user profile: %w", err)
	}

	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (u *UserService) UpdateUser(ctx context.Context, updatedUser *dto.UpdateUserRequest) error {
	// Validate the updated user
	if err := updatedUser.Validate(); err != nil {
		return serr.ValidationErr("UserService.UpdateUser", err.Error(), serr.ErrInvalidInput)
	}

	// Get the user from the database
	user, err := u.userPersistence.FindUserByID(ctx, updatedUser.ID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	if user == nil {
		return fmt.Errorf("user not found")
	}

	// Convert DTO to domain entity
	userEntity := &entity.User{
		ID:           updatedUser.ID,
		FirstName:    updatedUser.FirstName,
		LastName:     updatedUser.LastName,
		NationalCode: updatedUser.NationalCode,
		Birthday:     updatedUser.Birthday,
	}

	// Delegate the update operation to the domain service
	return u.userDomain.UpdateUser(ctx, userEntity)
}

func (u *UserService) AdminUpdateUser(ctx context.Context, request *dto.AdminUpdateUserRequest) error {
	// Validate the updated user
	if err := request.Validate(); err != nil {
		return serr.ValidationErr("UserService.AdminUpdateUser", err.Error(), serr.ErrInvalidInput)
	}

	// Get the user from the database
	user, err := u.userPersistence.FindUserByID(ctx, request.ID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	if user == nil {
		return fmt.Errorf("user not found")
	}

	// Convert DTO to domain entity
	userEntity := &entity.User{
		ID:           request.ID,
		FirstName:    request.FirstName,
		LastName:     request.LastName,
		NationalCode: request.NationalCode,
		Birthday:     request.Birthday,
		Role:         entity.Roles(request.Role),
		Phone:        request.Phone,
		IsVerified:   request.IsVerified,
	}

	// Validate the updated user
	if err := userEntity.Validate(); err != nil {
		return serr.ValidationErr("UserService.AdminUpdateUser", err.Error(), serr.ErrInvalidInput)
	}

	// Delegate the update operation to the domain service
	return u.userDomain.UpdateUser(ctx, userEntity)
}

func (u *UserService) GetUsers(ctx context.Context, limit, offset int64) ([]*entity.User, error) {
	return u.userDomain.GetUsers(ctx, limit, offset)
}
