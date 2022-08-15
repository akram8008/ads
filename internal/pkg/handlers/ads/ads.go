package ads

import (
	"ads/internal/model"
	"ads/internal/pkg/logger"
	"ads/internal/pkg/services/ads"
	"net/http"
	"strconv"
)

type Handlers interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
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
	h.adsService.Create()
}

// GetAll gets all ads
func (h *handlers) GetAll(w http.ResponseWriter, r *http.Request) {
	h.adsService.GetAll()
}

// GetByID gets ad by its ID
func (h *handlers) GetByID(w http.ResponseWriter, r *http.Request) {
	var resp model.ApiResponse
	defer resp.Respond(w)

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		resp.Set(http.StatusBadRequest, "id of ads can not be empty", nil)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.Set(http.StatusBadRequest, "id of ads must be integer", nil)
		return
	}

	receivedAds, err := h.adsService.GetByID(uint64(id))
	if err != nil {
		resp.Set(http.StatusInternalServerError, "Can not get ads with id: "+idStr, nil)
		return
	}
	if receivedAds.ID == 0 {
		resp.Set(http.StatusOK, "There are no ads with ID: "+idStr, nil)
		return
	}
	resp.Set(http.StatusOK, "Success", receivedAds)
}
