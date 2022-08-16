package helper

import (
	"ads/internal/domain"
	"ads/internal/pkg/models"
	"strings"
)

func ConvertFromAdsDomain(adsDomain domain.AdsRequest) models.Ads {
	return models.Ads{
		Name:        *adsDomain.Name,
		Description: *adsDomain.Description,
		Links:       strings.Join(*adsDomain.Links, "  "),
		Price:       *adsDomain.Price,
	}
}

func ConvertToAdsDomain(adsModel models.Ads, fields string) domain.AdsResponse {

	var description *string
	if strings.Contains(fields, models.AdsDescriptionField) {
		description = &adsModel.Description
	}

	var links *[]string
	if strings.Contains(fields, models.AdsLinksField) {
		linksSlice := strings.Split(adsModel.Links, "  ")
		links = &linksSlice
	}

	if description != nil || links != nil {
		return domain.AdsResponse{
			Description: description,
			Links:       links,
		}
	}

	return domain.AdsResponse{
		Name:     &adsModel.Name,
		Price:    &adsModel.Price,
		MainLink: &strings.Split(adsModel.Links, "  ")[0],
	}
}
