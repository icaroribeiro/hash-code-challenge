var mongoose = require('mongoose');

var schema = new mongoose.Schema({
    _id: mongoose.Schema.Types.ObjectId,
    price_in_cents: { 
        type: Number, 
        required: true 
    },
    title: { 
        type: String, 
        required: true 
    },
    description: { 
        type: String, 
        required: true 
    },
    discount: {
        type: {
            pct: { 
                type: Number, 
                required: true 
            },
            value_in_cents: { 
                type: {
                    value: {
                        type: Number,
                        required: true
                    }
                }, 
                required: true 
            },
        },
        required: false
    }
});

var product = new mongoose.model('Product', schema);

module.exports = product;