const dotenv = require('dotenv').config({path: __dirname + '/.env'});

var envVariablesMap = new Map();

try {
    envVariablesMap.set("TEST_DB_USERNAME", process.env.TEST_DB_USERNAME);
    envVariablesMap.set("TEST_DB_PASSWORD", process.env.TEST_DB_PASSWORD);
    envVariablesMap.set("TEST_DB_HOST", process.env.TEST_DB_HOST);
    envVariablesMap.set("TEST_DB_PORT", process.env.TEST_DB_PORT);
    envVariablesMap.set("TEST_DB_NAME", process.env.TEST_DB_NAME);

    envVariablesMap.set("TEST_GRPC_SERVER_HOST", process.env.TEST_GRPC_SERVER_HOST);
    envVariablesMap.set("TEST_GRPC_SERVER_PORT", process.env.TEST_GRPC_SERVER_PORT);

    for (const [key, value] of envVariablesMap.entries()) {
        if (value === undefined) {
            throw "Failed to read the " + key + " environment variable: it isn't set";
        }
    }
}
catch (err) {
    console.log(err);
    process.exit()
}

module.exports = envVariablesMap;