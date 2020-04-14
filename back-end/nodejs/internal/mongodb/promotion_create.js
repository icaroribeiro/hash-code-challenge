var Promotion = require('../models/promotion.js');

function CreatePromotion(promotion, callback) {
    Promotion.create(promotion)
        .then((doc) => {
            if (doc) {
                callback(null, {
                        id: doc._id,
                        code: doc.code,
                        title: doc.title,
                        description: doc.description,
                        max_discount_pct: doc.max_discount_pct
                    }
                )
            } else {
                callback(null, null);
            }
        })
        .catch(err => {
            callback(err, null);
        });
}

exports.CreatePromotion = CreatePromotion;