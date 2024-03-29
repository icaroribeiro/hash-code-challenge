package impl_test

import (
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services/impl"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services/server"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/mongodb"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/utils"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/test/bufconn"
    "log"
    "net"
    "os"
    "testing"
)

var envVariablesMap map[string]string

var datastore mongodb.Datastore

var ctx context.Context

var clientConn *grpc.ClientConn

var userServiceClient services.UserServiceClient

var productServiceClient services.ProductServiceClient

var promotionServiceClient services.PromotionServiceClient

var discountedDateServiceClient services.DiscountedDateServiceClient

func init() {
    var filenames []string
    var err error

    envVariablesMap = make(map[string]string)

    envVariablesMap["TEST_DB_USERNAME"] = ""
    envVariablesMap["TEST_DB_PASSWORD"] = ""
    envVariablesMap["TEST_DB_HOST"] = ""
    envVariablesMap["TEST_DB_PORT"] = ""
    envVariablesMap["TEST_DB_NAME"] = ""

    filenames = []string{"../../../../.test.env"}

    err = utils.GetEnvVariables(filenames, envVariablesMap)

    if err != nil {
        log.Fatal(err.Error())
    }
}

// It serves as a wrapper around the testMain function that allows to defer other functions.
// At the end, it finally passes the returned exit code to os.Exit().
func TestMain(m *testing.M) {
    var exitVal int

    // Before the tests.
    utils.InitializeRandomization()

    exitVal = testMain(m)

    // After the tests.
    defer clientConn.Close()
    defer datastore.Close()

    os.Exit(exitVal)
}

// It configures the settings before running the tests. It returns an integer denoting an exit code to be used
// in the TestMain function. In the case, if the exit code is 0 it denotes success while all other codes denote failure.
func testMain(m *testing.M) int {
    var dbConfig mongodb.DBConfig
    var err error
    var bufferSize int
    var listener *bufconn.Listener
    var grpcServer *grpc.Server
    var serviceServer server.ServiceServer
    var userServiceServer services.UserServiceServer
    var productServiceServer services.ProductServiceServer
    var promotionServiceServer services.PromotionServiceServer
    var discountedDateServiceServer services.DiscountedDateServiceServer

    dbConfig = mongodb.DBConfig{
        Username: envVariablesMap["TEST_DB_USERNAME"],
        Password: envVariablesMap["TEST_DB_PASSWORD"],
        Host:     envVariablesMap["TEST_DB_HOST"],
        Port:     envVariablesMap["TEST_DB_PORT"],
        Name:     envVariablesMap["TEST_DB_NAME"],
    }

    datastore, err = mongodb.InitializeDB(dbConfig)

    if err != nil {
        log.Printf("Failed to configure the database: %s", err.Error())
        return 1
    }

    bufferSize = 1024 * 1024

    listener = bufconn.Listen(bufferSize)

    grpcServer = grpc.NewServer()

    serviceServer = server.CreateServiceServer(datastore)

    userServiceServer = impl.NewUserServiceServer(serviceServer)
    services.RegisterUserServiceServer(grpcServer, userServiceServer)

    productServiceServer = impl.NewProductServiceServer(serviceServer, "", "")
    services.RegisterProductServiceServer(grpcServer, productServiceServer)

    promotionServiceServer = impl.NewPromotionServiceServer(serviceServer)
    services.RegisterPromotionServiceServer(grpcServer, promotionServiceServer)

    discountedDateServiceServer = impl.NewDiscountedDateServiceServer(serviceServer)
    services.RegisterDiscountedDateServiceServer(grpcServer, discountedDateServiceServer)

    go func() {
        err = grpcServer.Serve(listener)

        if err != nil {
            log.Fatalf("Failed to start the GRPC server: %s", err.Error())
        }
    }()

    ctx = context.Background()

    // With the approach below we can avoid starting up a service with a real port number, but still allowing testing
    // of network behavior. Here, we will have an in-memory connection without using OS-level resources like ports
    // that may or may not clean up quickly. The trick is setting the WithDialer option using the bufconn package
    // to create a listener that exposes its own dialer.
    bufDialer := func(listener *bufconn.Listener) func(context.Context, string) (net.Conn, error) {
        return func(ctx context.Context, url string) (net.Conn, error) {
            return listener.Dial()
        }
    }

    clientConn, err = grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer(listener)), grpc.WithInsecure())

    if err != nil {
        log.Printf("Failed to create a client connection: %s", err.Error())
        return 1
    }

    userServiceClient = services.NewUserServiceClient(clientConn)

    productServiceClient = services.NewProductServiceClient(clientConn)

    promotionServiceClient = services.NewPromotionServiceClient(clientConn)

    discountedDateServiceClient = services.NewDiscountedDateServiceClient(clientConn)

    return m.Run()
}
