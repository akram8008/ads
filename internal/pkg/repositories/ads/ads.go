package ads

import (
	"ads/internal/model"
	"ads/internal/pkg/db"
	"ads/internal/pkg/logger"
	"database/sql"
)

type Repository interface {
	Create()
	GetAll()
	GetByID(id uint64) (model.Ads, error)
}

type repository struct {
	db     *sql.DB
	logger logger.Logger
}

type Params struct {
	Logger logger.Logger
}

func New(p Params) Repository {
	return &repository{
		db:     db.New(db.Params{Logger: p.Logger}).Connection(),
		logger: p.Logger,
	}
}

// Create creates ad
func (r *repository) Create() {

}

// GetAll gets all ads
func (r *repository) GetAll() {

}

// GetByID gets ad by its ID
func (r *repository) GetByID(id uint64) (model.Ads, error) {
	query := `
		SELECT *
		FROM ads
		WHERE id = $1
	`

	var ads model.Ads

	err := r.db.QueryRow(query, id).Scan(&ads)
	if err != nil {
		r.logger.Logger().Error(err)
		return model.Ads{}, err
	}

	r.logger.Logger().Infof("[Database] Got ads with id: %d values: %v", id, ads)

	return ads, nil
}
