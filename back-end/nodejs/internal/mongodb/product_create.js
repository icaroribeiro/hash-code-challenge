var Product = require('../models/product.js');

async function CreateProduct(product, callback) {
    await Product.create(product)
        .then((doc) => {
            if (doc) {
                callback(null, {
                        id: doc._id,
                        price_in_cents: doc.price_in_cents,
                        title: doc.title,
                        description: doc.description
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

exports.CreateProduct = CreateProduct;