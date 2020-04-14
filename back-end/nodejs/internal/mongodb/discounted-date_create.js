var DiscountedDate = require('../models/discounted-date.js');

function CreateDiscountedDate(discountedDate, callback) {
    DiscountedDate.create(discountedDate)
        .then((doc) => {
            if (doc) {
                callback(null, {
                        id: doc._id,
                        title: doc.title,
                        description: doc.description,
                        discount_pct: doc.discount_pct,
                        date: doc.date
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

exports.CreateDiscountedDate = CreateDiscountedDate;