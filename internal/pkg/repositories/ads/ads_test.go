package ads

import (
	"ads/internal/pkg/logger"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name FROM ads WHERE id = $1")).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow("1", "some"))

	var repo Repository = &repository{db: db, logger: logger.New([]string{})}

	ad, err := repo.GetByID("1")
	assert.NoError(t, err)

	assert.Equal(t, "some", ad.Name)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
