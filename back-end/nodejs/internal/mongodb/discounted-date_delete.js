var DiscountedDate = require('../models/discounted-date.js');

async function DeleteDiscountedDate(id, callback) {
    await DiscountedDate.deleteOne({_id: id})
        .exec()
        .then((doc) => {
            callback(null, doc)
        })
        .catch(err => {
            callback(err, null);
        });
}

exports.DeleteDiscountedDate = DeleteDiscountedDate;