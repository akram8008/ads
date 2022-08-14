package ads

import (
	"ads/internal/model"
	"ads/internal/pkg/repositories/ads"
)

type Services interface {
	Create()
	GetAll()
	GetByID(id string) (model.Ads, error)
}

type services struct {
	repo ads.Repository
}

func New() Services {
	return &services{
		repo: ads.New(),
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
func (s *services) GetByID(id string) (model.Ads, error) {
	return s.repo.GetByID(id)
}
