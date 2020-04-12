const dotenv = require('dotenv').config({path: __dirname + '/.env'});

var envVariablesMap = new Map();

try {
    // The environment variables related to the database settings.
    envVariablesMap.set("DB_USERNAME", process.env.DB_USERNAME);
    envVariablesMap.set("DB_PASSWORD", process.env.DB_PASSWORD);
    envVariablesMap.set("DB_HOST", process.env.DB_HOST);
    envVariablesMap.set("DB_PORT", process.env.DB_PORT);
    envVariablesMap.set("DB_NAME", process.env.DB_NAME);

    // The environment variables related to the gRPC server.
    envVariablesMap.set("GRPC_SERVER_HOST", process.env.GRPC_SERVER_HOST);
    envVariablesMap.set("GRPC_SERVER_PORT", process.env.GRPC_SERVER_PORT);

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