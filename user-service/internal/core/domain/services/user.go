package services

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"goldvault/user-service/internal/core/application/ports"
	"goldvault/user-service/internal/core/domain/entity"
)

type UserService struct {
	userPersistence ports.UserPersistencePorts
	fileStorage     ports.FileStorage
}

func NewUserService(
	userPersistence ports.UserPersistencePorts,
	fileStorage ports.FileStorage,
) ports.UserDomainPorts {
	return &UserService{
		userPersistence: userPersistence,
		fileStorage:     fileStorage,
	}
}

func (u *UserService) CreateUser(ctx context.Context, phone string) (*entity.User, error) {
	userEntity := &entity.User{
		Phone:      phone,
		Role:       entity.RoleCustomer,
		IsVerified: true,
	}

	err := u.userPersistence.SaveUser(ctx, userEntity)
	if err != nil {
		return nil, err
	}

	return userEntity, nil
}

func (u *UserService) UpdateUser(ctx context.Context, updatedUser *entity.User) error {
	// Fetch existing user to ensure it exists
	existingUser, err := u.userPersistence.FindUserByID(ctx, updatedUser.ID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	if existingUser == nil {
		return fmt.Errorf("user not found")
	}

	// Update fields according to business logic
	existingUser.FirstName = updatedUser.FirstName
	existingUser.LastName = updatedUser.LastName
	existingUser.NationalCode = updatedUser.NationalCode
	existingUser.Birthday = updatedUser.Birthday
	existingUser.UpdatedAt = time.Now()

	// Persist the updated user
	if err := u.userPersistence.UpdateUser(ctx, existingUser); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (u *UserService) GetUsers(ctx context.Context, limit, offset int64) ([]*entity.User, error) {
	users, err := u.userPersistence.GetAllUsers(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	return users, nil
}

func (u *UserService) UploadNationalCard(ctx context.Context, file multipart.File, objectName string) error {
	bucketName := "national-card-ids"
	err := u.fileStorage.UploadFile(ctx, bucketName, objectName, file)
	if err != nil {
		return fmt.Errorf("failed to upload national card: %w", err)
	}

	return nil
}

func (u *UserService) UpdateNationalCardImage(ctx context.Context, userID int64, image string) error {
	err := u.userPersistence.UpdateNationalCardImage(ctx, userID, image)
	if err != nil {
		return fmt.Errorf("failed to update national card image: %w", err)
	}

	return nil
}
