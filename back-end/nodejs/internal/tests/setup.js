var grpc = require('grpc');
const { GRPC_SERVER_HOST, GRPC_SERVER_PORT } = require('../tests/env.js');
var server = new grpc.Server();

var ProductService = require('../../internal/grpc/services/product.js');
var ProductServiceImpl = require('../../internal/grpc/services/impl/product_impl.js');

const { TEST_DB_USERNAME, TEST_DB_PASSWORD, TEST_DB_HOST, TEST_DB_PORT, TEST_DB_NAME } = require('../tests/env.js');
const mongoose = require('mongoose');

beforeAll(async () => {
    var URL = `mongodb://${TEST_DB_USERNAME}:${TEST_DB_PASSWORD}@${TEST_DB_HOST}:${TEST_DB_PORT}/${TEST_DB_NAME}?authSource=admin`;

    mongoose.connect(URL, {
            useNewUrlParser: true, 
            useUnifiedTopology: true 
        });
    
    var db = mongoose.connection;
    
    db.on('error', err => {
        console.log("Failed to connect to the database:", err);
    });

    try {
        var grpcHost = GRPC_SERVER_HOST;
    
        if (!grpcHost) {
            throw "Failed to read the GRPC_SERVER_HOST environment variable: it isn't set";
        }
    
        var grpcPort = GRPC_SERVER_PORT;
    
        if (!grpcPort) {
            throw "Failed to read the GRPC_SERVER_PORT environment variable: it isn't set";
        }
    }
    catch (err) {
        console.log(err);
        throw err;
    }
    
    var grpcAddress = grpcHost + ":" + grpcPort;

    server.addService(ProductService.serviceDescriptor, ProductServiceImpl.serviceMap);

    server.bind(grpcAddress, grpc.ServerCredentials.createInsecure());

    server.start();
});

afterAll(async () => {
    await mongoose.connection.close();
    server.forceShutdown();
});