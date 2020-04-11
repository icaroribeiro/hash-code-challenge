package mongodb_test

import (
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/mongodb"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/utils"
    "github.com/joho/godotenv"
    "log"
    "os"
    "testing"
)

var envVariablesMap map[string]string

var datastore mongodb.Datastore

func init() {
    var err error

    err = godotenv.Load("../../.test.env")

    if err != nil {
        log.Fatalf("Failed to load the .test.env file: %s", err.Error())
    }

    envVariablesMap = make(map[string]string)

    envVariablesMap["TEST_DB_USERNAME"] = ""
    envVariablesMap["TEST_DB_PASSWORD"] = ""
    envVariablesMap["TEST_DB_HOST"] = ""
    envVariablesMap["TEST_DB_PORT"] = ""
    envVariablesMap["TEST_DB_NAME"] = ""

    err = utils.GetEnvVariables(envVariablesMap)

    if err != nil {
        log.Fatal(err.Error())
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
    var dbConfig mongodb.DBConfig
    var err error

    dbConfig = mongodb.DBConfig{
        Username: envVariablesMap["TEST_DB_USERNAME"],
        Password: envVariablesMap["TEST_DB_PASSWORD"],
        Host:     envVariablesMap["TEST_DB_HOST"],
        Port:     envVariablesMap["TEST_DB_PORT"],
        Name:     envVariablesMap["TEST_DB_NAME"],
    }

    datastore, err = mongodb.InitializeDB(dbConfig)

    if err != nil {
        log.Printf("Failed to configure the database: %s", err.Error())
        return 1
    }

    return m.Run()
}
