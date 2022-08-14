package logger

import (
	"go.uber.org/zap"
	"log"
)

type Logger interface {
	Logger() *zap.SugaredLogger
}

type logger struct {
	logger *zap.SugaredLogger
}

func (l *logger) Logger() *zap.SugaredLogger {
	return l.logger
}
func New(paths []string) Logger {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = paths
	cfg.DisableStacktrace = true
	l, err := cfg.Build()
	if err != nil {
		log.Fatalf("%v", err)
	}

	return &logger{
		logger: l.Sugar(),
	}
}
