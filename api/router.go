package api

import (
	v1 "github.com/NafisaTojiboyeva/api-gateway/api/handlers/v1"
	"github.com/gin-gonic/gin"

	"github.com/NafisaTojiboyeva/api-gateway/config"
	"github.com/NafisaTojiboyeva/api-gateway/pkg/logger"
	"github.com/NafisaTojiboyeva/api-gateway/service"
)

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager service.IServiceManager
}

// New ...
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router.Group("/v1")
	api.POST("/todos", handlerV1.CreateTask)
	api.GET("/todos/:id", handlerV1.GetTask)
	api.GET("/todos", handlerV1.ListTasks)
	api.PUT("/todos/:id", handlerV1.UpdateTask)
	api.DELETE("/todos/:id", handlerV1.DeleteTask)
	api.GET("/todos/overdue", handlerV1.ListOverdueTasks)

	return router
}
