package serr_test

import (
	"database/sql"
	"errors"
	"testing"

	"goldvault/user-service/pkg/serr"

	"github.com/stretchr/testify/assert"
)

func TestNewDBError(t *testing.T) {
	// DB errors' status must be 500 to avoid error leaking unless it's a not found error
	t.Run("internal error", func(t *testing.T) {
		t.Parallel()
		err := errors.New("random err")
		dbErr := serr.DBError("Test", "Test", err)
		assert.Equal(t, dbErr.(*serr.ServiceError).Code, 500)
	})
	t.Run("sql not found db error", func(t *testing.T) {
		t.Parallel()
		dbErr := serr.DBError("Test", "Test", sql.ErrNoRows)
		assert.Equal(t, dbErr.(*serr.ServiceError).Code, 404)
	})
}
