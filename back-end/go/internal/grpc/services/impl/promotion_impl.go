package impl

import (
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services/server"
)

type PromotionServiceServer struct {
	ServiceServer server.ServiceServer
}

func NewPromotionServiceServer(serviceServer server.ServiceServer) services.PromotionServiceServer {
	return &PromotionServiceServer{
		ServiceServer: serviceServer,
	}
}
