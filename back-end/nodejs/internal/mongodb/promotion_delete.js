var Promotion = require('../models/promotion.js');

function DeletePromotion(id, callback) {
    Promotion.deleteOne({_id: id})
        .exec()
        .then((doc) => {
            callback(null, doc)
        })
        .catch(err => {
            callback(err, null);
        });
}

exports.DeletePromotion = DeletePromotion;