var DiscountedDate = require('../models/discounted-date.js');

function DeleteDiscountedDate(id, callback) {
    DiscountedDate.deleteOne({_id: id})
        .exec()
        .then((doc) => {
            callback(null, doc)
        })
        .catch(err => {
            callback(err, null);
        });
}

exports.DeleteDiscountedDate = DeleteDiscountedDate;