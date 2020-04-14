var envVariablesMap = require('./env.js');
var mongodb = require('../../internal/mongodb/index.js');
var server = require('../../internal/grpc/services/server/index.js');

var grpc = require('grpc');
var grpcServer = new grpc.Server();

var ProductServiceImpl = require('../../internal/grpc/services/impl/product_impl.js');
var ProductService = require('../../internal/grpc/services/product.js');

var dbConfig = {
    Username: envVariablesMap.get('TEST_DB_USERNAME'),
    Password: envVariablesMap.get('TEST_DB_PASSWORD'),
    Host:     envVariablesMap.get('TEST_DB_HOST'),
    Port:     envVariablesMap.get('TEST_DB_PORT'),
    Name:     envVariablesMap.get('TEST_DB_NAME')
};

var datastore = mongodb.InitializeDB(dbConfig);

var serviceServer = server.CreateServiceServer(datastore);

var grpcAddress = envVariablesMap.get("TEST_GRPC_SERVER_HOST") + ":" + envVariablesMap.get("TEST_GRPC_SERVER_PORT");

var ProductServiceServer = ProductServiceImpl.NewProductServiceServer(serviceServer);
grpcServer.addService(ProductService.serviceDescriptor, ProductServiceServer);

grpcServer.bind(grpcAddress, grpc.ServerCredentials.createInsecure());

grpcServer.start();

afterAll(async () => {
    mongodb.Close();
    grpcServer.forceShutdown();
});

module.exports = {
    GrpcAddress: grpcAddress,
    Datastore: serviceServer.Datastore
};