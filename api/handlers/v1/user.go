package v1

import (
	"context"
	l "github.com/NafisaTojiboyeva/api-gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"time"

	pb "github.com/NafisaTojiboyeva/api-gateway/genproto"
)

func (h *handlerV1) CreateTask(c *gin.Context) {
	var (
		body        pb.Task
		jspbmarshal protojson.MarshalOptions
	)
	jspbmarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.ToDoService().Create(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}
