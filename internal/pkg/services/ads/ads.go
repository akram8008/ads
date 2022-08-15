package ads

import (
	"ads/internal/model"
	"ads/internal/pkg/logger"
	"ads/internal/pkg/repositories/ads"
)

type Services interface {
	Create()
	GetAll()
	GetByID(id uint64) (model.Ads, error)
}

type services struct {
	repo   ads.Repository
	logger logger.Logger
}

type Params struct {
	Logger logger.Logger
}

func New(p Params) Services {
	return &services{
		repo:   ads.New(ads.Params{Logger: p.Logger}),
		logger: p.Logger,
	}
}

// Create creates ad
func (s *services) Create() {
	s.repo.Create()
}

// GetAll gets all ads
func (s *services) GetAll() {
	s.repo.GetAll()
}

// GetByID gets ad by its ID
func (s *services) GetByID(id uint64) (model.Ads, error) {
	return s.repo.GetByID(id)
}
