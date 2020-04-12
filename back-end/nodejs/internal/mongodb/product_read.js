var Product = require('../models/product.js');

function GetAllProducts(callback) {
    Product.find()
        .exec()
        .then((docs) => {
            if (docs.length >= 0) {
                callback(null, { 
                        products: docs.map(doc => {
                            return {
                                id: doc._id,
                                price_in_cents: doc.price_in_cents,
                                title: doc.title,
                                description: doc.description
                            }
                        })           
                    }
                )
            } else {
                callback(null, null);
            }
        })
        .catch(err => {
            console.log(err);
            callback(err, null);
        });
}

function GetProduct(id, callback) {
    Product.findById(id)
        .exec()
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
            console.log(err);
            callback(err, null);
        });
}

exports.GetAllProducts = GetAllProducts;
exports.GetProduct = GetProduct;