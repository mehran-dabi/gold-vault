package ports

import (
	"context"

	"goldvault/user-service/internal/core/domain/entity"
)

type (
	UserPersistencePorts interface {
		SaveUser(ctx context.Context, user *entity.User) error
		FindUserByPhone(ctx context.Context, phone string) (*entity.User, error)
		FindUserByID(ctx context.Context, userID int64) (*entity.User, error)
		UpdateUser(ctx context.Context, user *entity.User) error
		GetAllUsers(ctx context.Context, limit, offset int64) ([]*entity.User, error)
	}

	UserDomainPorts interface {
		CreateUser(ctx context.Context, phone string) (*entity.User, error)
		UpdateUser(ctx context.Context, user *entity.User) error
		GetUsers(ctx context.Context, limit, offset int64) ([]*entity.User, error)
	}
)
