package internal

import (
	"ads/internal/pkg/config"
	"ads/internal/pkg/db"
	"ads/internal/pkg/handlers/ads"
	"ads/internal/pkg/logger"
	"ads/internal/pkg/server"
)

func Up() {
	cfg := config.New("configs/configs.json")
	logger := logger.New([]string{"logs/ads.log"})
	db.Connect(cfg, logger.Logger())
	srv := server.New(server.Params{
		Config:   *cfg,
		Logger:   logger,
		Handlers: ads.New(),
	})

	logger.Logger().Fatal(srv.ListenAndServe())
}
