var mongoose = require('mongoose');
var Promotion = require('../../../services/promotions/index.js');

function GetAllProducts(s) {
    return function (call, callback) {
        s.ServiceServer.Datastore.GetAllProducts(async function (err, data) {
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
}

function GetProduct(s) {
    return function (call, callback) {
        s.ServiceServer.Datastore.GetProduct(call.request.id, async function (err, data) {
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
}

exports.GetAllProducts = GetAllProducts;
exports.GetProduct = GetProduct;