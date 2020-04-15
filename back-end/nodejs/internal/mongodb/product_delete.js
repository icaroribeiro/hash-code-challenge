var Product = require('../models/product.js');

async function DeleteProduct(id, callback) {
    await Product.deleteOne({_id: id})
        .exec()
        .then((doc) => {
            callback(null, doc)
        })
        .catch(err => {
            callback(err, null);
        });
}

exports.DeleteProduct = DeleteProduct;