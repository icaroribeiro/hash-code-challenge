var ProductDBService = {};

var mongoose = require('mongoose');

var Promotion = require('../../../services/promotions/index.js');

var fileMap = require('require-all')({
    dirname : __dirname + "/../../../mongodb/",
    filter  : function (filename) {
        if (filename.includes('product')) {
            return filename;
        }
    }
});

for (const i in fileMap) {
    for (const j in fileMap[i]) {
        ProductDBService[j] = fileMap[i][j];
    }
}

function GetAllProducts(call, callback) {
    ProductDBService.GetAllProducts(async function (err, data) {
        if (err) {
            callback(err, []);
            return;
        }

        var xUserId = "";

        var params = call.metadata.get("X-USER-ID");

        if (params.length != 0) {
            xUserId = params[0];

            if (!mongoose.Types.ObjectId.isValid(xUserId)) {
                xUserId = "";
            }
        }

        await Promotion.EvaluatePromotions(data, xUserId);

        callback(null, data);
    });
}

function GetProduct(call, callback) {
    ProductDBService.GetProduct(call.request.id, async function (err, data) {
        if (err) {
            callback(err, []);
            return;
        }

        var xUserId = "";

        var params = call.metadata.get("X-USER-ID");

        if (params.length != 0) {
            xUserId = params[0];

            if (!mongoose.Types.ObjectId.isValid(xUserId)) {
                xUserId = "";
            }
        }

        await Promotion.EvaluatePromotions(data, xUserId);

        callback(null, data);
    });
}

exports.GetAllProducts = GetAllProducts;

exports.GetProduct = GetProduct;