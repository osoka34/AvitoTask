package main

import (
	"AvitoTask/config"
	"AvitoTask/internal/httpServer"
	"AvitoTask/pkg/logger"
	"context"
	"fmt"
	"log"
)

func main() {

	ctx := context.Background()

	viperInstance, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Cannot load config. Error: {%s}", err.Error())
	}

	cfg, err := config.ParseConfig(viperInstance)
	if err != nil {
		log.Fatalf("Cannot parse config. Error: {%s}", err.Error())
	}

	logger := logger.NewLogger(cfg).Sugar()
	defer logger.Sync()

	logger.Infof("file path: %s", cfg.Statistic.FilePath)

	s := httpServer.NewServer(cfg, logger)
	if err = s.Run(ctx); err != nil {
		logger.Error(fmt.Sprint(err))
	}
}
