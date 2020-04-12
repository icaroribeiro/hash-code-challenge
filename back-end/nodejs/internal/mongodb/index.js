var mongoose = require('mongoose');
var envVariablesMap = require('../../env.js');

var URL = "mongodb://";
    URL += envVariablesMap.get('DB_USERNAME') + ":";
    URL += envVariablesMap.get('DB_PASSWORD') + "@";
    URL += envVariablesMap.get('DB_HOST') + ":";
    URL += envVariablesMap.get('DB_PORT') + "/";
    URL += envVariablesMap.get('DB_NAME');
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

var fileMap = require('require-all')({
    dirname : __dirname,
    filter  : function (filename) {
        if (filename.includes('index')) {
            return;
        }
        return filename;
    }
});

for (const i in fileMap) {
    for (const j in fileMap[i]) {
        datastore[j] = fileMap[i][j];
    }
}

exports.Datastore = datastore;