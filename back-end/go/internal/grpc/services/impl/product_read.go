package impl

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/entities"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
)

func (p *ProductServiceServer) GetAllProducts(ctx context.Context, 
	e *empty.Empty) (*services.GetAllProductsResponse, error) {
		var xUserId string
		var incomingMetadata metadata.MD
		var hasMetada bool
		var hasXUserId bool
		var values []string
		var grpcAddress string
		var clientConn *grpc.ClientConn
		var err error
		var client services.ProductServiceClient
		var request empty.Empty
		var outgoingCtx context.Context
		var response *services.GetAllProductsResponse
		var products []models.Product
		var product models.Product
		var productEntity *entities.Product

		xUserId = ""

		incomingMetadata, hasMetada = metadata.FromIncomingContext(ctx)

		// Retrieve data from a map based on metadata keys to values.
		if hasMetada {
			if len(incomingMetadata) > 0 {
				values, hasXUserId = incomingMetadata["X-USER-ID"]

				if hasXUserId {
					xUserId = values[0]
				} else {
					values, hasXUserId = incomingMetadata["x-user-id"]

					if hasXUserId {
						xUserId = values[0]
					}
				}
			}
		}

		// Try to communicate with the microservice 1 to obtain the list of all products after evaluating
		// if any discount must be applied based on criteria related to discounted dates.
		grpcAddress = fmt.Sprintf("%s:%s", p.GrpcHost, p.GrpcPort)

		clientConn, err = grpc.Dial(grpcAddress, grpc.WithInsecure())

		if err != nil {
			return nil, status.Error(codes.Internal,
				fmt.Sprintf("Failed to establish client connection to %s: %s", grpcAddress, err.Error()))
		}

		defer clientConn.Close()

		client = services.NewProductServiceClient(clientConn)

		request = empty.Empty{}

		outgoingCtx = context.Background()

		if hasXUserId && (xUserId != "") {
			outgoingCtx = metadata.AppendToOutgoingContext(outgoingCtx, "X-USER-ID", xUserId)
		}

		response, err = client.GetAllProducts(outgoingCtx, &request)

		if err != nil {
			log.Printf("Failed to get the list of all products using the microservice 1: %s", err.Error())

			// In case of the microservice 1 goes down, the microservice 2 must work normally
			// and then return the list of all products without applying any discount.
			products, err = p.ServiceServer.Datastore.GetAllProducts()

			if err != nil {
				return nil, status.Error(codes.Internal,
					fmt.Sprintf("Failed to get the list of all products: %s", err.Error()))
			}

			response = &services.GetAllProductsResponse{}

			for _, product = range products {
				productEntity = &entities.Product{
					Id:           product.ID.Hex(),
					PriceInCents: int32(product.PriceInCents),
					Title:        product.Title,
					Description:  product.Description,
				}

				response.Products = append(response.Products, productEntity)
			}

			return response, nil
		}

		return response, nil
}

func (p *ProductServiceServer) GetProduct(ctx context.Context, 
	request *services.GetProductRequest) (*entities.Product, error) {
		var xUserId string
		var incomingMetadata metadata.MD
		var hasMetada bool
		var hasXUserId bool
		var values []string
		var grpcAddress string
		var clientConn *grpc.ClientConn
		var err error
		var client services.ProductServiceClient
		var outgoingCtx context.Context
		var response *entities.Product
		var product models.Product

		if request.Id == "" {
			return nil, status.Error(codes.InvalidArgument,
				"The id is required and must be set to a non-empty value in the request URL")
		}

		xUserId = ""

		incomingMetadata, hasMetada = metadata.FromIncomingContext(ctx)

		// Retrieve data from a map based on metadata keys to values.
		if hasMetada {
			if len(incomingMetadata) > 0 {
				values, hasXUserId = incomingMetadata["X-USER-ID"]

				if hasXUserId {
					xUserId = values[0]
				} else {
					values, hasXUserId = incomingMetadata["x-user-id"]

					if hasXUserId {
						xUserId = values[0]
					}
				}
			}
		}

		// Try to communicate with the microservice 1 to obtain the list of all products after evaluating
		// if any discount must be applied based on criteria related to discounted dates.
		grpcAddress = fmt.Sprintf("%s:%s", p.GrpcHost, p.GrpcPort)

		clientConn, err = grpc.Dial(grpcAddress, grpc.WithInsecure())

		if err != nil {
			return nil, status.Error(codes.Internal,
				fmt.Sprintf("Failed to establish client connection to %s: %s", grpcAddress, err.Error()))
		}

		defer clientConn.Close()

		client = services.NewProductServiceClient(clientConn)

		outgoingCtx = context.Background()

		if hasXUserId && (xUserId != "") {
			outgoingCtx = metadata.AppendToOutgoingContext(outgoingCtx, "X-USER-ID", xUserId)
		}

		response, err = client.GetProduct(outgoingCtx, request)

		if err != nil {
			log.Printf("Failed to get the product with the id %s using the microservice 1: %s", request.Id, err.Error())

			// In case of the microservice 1 goes down, the microservice 2 must work normally
			// and then return a specific product without applying any discount.
			product, err = p.ServiceServer.Datastore.GetProduct(request.Id)

			if err != nil {
				return nil, status.Error(codes.Internal,
					fmt.Sprintf("Failed to get the product with the id %s: %s", request.Id, err.Error()))
			}

			if product.ID.IsZero() {
				return nil, status.Error(codes.NotFound,
					fmt.Sprintf("Failed to get the product with the id %s: the product wasn't found", request.Id))
			}

			response = &entities.Product{
				Id:           product.ID.Hex(),
				PriceInCents: int32(product.PriceInCents),
				Title:        product.Title,
				Description:  product.Description,
			}

			return response, nil
		}

		return response, nil
}
