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

	log := logger.New([]string{"logs/ads.log"})

	db.Connect(cfg, log.Logger())
	defer db.New(db.Params{}).CloseConnection()

	srv := server.New(server.Params{
		Config:   *cfg,
		Logger:   log,
		Handlers: ads.New(ads.Params{Logger: log}),
	})

	log.Logger().Infoln("[Server] listening on ", cfg.Sever)
	log.Logger().Fatal(srv.ListenAndServe())
}
