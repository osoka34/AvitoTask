package main

import (
	"AvitoTask/config"
	"AvitoTask/internal/httpServer"
	"AvitoTask/pkg/logger"
	"fmt"
	"log"
)

func main() {

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

	logger.Info("Hello from Zap!")

	s := httpServer.NewServer(cfg, logger)
	if err = s.Run(); err != nil {
		logger.Error(fmt.Sprint(err))
	}
}
