var envVariablesMap = require('../../env.js');
var mongodb = require('../../internal/mongodb/index.js');
var server = require('../../internal/grpc/services/server/index.js');

var grpc = require('grpc');

var ProductServiceImpl = require('../../internal/grpc/services/impl/product_impl.js');
var ProductService = require('../../internal/grpc/services/product.js');

var dbConfig = {
    Username: envVariablesMap.get('DB_USERNAME'),
    Password: envVariablesMap.get('DB_PASSWORD'),
    Host:     envVariablesMap.get('DB_HOST'),
    Port:     envVariablesMap.get('DB_PORT'),
    Name:     envVariablesMap.get('DB_NAME')
};

var datastore = mongodb.InitializeDB(dbConfig);

var serviceServer = server.CreateServiceServer(datastore);

var grpcAddress = envVariablesMap.get("GRPC_SERVER_HOST") + ":" + envVariablesMap.get("GRPC_SERVER_PORT");

console.log('Starting the GRPC server connection on', grpcAddress);

var grpcServer = new grpc.Server();

var ProductServiceServer = ProductServiceImpl.NewProductServiceServer(serviceServer);
grpcServer.addService(ProductService.serviceDescriptor, ProductServiceServer);

grpcServer.bind(grpcAddress, grpc.ServerCredentials.createInsecure());

grpcServer.start();