function CreateServiceServer(datastore) {
    // This object is an abstraction of the server that allows to "attach" some resources in order to make them
    // available during the API requests. Here, it's used to store other object that holds attributes to manage the data.
    return serviceServer = {
        Datastore: datastore,
    }
}

exports.CreateServiceServer = CreateServiceServer;