package db

import (
	"ads/internal/pkg/models"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"ads/internal/domain"
	"ads/internal/pkg/logger"
)

type Database interface {
	Connection() *gorm.DB
	CloseConnection()
}

type database struct {
	db     *gorm.DB
	logger logger.Logger
}

func (d *database) Connection() *gorm.DB {
	return d.db
}

func (d *database) CloseConnection() {
	db, err := d.db.DB()
	if err != nil {
		d.logger.Logger().Error(err)
	}
	err = db.Close()
	if err != nil {
		d.logger.Logger().Error(err)
	}
}

var dbConn *gorm.DB

func Connect(cfg *domain.Config, logger *zap.SugaredLogger) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalln("Can not connect to database: ", err)
	}

	err = db.AutoMigrate(&models.Ads{})
	if err != nil {
		logger.Fatalln("Can not migrate tables : ", err)
	}

	logger.Infoln("Database successfully connected")

	dbConn = db
}

type Params struct {
	Logger logger.Logger
}

func New(p Params) Database {
	return &database{db: dbConn, logger: p.Logger}
}
