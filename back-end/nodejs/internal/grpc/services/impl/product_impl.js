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

function NewProductServiceServer(serviceServer) {
    var server = {
        Datastore: serviceServer.Datastore
    }

    var ProductServices = {};

    for (const i in fileMap) {
        for (const j in fileMap[i]) {
            ProductServices[j] = fileMap[i][j];
        }
    }
    
    var ProductServiceServer = {
        GetAllProducts: ProductServices.GetAllProducts(server),
        GetProduct: ProductServices.GetProduct(server)
    }

    return ProductServiceServer;
}

exports.NewProductServiceServer = NewProductServiceServer;