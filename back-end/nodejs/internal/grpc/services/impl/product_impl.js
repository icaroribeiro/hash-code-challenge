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
    GetAllProducts: ProductService.GetAllProducts,
    GetProduct: ProductService.GetProduct
}

exports.serviceMap = serviceMap;