var grpc = require('grpc');

var protoPath = require('path').join(__dirname, '../../../internal/proto/services/product.proto');

var protoLoader = require('@grpc/proto-loader');

var packageDefinition = protoLoader.loadSync(
    protoPath,{
            keepCase: true,
            longs: String,
            enums: String,
            defaults: true,
            oneofs: true
        });

var protoDescriptor = grpc.loadPackageDefinition(packageDefinition);

exports.stubConstructor = protoDescriptor.services.ProductService;

exports.serviceDescriptor = protoDescriptor.services.ProductService.service;