package ads

import (
	"ads/internal/pkg/db"
	"ads/internal/pkg/logger"
	"ads/internal/pkg/models"
	"gorm.io/gorm"
)

type Repository interface {
	Create(models.Ads) (uint64, error)
	GetAll(int, string, string) ([]models.Ads, error)
	GetByID(uint64) (models.Ads, error)
}

type repository struct {
	db     *gorm.DB
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
func (r *repository) Create(ads models.Ads) (uint64, error) {
	err := r.db.Create(&ads).Error
	return ads.ID, err
}

// GetAll gets all ads
func (r *repository) GetAll(page int, price, createdDate string) ([]models.Ads, error) {

	query := r.db.Model(&models.Ads{}).Offset((page - 1) * 10).Limit(10)

	if price == "desc" || price == "asc" {
		query = query.Order("price " + price)
	}

	if createdDate == "desc" || createdDate == "asc" {
		query = query.Order("created_date " + createdDate)
	}

	var ads []models.Ads
	err := query.Find(&ads).Error

	return ads, err
}

// GetByID gets ad by its ID
func (r *repository) GetByID(id uint64) (models.Ads, error) {
	var ads models.Ads

	err := r.db.Where("id = ?", id).First(&ads).Error

	return ads, err
}
