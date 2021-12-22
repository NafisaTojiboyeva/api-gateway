package service

import (
	"fmt"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"

	"github.com/NafisaTojiboyeva/api-gateway/config"
	pb "github.com/NafisaTojiboyeva/api-gateway/genproto"
)

type IServiceManager interface {
	ToDoService() pb.ToDoServiceClient
}

type serviceManager struct {
	todoService pb.ToDoServiceClient
}

func (s *serviceManager) ToDoService() pb.ToDoServiceClient {
	return s.todoService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connToDo, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.ToDoServiceHost, conf.ToDoServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		todoService: pb.NewToDoServiceClient(connToDo),
	}

	return serviceManager, nil
}
