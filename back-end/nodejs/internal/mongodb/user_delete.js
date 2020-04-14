var User = require('../models/user.js');

function DeleteUser(id, callback) {
    User.deleteOne({_id: id})
        .exec()
        .then((doc) => {
            callback(null, doc)
        })
        .catch(err => {
            callback(err, null);
        });
}

exports.DeleteUser = DeleteUser;