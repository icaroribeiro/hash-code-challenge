var ProductService = {};

var fileMap = require('require-all')({
    dirname : __dirname,
    filter  : function (fileName) {
        if (fileName.includes('product')) {
            if (fileName.includes('impl')) {
                return;
            }
            return fileName;
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