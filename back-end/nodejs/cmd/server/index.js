var grpc = require('grpc');

var envVariablesMap = require('../../env.js');

var ProductService = require('../../internal/grpc/services/product.js');
var ProductServiceImpl = require('../../internal/grpc/services/impl/product_impl.js');

require('../../internal/mongodb/index.js');

var grpcAddress = envVariablesMap.get("GRPC_SERVER_HOST") + ":" + envVariablesMap.get("GRPC_SERVER_PORT");

console.log('Starting the GRPC server connection on', grpcAddress);

var server = new grpc.Server();

server.addService(ProductService.serviceDescriptor, ProductServiceImpl.serviceMap);

server.bind(grpcAddress, grpc.ServerCredentials.createInsecure());

server.start();