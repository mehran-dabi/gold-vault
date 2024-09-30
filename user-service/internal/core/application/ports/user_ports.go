package ports

import (
	"context"
	"mime/multipart"

	"goldvault/user-service/internal/core/domain/entity"
)

type (
	UserPersistencePorts interface {
		SaveUser(ctx context.Context, user *entity.User) error
		FindUserByPhone(ctx context.Context, phone string) (*entity.User, error)
		FindUserByID(ctx context.Context, userID int64) (*entity.User, error)
		UpdateUser(ctx context.Context, user *entity.User) error
		GetAllUsers(ctx context.Context, limit, offset int64) ([]*entity.User, error)
		UpdateNationalCardImage(ctx context.Context, userID int64, nationalCardImage string) error
	}

	UserDomainPorts interface {
		CreateUser(ctx context.Context, phone string) (*entity.User, error)
		UpdateUser(ctx context.Context, user *entity.User) error
		GetUsers(ctx context.Context, limit, offset int64) ([]*entity.User, error)
		UploadNationalCard(ctx context.Context, file multipart.File, objectName string) error
		UpdateNationalCardImage(ctx context.Context, userID int64, image string) error
	}
)
