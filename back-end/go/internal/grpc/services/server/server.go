package server

import (
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/mongodb"
)

// This structure is an abstraction of the server that allows to "attach" some resources in order to make them
// available during the API requests. Here, it's used to store other structure that holds attributes to manage the data.
type ServiceServer struct {
    Datastore mongodb.Datastore
}

func CreateServiceServer(datastore mongodb.Datastore) ServiceServer {
    return ServiceServer {
        Datastore: datastore,
    }
}