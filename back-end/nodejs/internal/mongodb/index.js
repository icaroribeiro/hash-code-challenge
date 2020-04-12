var mongoose = require('mongoose');
var envVariablesMap = require('../../env.js');

var URL = "mongodb://";
    URL += envVariablesMap.get('DB_USERNAME') + ":";
    URL += envVariablesMap.get('DB_PASSWORD') + "@";
    URL += envVariablesMap.get('DB_HOST') + ":";
    URL += envVariablesMap.get('DB_PORT') + "/";
    URL += envVariablesMap.get('DB_NAME');
    URL += ":?authSource=admin";

mongoose.connect(URL, {
        useNewUrlParser: true, 
        useUnifiedTopology: true 
    });

var db = mongoose.connection;

db.on('error', err => {
    console.log("Failed to connect to the database:", err);
});