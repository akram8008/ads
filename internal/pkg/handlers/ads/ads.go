package ads

import (
	"ads/internal/domain"
	"ads/internal/pkg/logger"
	"ads/internal/pkg/services/ads"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Handlers interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	validateAds(body domain.AdsRequest) error
}

type handlers struct {
	adsService ads.Services
	logger     logger.Logger
}

type Params struct {
	Logger logger.Logger
}

func New(p Params) Handlers {
	return &handlers{
		adsService: ads.New(ads.Params{Logger: p.Logger}),
		logger:     p.Logger,
	}
}

// Create creates ad
func (h *handlers) Create(w http.ResponseWriter, r *http.Request) {
	var resp domain.ApiResponse
	defer resp.Respond(w)

	var newAdsReq domain.AdsRequest
	err := json.NewDecoder(r.Body).Decode(&newAdsReq)
	if err != nil {
		resp.Set(http.StatusBadRequest, "Wrong ads model in body request", nil)
		h.logger.Logger().Error(err)
		return
	}

	if err = h.validateAds(newAdsReq); err != nil {
		resp.Set(http.StatusBadRequest, "Validation error: "+err.Error(), nil)
		h.logger.Logger().Error(fmt.Sprintf("Validation error. Bad Ads - model in body request: %v Error: %v", newAdsReq, err.Error()))
		return
	}

	ID, err := h.adsService.Create(newAdsReq)
	if err != nil {
		resp.Set(http.StatusInternalServerError, "Can not create new ads", nil)
		h.logger.Logger().Error(err)
		return
	}

	resp.Set(http.StatusOK, "Success", ID)

}

// GetAll gets all ads
func (h *handlers) GetAll(w http.ResponseWriter, r *http.Request) {
	var resp domain.ApiResponse
	defer resp.Respond(w)

	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		resp.Set(http.StatusBadRequest, "page must be integer", nil)
		h.logger.Logger().Error("Can not convert " + pageStr + " into integer")
		return
	}

	price := r.URL.Query().Get("price")
	createdDate := r.URL.Query().Get("created_date")

	receivedAds, err := h.adsService.GetAll(page, price, createdDate)
	if err != nil {
		resp.Set(http.StatusInternalServerError, "can not get all ads", nil)
		h.logger.Logger().Error(err)
		return
	}
	resp.Set(http.StatusOK, "Success", receivedAds)
}

// GetByID gets ad by its ID
func (h *handlers) GetByID(w http.ResponseWriter, r *http.Request) {
	var resp domain.ApiResponse
	defer resp.Respond(w)

	// getting id from query
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		resp.Set(http.StatusBadRequest, "id of ads can not be empty", nil)
		h.logger.Logger().Error("Empty ID")
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.Set(http.StatusBadRequest, "id of ads must be integer", nil)
		h.logger.Logger().Error("Can not convert " + idStr + " into integer")
		return
	}

	// getting fields from query
	fields := r.URL.Query().Get("fields")

	receivedAds, err := h.adsService.GetByID(uint64(id), fields)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.Set(http.StatusOK, "There are no ads with ID: "+idStr, nil)
		h.logger.Logger().Error(err)
		return
	}
	if err != nil {
		resp.Set(http.StatusInternalServerError, "Can not get ads", nil)
		h.logger.Logger().Error(err)
		return
	}

	resp.Set(http.StatusOK, "Success", receivedAds)
}
