var mongoose = require('mongoose');

const { DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME } = require('../../env');

var URL = `mongodb://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?authSource=admin`;

mongoose.connect(URL, {
        useNewUrlParser: true, 
        useUnifiedTopology: true 
    });

var db = mongoose.connection;

db.on('error', err => {
    console.log("Failed to connect to the database:", err);
});