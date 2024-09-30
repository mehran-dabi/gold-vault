package persistence

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"goldvault/user-service/internal/core/application/ports"
	"goldvault/user-service/internal/core/domain/entity"
	"goldvault/user-service/internal/infrastructure/persistence/models"
	"goldvault/user-service/internal/infrastructure/persistence/queries"
	"goldvault/user-service/pkg/serr"
)

// UserPersistence is the persistence adapter for the User entity
type UserPersistence struct {
	db *sql.DB
}

// NewUserPersistence creates a new UserPersistence instance
func NewUserPersistence(db *sql.DB) ports.UserPersistencePorts {
	return &UserPersistence{db: db}
}

// SaveUser creates a new user in the database and updates the user with the generated ID and timestamps
func (p *UserPersistence) SaveUser(ctx context.Context, user *entity.User) error {
	dbModel, err := models.ToUserDB(user)
	if err != nil {
		return err
	}

	err = p.db.QueryRowContext(
		ctx,
		queries.CreateUser,
		dbModel.Phone,
		dbModel.FirstName,
		dbModel.LastName,
		dbModel.Role,
		dbModel.IsVerified,
		dbModel.NationalCode,
		dbModel.NationalCardImage,
		dbModel.Birthday,
	).Scan(&dbModel.ID, &dbModel.CreatedAt, &dbModel.UpdatedAt)
	if err != nil {
		return serr.DBError("SaveUser", "user_persistence_adapter", err)
	}

	// Update the entity after successful creation
	user.ID = dbModel.ID
	user.CreatedAt = dbModel.CreatedAt
	user.UpdatedAt = dbModel.UpdatedAt

	return nil
}

// FindUserByPhone retrieves a user by their phone number
func (p *UserPersistence) FindUserByPhone(ctx context.Context, phone string) (*entity.User, error) {
	row := p.db.QueryRowContext(ctx, queries.GetUserByPhone, phone)
	var user models.User
	err := user.Scan(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // User not found
		}
		return nil, serr.DBError("FindUserByPhone", "user_persistence_adapter", err)
	}
	userEntity, err := user.ToUserEntity()
	if err != nil {
		return nil, serr.DBError("FindUserByPhone", "user_persistence_adapter", err)
	}
	return userEntity, nil
}

func (p *UserPersistence) FindUserByID(ctx context.Context, userID int64) (*entity.User, error) {
	row := p.db.QueryRowContext(ctx, queries.GetUserByID, userID)
	var user models.User
	err := user.Scan(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // User not found
		}
		return nil, serr.DBError("FindUserByID", "user_persistence_adapter", err)
	}
	userEntity, err := user.ToUserEntity()
	if err != nil {
		return nil, serr.DBError("FindUserByID", "user_persistence_adapter", err)
	}
	return userEntity, nil
}

// UpdateUser updates a user in the database
func (p *UserPersistence) UpdateUser(ctx context.Context, user *entity.User) error {
	dbModel, err := models.ToUserDB(user)
	if err != nil {
		return err
	}

	result, err := p.db.ExecContext(
		ctx,
		queries.UpdateUser,
		dbModel.Phone,
		dbModel.FirstName,
		dbModel.LastName,
		dbModel.Role,
		dbModel.IsVerified,
		dbModel.NationalCode,
		dbModel.Birthday,
		dbModel.ID,
	)
	if err != nil {
		return serr.DBError("UpdateUser", "user_persistence_adapter", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return serr.DBError("UpdateUser", "user_persistence_adapter", err)
	}
	if rowsAffected == 0 {
		return serr.DBError("UpdateUser", "user_persistence_adapter", fmt.Errorf("no rows affected"))
	}

	return nil
}

func (p *UserPersistence) GetAllUsers(ctx context.Context, limit, offset int64) ([]*entity.User, error) {
	rows, err := p.db.QueryContext(ctx, queries.GetAllUsers, limit, offset)
	if err != nil {
		return nil, serr.DBError("GetAllUsers", "user_persistence_adapter", err)
	}
	defer rows.Close()

	users := make([]*entity.User, 0)
	for rows.Next() {
		var user models.User
		err := user.Scan(rows)
		if err != nil {
			return nil, serr.DBError("GetAllUsers", "user_persistence_adapter", err)
		}
		userEntity, err := user.ToUserEntity()
		if err != nil {
			return nil, serr.DBError("GetAllUsers", "user_persistence_adapter", err)
		}
		users = append(users, userEntity)
	}

	return users, nil
}

func (p *UserPersistence) UpdateNationalCardImage(ctx context.Context, userID int64, nationalCardImage string) error {
	_, err := p.db.ExecContext(ctx, queries.UpdateUserNationalCardImage, nationalCardImage, userID)
	if err != nil {
		return serr.DBError("UpdateNationalCardImage", "user_persistence_adapter", err)
	}
	return nil
}
