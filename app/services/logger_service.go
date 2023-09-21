package services

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() error {
	var err error
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stdout", "./public/logs/app.log"}
	Logger, err = cfg.Build()
	if err != nil {
		return err
	}
	defer Logger.Sync()
	return nil
}
