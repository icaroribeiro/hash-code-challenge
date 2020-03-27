package impl

import (
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services/server"
)

type ProductServiceServer struct {
	ServiceServer server.ServiceServer
	// The gRPC host and port used to connect to the microservice 1.
	GrpcHost string
	GrpcPort string
}

func NewProductServiceServer(serviceServer server.ServiceServer, grpcHost string, grpcPort string) services.ProductServiceServer {
	return &ProductServiceServer{
		ServiceServer: serviceServer,
		GrpcHost:      grpcHost,
		GrpcPort:      grpcPort,
	}
}
