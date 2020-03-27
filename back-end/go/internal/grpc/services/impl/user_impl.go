package impl

import (
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services/server"
)

type UserServiceServer struct {
	ServiceServer server.ServiceServer
}

func NewUserServiceServer(serviceServer server.ServiceServer) services.UserServiceServer {
	return &UserServiceServer{
		ServiceServer: serviceServer,
	}
}
