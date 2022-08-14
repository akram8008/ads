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
	GetByID(id string) (model.Ads, error)
}

type repository struct {
	db     *sql.DB
	logger logger.Logger
}

func New() Repository {
	return &repository{
		db: db.New().Connection(),
	}
}

// Create creates ad
func (r *repository) Create() {

}

// GetAll gets all ads
func (r *repository) GetAll() {

}

// GetByID gets ad by its ID
func (r *repository) GetByID(id string) (model.Ads, error) {
	query := `
		SELECT id, name
		FROM ads
		WHERE id = $1
	`

	var ad model.Ads

	err := r.db.QueryRow(query, id).Scan(&ad.ID, &ad.Name)
	if err != nil {
		r.logger.Logger().Error(err)
		return model.Ads{}, err
	}

	r.logger.Logger().Infoln("got it")

	return ad, nil
}
