var mongodb = require('../../../mongodb/index.js');

// This variable is an abstraction of the server that allows to "attach" some resources in order to make them
// available during the API requests. Here, it's used to store other variable that holds attributes to manage the data.
var s = {
    Datastore: mongodb.Datastore
};

var ProductService = {};

var fileMap = require('require-all')({
    dirname : __dirname,
    filter  : function (filename) {
        if (filename.includes('product')) {
            if (filename.includes('impl')) {
                return;
            }
            return filename;
        }
    }
});

for (const i in fileMap) {
    for (const j in fileMap[i]) {
        ProductService[j] = fileMap[i][j];
    }
}

var serviceMap = {
    GetAllProducts: ProductService.GetAllProducts(s),
    GetProduct: ProductService.GetProduct(s)
}

exports.serviceMap = serviceMap;