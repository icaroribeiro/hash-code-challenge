package impl

import (
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services/server"
)

type DiscountedDateServiceServer struct {
	ServiceServer server.ServiceServer
}

func NewDiscountedDateServiceServer(serviceServer server.ServiceServer) services.DiscountedDateServiceServer {
	return &DiscountedDateServiceServer{
		ServiceServer: serviceServer,
	}
}
