package mongodb

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// The DBConfig stores all the parameters to configure the database settings.
type DBConfig struct {
    DBUsername string
    DBPassword string
    DBHost     string
    DBPort     string
    DBName     string
}

// The Datastore groups all the variables necessary to connect and work with data
// by means of a collection of statements used to interface with our backing database.
type Datastore struct {
    DB      *mongo.Database
    Client  *mongo.Client
    Context context.Context
}

func InitializeDB(dbConfig DBConfig) (Datastore, error) {
    var authCredential options.Credential
    var connString string
    var err error
    var client *mongo.Client
    var ctx context.Context
    var db *mongo.Database

    // Generate an authentication credential.
    authCredential, err = GenerateAuthCredential(dbConfig.DBUsername, dbConfig.DBPassword)

    if err != nil {
        return Datastore{}, err
    }

    // Set up the connection string of the database.
    connString = SetUpConnString(dbConfig.DBHost, dbConfig.DBPort)

    ctx = context.Background()

    client, err = mongo.Connect(ctx, options.Client().SetAuth(authCredential).ApplyURI(connString))

    if err != nil {
        return Datastore{}, err
    }

    err = client.Ping(ctx, nil)

    if err != nil {
        return Datastore{}, err
    }

    db = client.Database(dbConfig.DBName)

    return Datastore{DB: db, Client: client, Context: ctx}, nil
}

// It defines a credential containing options for configuring authentication.
func GenerateAuthCredential(dbUsername, dbPassword string) (options.Credential, error) {
    var authCredential options.Credential

    authCredential = options.Credential{
        Username: dbUsername,
        Password: dbPassword,
    }

    return authCredential, nil
}

// It builds the connection string of the database.
func SetUpConnString(dbHost, dbPort string) string {
    var connString string

    connString = fmt.Sprintf("mongodb://%s:%s",
        dbHost,
        dbPort,
    )

    return connString
}

// It closes the sockets to the topology referenced by this Client.
func (d *Datastore) Close() error {
    return d.Client.Disconnect(d.Context)
}
