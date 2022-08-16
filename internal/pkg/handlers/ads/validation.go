package ads

import (
	"ads/internal/domain"
	"errors"
	"net/url"
)

func (h *handlers) validateAds(newAdsReq domain.AdsRequest) error {
	// Validating name
	if newAdsReq.Name == nil || *newAdsReq.Name == "" {
		return errors.New("no name provided")
	} else if len([]rune(*newAdsReq.Name)) > 200 {
		return errors.New("name length is limited to 200 characters")
	}

	// Validating description
	if newAdsReq.Description == nil || *newAdsReq.Description == "" {
		return errors.New("no description provided")
	} else if len([]rune(*newAdsReq.Description)) > 1000 {
		return errors.New("description length is limited to 1000 characters")
	}

	// Validating links
	if newAdsReq.Links == nil || len(*newAdsReq.Links) == 0 {
		return errors.New("no links provided")
	} else if len(*newAdsReq.Links) > 3 {
		return errors.New("links limited to 3 links")
	}
	for _, link := range *newAdsReq.Links {
		_, err := url.ParseRequestURI(link)
		if err != nil {
			return errors.New(link + " doesn't look like a link")
		}
	}

	// Validating price
	if newAdsReq.Price == nil {
		return errors.New("no price provided")
	} else if *newAdsReq.Price < 0 {
		return errors.New("price can not be negative")
	}

	return nil
}
