package ads

import (
	"ads/internal/pkg/services/ads"
	"net/http"
)

type Handlers interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
}

type handlers struct {
	adsService ads.Services
}

func New() Handlers {
	return &handlers{
		adsService: ads.New(),
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
	id := r.URL.Query().Get("id")
	h.adsService.GetByID(id)
}
