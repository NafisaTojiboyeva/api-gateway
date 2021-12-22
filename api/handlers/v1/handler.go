package v1

import (
	"github.com/NafisaTojiboyeva/api-gateway/config"
	"github.com/NafisaTojiboyeva/api-gateway/pkg/logger"
	"github.com/NafisaTojiboyeva/api-gateway/service"
)

type handlerV1 struct {
	log            logger.Logger
	serviceManager service.IServiceManager
	cfg            config.Config
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Logger         logger.Logger
	ServiceManager service.IServiceManager
	Cfg            config.Config
}

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
	}
}
