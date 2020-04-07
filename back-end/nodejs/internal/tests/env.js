const dotenv = require('dotenv').config({path: __dirname + '/.env'});

module.exports = {
    GRPC_SERVER_HOST: process.env.GRPC_SERVER_HOST,
    GRPC_SERVER_PORT: process.env.GRPC_SERVER_PORT,
    DB_USERNAME: process.env.DB_USERNAME,
    DB_PASSWORD: process.env.DB_PASSWORD,    
    DB_HOST: process.env.DB_HOST,
    DB_PORT: process.env.DB_PORT,
    DB_NAME: process.env.DB_NAME
};