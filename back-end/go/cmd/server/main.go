package main

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services/impl"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services/server"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/middlewares"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/mongodb"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

func init() {
	var err error

	// Load the variables from .env file into the system.
	err = godotenv.Load("./.env")

	if err != nil {
		err = godotenv.Load("../../.env")

		if err != nil {
			log.Fatalf("Failed to load the .env file: %s", err.Error())
		}
	}
}

func main() {
	var dbUsername string
	var isSet bool
	var dbPassword string
	var dbHost string
	var dbPort string
	var dbName string
	var dbConfig mongodb.DBConfig
	var datastore mongodb.Datastore
	var err error
	var grpcHost string
	var grpcPort string
	var grpcAddress string
	var listener net.Listener
	var grpcServer *grpc.Server
	var serviceServer server.ServiceServer
	var userServiceServer services.UserServiceServer
	var productServiceServer services.ProductServiceServer
	var promotionServiceServer services.PromotionServiceServer
	var discountedDateServiceServer services.DiscountedDateServiceServer
	var ctx context.Context
	var cancel context.CancelFunc
	var mux *runtime.ServeMux
	var httpHost string
	var httpPort string
	var httpAddress string
	var grpcOpts []grpc.DialOption

	dbUsername, isSet = os.LookupEnv("DB_USERNAME")

	if !isSet {
		log.Fatal("Failed to read the DB_USERNAME environment variable: it isn't set")
	}

	dbPassword, isSet = os.LookupEnv("DB_PASSWORD")

	if !isSet {
		log.Fatal("Failed to read the DB_PASSWORD environment variable: it isn't set")
	}

	dbHost, isSet = os.LookupEnv("DB_HOST")

	if !isSet {
		log.Fatal("Failed to read the DB_HOST environment variable: it isn't set")
	}

	dbPort, isSet = os.LookupEnv("DB_PORT")

	if !isSet {
		log.Fatal("Failed to read the DB_PORT environment variable: it isn't set")
	}

	dbName, isSet = os.LookupEnv("DB_NAME")

	if !isSet {
		log.Fatal("Failed to read the DB_NAME environment variable: it isn't set")
	}

	dbConfig = mongodb.DBConfig{
		DBUsername: dbUsername,
		DBPassword: dbPassword,
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBName:     dbName,
	}

	// Initialize the database.
	datastore, err = mongodb.InitializeDB(dbConfig)

	if err != nil {
		log.Fatal("Failed to configure the database: ", err.Error())
	}

	// Check the gRPC environment variables for the service 2.
	grpcHost, isSet = os.LookupEnv("GRPC_SERVER_HOST")

	if !isSet {
		log.Fatal("Failed to read the GRPC_SERVER_HOST environment variable: it isn't set")
	}

	grpcPort, isSet = os.LookupEnv("GRPC_SERVER_PORT")

	if !isSet {
		log.Fatal("Failed to read the GRPC_SERVER_PORT environment variable: it isn't set")
	}

	grpcAddress = fmt.Sprintf("%s:%s", grpcHost, grpcPort)

	listener, err = net.Listen("tcp", grpcAddress)

	if err != nil {
		log.Fatalf("Failed to created the server on local tcp network address %s: %s", grpcAddress, err.Error())
	}

	grpcServer = grpc.NewServer()

	serviceServer = server.CreateServiceServer(datastore)

	// Check the gRPC environment variables for the microservice 1.
	grpcHost, isSet = os.LookupEnv("GRPC_SERVER_HOST_MS_1")

	if !isSet {
		log.Fatal("Failed to read the GRPC_SERVER_HOST_MS_1 environment variable: it isn't set")
	}

	grpcPort, isSet = os.LookupEnv("GRPC_SERVER_PORT_MS_1")

	if !isSet {
		log.Fatal("Failed to read the GRPC_SERVER_PORT_MS_1 environment variable: it isn't set")
	}

	userServiceServer = impl.NewUserServiceServer(serviceServer)
	services.RegisterUserServiceServer(grpcServer, userServiceServer)

	productServiceServer = impl.NewProductServiceServer(serviceServer, grpcHost, grpcPort)
	services.RegisterProductServiceServer(grpcServer, productServiceServer)

	promotionServiceServer = impl.NewPromotionServiceServer(serviceServer)
	services.RegisterPromotionServiceServer(grpcServer, promotionServiceServer)

	discountedDateServiceServer = impl.NewDiscountedDateServiceServer(serviceServer)
	services.RegisterDiscountedDateServiceServer(grpcServer, discountedDateServiceServer)

	log.Printf("Starting the GRPC server connection on %s", grpcAddress)

	go func() {
		err = grpcServer.Serve(listener)

		if err != nil {
			log.Fatalf("Failed to start the GRPC server connection to %s: %s", grpcAddress, err.Error())
		}
	}()

	ctx = context.Background()

	ctx, cancel = context.WithCancel(ctx)

	defer cancel()

	mux = runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(middlewares.CustomHeaderMatcher))

	// Check the http environment variables.
	httpHost, isSet = os.LookupEnv("HTTP_SERVER_HOST")

	if !isSet {
		log.Fatal("Failed to read the HTTP_SERVER_HOST environment variable: it isn't set")
	}

	httpPort, isSet = os.LookupEnv("HTTP_SERVER_PORT")

	if !isSet {
		log.Fatal("Failed to read the HTTP_SERVER_PORT environment variable: it isn't set")
	}

	httpAddress = fmt.Sprintf("%s:%s", httpHost, httpPort)

	grpcOpts = []grpc.DialOption{grpc.WithInsecure()}

	err = services.RegisterUserServiceHandlerFromEndpoint(ctx, mux, grpcAddress, grpcOpts)

	if err != nil {
		log.Fatalf("Failed to register the endpoints of the user service: %s", err.Error())
	}

	err = services.RegisterProductServiceHandlerFromEndpoint(ctx, mux, grpcAddress, grpcOpts)

	if err != nil {
		log.Fatalf("Failed to register the endpoints of the product service: %s", err.Error())
	}

	err = services.RegisterPromotionServiceHandlerFromEndpoint(ctx, mux, grpcAddress, grpcOpts)

	if err != nil {
		log.Fatalf("Failed to register the endpoints of the promotion service: %s", err.Error())
	}

	err = services.RegisterDiscountedDateServiceHandlerFromEndpoint(ctx, mux, grpcAddress, grpcOpts)

	if err != nil {
		log.Fatalf("Failed to register the endpoints of the discounted date service: %s", err.Error())
	}

	log.Printf("Starting the HTTP server connection on %s", httpAddress)

	go func() {
		err = http.ListenAndServe(httpAddress, mux)

		if err != nil {
			log.Fatalf("Failed to start the HTTP server connection to %s: %s", httpAddress, err.Error())
		}
	}()

	// Graceful disconnect.
	WaitForShutdown()

	grpcServer.Stop()

	listener.Close()

	datastore.Close()

	log.Println("Done")
}

func WaitForShutdown() {
	var interruptChan chan os.Signal

	// Create a channel to receive OS signals.
	interruptChan = make(chan os.Signal)

	// Relay os.Interrupt to our channel (os.Interrupt = CTRL+C)
	// ignoring other incoming signals.
	signal.Notify(interruptChan, os.Interrupt)

	// Block the main routine so that to keep it running until a signal is received.
	// If the main routine is shut down, the child one that is serving the server will shut down as well.
	<-interruptChan

	log.Println("Shutting down the server...")
}
