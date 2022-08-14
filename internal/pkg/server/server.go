package server

import (
	"ads/internal/model"
	"ads/internal/pkg/handlers/ads"
	"ads/internal/pkg/logger"
	"github.com/gorilla/mux"
	"net/http"
)

type Params struct {
	Config   model.Config
	Handlers ads.Handlers
	Logger   logger.Logger
}

func New(p Params) *http.Server {
	muxRouter := mux.NewRouter()

	muxRouter.Handle("/ad", http.HandlerFunc(p.Handlers.Create)).Methods("POST")
	muxRouter.Handle("/ad", http.HandlerFunc(p.Handlers.GetByID)).Methods("GET")
	muxRouter.Handle("/ads", http.HandlerFunc(p.Handlers.GetAll)).Methods("GET")

	server := http.Server{
		Addr:    p.Config.Host + ":" + p.Config.Port,
		Handler: muxRouter,
	}

	return &server
}
