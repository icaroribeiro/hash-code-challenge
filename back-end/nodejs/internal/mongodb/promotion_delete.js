var Promotion = require('../models/promotion.js');

async function DeletePromotion(id, callback) {
    await Promotion.deleteOne({_id: id})
        .exec()
        .then((doc) => {
            callback(null, doc)
        })
        .catch(err => {
            callback(err, null);
        });
}

exports.DeletePromotion = DeletePromotion;