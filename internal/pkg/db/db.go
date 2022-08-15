package db

import (
	"ads/internal/model"
	"ads/internal/pkg/logger"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type Database interface {
	Connection() *sql.DB
	CloseConnection() error
}

type database struct {
	db     *sql.DB
	logger logger.Logger
}

func (d *database) Connection() *sql.DB {
	return d.db
}

func (d *database) CloseConnection() error {
	return d.db.Close()
}

var dbConn *sql.DB

func Connect(cfg *model.Config, logger *zap.SugaredLogger) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	logger.Infoln("Database - successfully connected!")

	var tmp string
	err = db.QueryRow("select 'done'").Scan(&tmp)
	if err != nil {
		panic(err)
	}

	logger.Infoln("tmp: ", tmp)

	dbConn = db
}

type Params struct {
	Logger logger.Logger
}

func New(p Params) Database {
	return &database{db: dbConn, logger: p.Logger}
}
