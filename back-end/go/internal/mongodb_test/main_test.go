package mongodb_test

import (
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/mongodb"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/utils"
    "github.com/joho/godotenv"
    "log"
    "os"
    "testing"
)

var datastore mongodb.Datastore

func init() {
    var err error

    err = godotenv.Load("../../.test.env")

    if err != nil {
        log.Fatalf("Failed to load the .test.env file: %s", err.Error())
    }
}

func TestMain(m *testing.M) {
    var exitVal int

    // Before the tests.
    utils.InitializeRandomization()

    exitVal = testMain(m)

    // After the tests.
    defer datastore.Close()

    os.Exit(exitVal)
}

func testMain(m *testing.M) int {
    var dbUsername string
    var dbPassword string
    var dbHost string
    var dbPort string
    var dbName string
    var isSet bool
    var dbConfig mongodb.DBConfig
    var err error

    dbUsername, isSet = os.LookupEnv("DB_USERNAME")

    if !isSet {
        log.Print("Failed to read the DB_USERNAME environment variable: it isn't set")
        return 1
    }

    dbPassword, isSet = os.LookupEnv("DB_PASSWORD")

    if !isSet {
        log.Print("Failed to read the DB_PASSWORD environment variable: it isn't set")
        return 1
    }

    dbHost, isSet = os.LookupEnv("DB_HOST")

    if !isSet {
        log.Print("Failed to read the DB_HOST environment variable: it isn't set")
        return 1
    }

    dbPort, isSet = os.LookupEnv("DB_PORT")

    if !isSet {
        log.Print("Failed to read the DB_PORT environment variable: it isn't set")
        return 1
    }

    dbName, isSet = os.LookupEnv("DB_NAME")

    if !isSet {
        log.Print("Failed to read the DB_NAME environment variable: it isn't set")
        return 1
    }

    dbConfig = mongodb.DBConfig{
        Username: dbUsername,
        Password: dbPassword,
        Host:     dbHost,
        Port:     dbPort,
        Name:     dbName,
    }

    datastore, err = mongodb.InitializeDB(dbConfig)

    if err != nil {
        log.Printf("Failed to configure the database: %s", err.Error())
        return 1
    }

    return m.Run()
}
