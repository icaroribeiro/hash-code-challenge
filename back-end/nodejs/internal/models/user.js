var mongoose = require('mongoose');

var schema = new mongoose.Schema({
    _id: mongoose.Schema.Types.ObjectId,
    first_name: { 
        type: String, 
        required: true 
    },
    last_name: { 
        type: String, 
        required: true 
    },
    date_of_birth: { 
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

var user = new mongoose.model('User', schema);

module.exports = user;