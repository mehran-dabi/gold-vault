package queries

import "goldvault/user-service/internal/infrastructure/persistence/models"

const (
	// CreateUser is a query to create a new user in the database and returns the id, created_at, and updated_at fields.
	CreateUser = `
        INSERT INTO ` + models.UsersTableName + ` (` + models.UserColumnsNoID + `)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())
        RETURNING id, created_at, updated_at;
    `

	// GetUserByPhone is a query to get a user by their phone number.
	GetUserByPhone = `
        SELECT ` + models.UserColumns + `
        FROM ` + models.UsersTableName + `
        WHERE phone = $1;
    `

	// GetUserByID is a query to get a user by their ID.
	GetUserByID = `
		SELECT ` + models.UserColumns + `
		FROM ` + models.UsersTableName + `
		WHERE id = $1;
	`
	// UpdateUser is a query to update a user's name and verification status.
	UpdateUser = `
        UPDATE ` + models.UsersTableName + `
        SET phone = $1, first_name = $2, last_name = $3, role = $4, is_verified = $5, national_code = $6, birthday = $7, updated_at = NOW()
        WHERE id = $8;
    `

	// GetAllUsers is a query to get all users with pagination.
	GetAllUsers = `
		SELECT ` + models.UserColumns + `
		FROM ` + models.UsersTableName + `
		LIMIT $1 OFFSET $2;
	`

	// UpdateUserNationalCardImage is a query to update a user's national card image.
	UpdateUserNationalCardImage = `
		UPDATE ` + models.UsersTableName + `
		SET national_card_image = $1, updated_at = NOW()
		WHERE id = $2;
	`
)
