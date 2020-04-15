var ProductServices = {};

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
        ProductServices[j] = fileMap[i][j];
    }
}

function NewProductServiceServer(serviceServer) {
    var server = {
        ServiceServer: serviceServer
    }

    return ProductServiceServer = {
        GetAllProducts: ProductServices.GetAllProducts(server),
        GetProduct: ProductServices.GetProduct(server)
    }
}

exports.NewProductServiceServer = NewProductServiceServer;