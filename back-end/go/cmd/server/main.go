package main

import (
    "fmt"
    "github.com/grpc-ecosystem/grpc-gateway/runtime"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services/impl"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services/server"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/middlewares"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/mongodb"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/utils"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "log"
    "net"
    "net/http"
    "os"
    "os/signal"
)

var envVariablesMap map[string]string

func init() {
    var filenames []string
    var err error

    envVariablesMap = make(map[string]string)

    // The environment variables related to the database settings.
    envVariablesMap["DB_USERNAME"] = ""
    envVariablesMap["DB_PASSWORD"] = ""
    envVariablesMap["DB_HOST"] = ""
    envVariablesMap["DB_PORT"] = ""
    envVariablesMap["DB_NAME"] = ""
    
    // The environment variables related to the gRPC server for the microservice 2.
    envVariablesMap["GRPC_SERVER_HOST"] = ""
    envVariablesMap["GRPC_SERVER_PORT"] = ""

    // The environment variables related to the gRPC server for the microservice 1.
    envVariablesMap["GRPC_SERVER_HOST_MS_1"] = ""
    envVariablesMap["GRPC_SERVER_PORT_MS_1"] = ""

    // The environment variables related to the HTTP server.
    envVariablesMap["HTTP_SERVER_HOST"] = ""
    envVariablesMap["HTTP_SERVER_PORT"] = ""

    // The environment files from where the variables will be loaded.
    filenames = []string{"../../.env"}

    err = utils.GetEnvVariables(filenames, envVariablesMap)

    if err != nil {
        log.Fatal(err.Error())
    }
}

func main() {
    var dbConfig mongodb.DBConfig
    var datastore mongodb.Datastore
    var err error
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
    var httpAddress string
    var grpcOpts []grpc.DialOption

    dbConfig = mongodb.DBConfig{
        Username: envVariablesMap["DB_USERNAME"],
        Password: envVariablesMap["DB_PASSWORD"],
        Host:     envVariablesMap["DB_HOST"],
        Port:     envVariablesMap["DB_PORT"],
        Name:     envVariablesMap["DB_NAME"],
    }

    // Initialize the database.
    datastore, err = mongodb.InitializeDB(dbConfig)

    if err != nil {
        log.Fatal("Failed to configure the database: ", err.Error())
    }

    grpcAddress = fmt.Sprintf("%s:%s", envVariablesMap["GRPC_SERVER_HOST"], envVariablesMap["GRPC_SERVER_PORT"])

    listener, err = net.Listen("tcp", grpcAddress)

    if err != nil {
        log.Fatalf("Failed to created the server on local tcp network address %s: %s", grpcAddress, err.Error())
    }

    grpcServer = grpc.NewServer()

    serviceServer = server.CreateServiceServer(datastore)

    // Register the gRPC services.
    userServiceServer = impl.NewUserServiceServer(serviceServer)
    services.RegisterUserServiceServer(grpcServer, userServiceServer)

    productServiceServer = impl.NewProductServiceServer(serviceServer, 
        envVariablesMap["GRPC_SERVER_HOST_MS_1"], envVariablesMap["GRPC_SERVER_PORT_MS_1"])
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

    httpAddress = fmt.Sprintf("%s:%s", envVariablesMap["HTTP_SERVER_HOST"], envVariablesMap["HTTP_SERVER_PORT"])

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

    err = datastore.Close()

    if err != nil {
        log.Fatalf("Failed to close the database: %s", err.Error())
    }

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
