const dotenv = require('dotenv').config({path: __dirname + '/.env'});

module.exports = {
    TEST_GRPC_SERVER_HOST: process.env.TEST_GRPC_SERVER_HOST,
    TEST_GRPC_SERVER_PORT: process.env.TEST_GRPC_SERVER_PORT,
    TEST_DB_USERNAME: process.env.TEST_DB_USERNAME,
    TEST_DB_PASSWORD: process.env.TEST_DB_PASSWORD,    
    TEST_DB_HOST: process.env.TEST_DB_HOST,    
    TEST_DB_PORT: process.env.TEST_DB_PORT,    
    TEST_DB_NAME: process.env.TEST_DB_NAME
};