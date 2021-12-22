package main

import (
	"github.com/NafisaTojiboyeva/api-gateway/api"
	"github.com/NafisaTojiboyeva/api-gateway/config"
	"github.com/NafisaTojiboyeva/api-gateway/pkg/logger"
	"github.com/NafisaTojiboyeva/api-gateway/service"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	serviceManager, err := service.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Error("failed to run http server", logger.Error(err))
		panic(err)
	}
}
