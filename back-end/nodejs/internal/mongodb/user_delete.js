var User = require('../models/user.js');

async function DeleteUser(id, callback) {
    await User.deleteOne({_id: id})
        .exec()
        .then((doc) => {
            callback(null, doc)
        })
        .catch(err => {
            callback(err, null);
        });
}

exports.DeleteUser = DeleteUser;