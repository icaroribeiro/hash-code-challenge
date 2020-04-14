var Promotion = require('../models/product.js');

function UpdatePromotion(id, updateOptions, callback) {
    Promotion.updateOne({_id: id},{$set: updateOptions})
        .exec()
        .then((doc) => {
            callback(null, doc);
        })
        .catch(err => {
            callback(err, null);
        });
}

exports.UpdatePromotion = UpdatePromotion;