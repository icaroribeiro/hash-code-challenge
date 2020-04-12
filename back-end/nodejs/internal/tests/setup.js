var envVariablesMap = require('../../env.js');
const mongoose = require('mongoose');

var grpc = require('grpc');
var server = new grpc.Server();

var ProductService = require('../../internal/grpc/services/product.js');
var ProductServiceImpl = require('../../internal/grpc/services/impl/product_impl.js');

beforeAll(async () => {
    var URL = "mongodb://";
        URL += envVariablesMap.get('TEST_DB_USERNAME') + ":";
        URL += envVariablesMap.get('TEST_DB_PASSWORD') + "@";
        URL += envVariablesMap.get('TEST_DB_HOST') + ":";
        URL += envVariablesMap.get('TEST_DB_PORT') + "/";
        URL += envVariablesMap.get('TEST_DB_NAME');
        URL += ":?authSource=admin";

    mongoose.connect(URL, {
            useNewUrlParser: true, 
            useUnifiedTopology: true 
        });
    
    var db = mongoose.connection;
    
    db.on('error', err => {
        console.log("Failed to connect to the database:", err);
    });
    
    var grpcAddress = envVariablesMap.get("TEST_GRPC_SERVER_HOST") + ":" + envVariablesMap.get("TEST_GRPC_SERVER_PORT");

    server.addService(ProductService.serviceDescriptor, ProductServiceImpl.serviceMap);

    server.bind(grpcAddress, grpc.ServerCredentials.createInsecure());

    server.start();
});

afterAll(async () => {
    await mongoose.connection.close();
    server.forceShutdown();
});