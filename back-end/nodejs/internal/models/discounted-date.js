var mongoose = require('mongoose');

var schema = new mongoose.Schema({
    _id: mongoose.Schema.Types.ObjectId,
    title: { 
        type: String, 
        required: true 
    },
    description: { 
        type: String, 
        required: true 
    },
    discount_pct: {
        type: Number, 
        required: true 
    },
    date: { 
        type: {
            year: { 
                type: Number, 
                required: true 
            },
            month: { 
                type: Number, 
                required: true 
            },
            day: { 
                type: Number, 
                required: true 
            }
        },
        required: true 
    }
});

var discountedDate = new mongoose.model('DiscountedDate', schema, 'discountedDates');

module.exports = discountedDate;