var mongoose = require('mongoose');

var fileMap = require('require-all')({
    dirname : __dirname,
    filter  : function (filename) {
        if (filename.includes('index')) {
            return;
        }
        return filename;
    }
});

function InitializeDB(dbConfig) {
    var URL = "mongodb://";
        URL += dbConfig.Username + ":";
        URL += dbConfig.Password + "@";
        URL += dbConfig.Host + ":";
        URL += dbConfig.Port + "/";
        URL += dbConfig.Name;
        URL += "?authSource=admin";

    mongoose.connect(URL, {
            useNewUrlParser: true, 
            useUnifiedTopology: true 
        });

    var db = mongoose.connection;

    db.on('error', err => {
        console.log("Failed to connect to the database:", err);
        process.exit();
    });

    var datastore = {};

    for (const i in fileMap) {
        for (const j in fileMap[i]) {
            datastore[j] = fileMap[i][j];
        }
    }

    return datastore;
}

function Close() {
    mongoose.connection.close();
}

exports.InitializeDB = InitializeDB;
exports.Close = Close;