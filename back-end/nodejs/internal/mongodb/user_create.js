var User = require('../models/user.js');

function CreateUser(user, callback) {
    User.create(user)
        .then((doc) => {
            if (doc) {
                callback(null, {
                        id: doc._id,
                        first_name: doc.first_name,
                        last_name: doc.last_name,
                        date_of_birth: doc.date_of_birth
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

exports.CreateUser = CreateUser;