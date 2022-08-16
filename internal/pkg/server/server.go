package server

import (
	"ads/internal/domain"
	"ads/internal/pkg/handlers/ads"
	"ads/internal/pkg/logger"
	"github.com/gorilla/mux"
	"net/http"
)

type Params struct {
	Config   domain.Config
	Handlers ads.Handlers
	Logger   logger.Logger
}

func New(p Params) *http.Server {
	muxRouter := mux.NewRouter()

	muxRouter.Handle("/ads", http.HandlerFunc(p.Handlers.Create)).Methods("POST")
	muxRouter.Handle("/ads", http.HandlerFunc(p.Handlers.GetByID)).Methods("GET")
	muxRouter.Handle("/ads/all", http.HandlerFunc(p.Handlers.GetAll)).Methods("GET")

	server := http.Server{
		Addr:    p.Config.Sever.Host + ":" + p.Config.Sever.Port,
		Handler: muxRouter,
	}

	return &server
}
