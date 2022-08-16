package ads

import (
	"ads/internal/domain"
	"ads/internal/pkg/logger"
	"ads/internal/pkg/repositories/ads"
	"ads/internal/pkg/services/helper"
)

type Services interface {
	Create(domain.AdsRequest) (uint64, error)
	GetAll(int, string, string) ([]domain.AdsResponse, error)
	GetByID(uint64, string) (domain.AdsResponse, error)
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
func (s *services) Create(adsReq domain.AdsRequest) (uint64, error) {
	return s.repo.Create(helper.ConvertFromAdsDomain(adsReq))
}

// GetAll gets all ads
func (s *services) GetAll(page int, price, createdDate string) ([]domain.AdsResponse, error) {
	receivedAds, err := s.repo.GetAll(page, price, createdDate)
	if err != nil {
		return []domain.AdsResponse{}, err
	}

	var newAds []domain.AdsResponse
	for _, item := range receivedAds {
		newAds = append(newAds, helper.ConvertToAdsDomain(item, ""))
	}

	return newAds, nil
}

// GetByID gets ad by its ID
func (s *services) GetByID(id uint64, fields string) (domain.AdsResponse, error) {
	newAds, err := s.repo.GetByID(id)
	if err != nil {
		return domain.AdsResponse{}, err
	}
	return helper.ConvertToAdsDomain(newAds, fields), nil
}
