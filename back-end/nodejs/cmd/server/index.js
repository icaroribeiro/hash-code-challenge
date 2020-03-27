var grpc = require('grpc');

const { GRPC_SERVER_HOST, GRPC_SERVER_PORT } = require('../../env.js');

var ProductService = require('../../internal/grpc/services/product.js');
var ProductServiceImpl = require('../../internal/grpc/services/impl/product_impl.js');

require('../../internal/models/product.js');

require('../../internal/mongodb/index.js');

try {
    var grpcHost = GRPC_SERVER_HOST;

    if (!grpcHost) {
        throw "Faield to read the GRPC_SERVER_HOST environment variable: it isn't set";
    }

    var grpcPort = GRPC_SERVER_PORT;

    if (!grpcPort) {
        throw "Failed to read the GRPC_SERVER_PORT environment variable: it isn't set";
    }
}
catch (err) {
    console.log(err);
    process.exit()
}

var grpcAddress = grpcHost + ":" + grpcPort;

console.log('Starting the GRPC server connection on', grpcAddress);

var server = new grpc.Server();

server.addService(ProductService.serviceDescriptor, ProductServiceImpl.serviceMap);

server.bind(grpcAddress, grpc.ServerCredentials.createInsecure());

server.start();