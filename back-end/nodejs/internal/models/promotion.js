var mongoose = require('mongoose');

var schema = new mongoose.Schema({
    _id: mongoose.Schema.Types.ObjectId,
    code: { 
        type: String, 
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
    max_discount_pct: {
        type: Number, 
        required: true 
    },
    products: {
        type: [{
            type: String
        }],
        require: false
    }
});

var promotion = new mongoose.model('Promotion', schema);

module.exports = promotion;