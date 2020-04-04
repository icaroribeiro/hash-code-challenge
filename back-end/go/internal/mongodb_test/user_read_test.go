package mongodb_test

import (
    "encoding/json"
    "fmt"
    "github.com/google/go-cmp/cmp"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/utils"
    "testing"
    "time"
)

func TestGetAllUsers(t *testing.T) {
    var rDate time.Time
    var user models.User
    var body string
    var err error
    var bodyBytes []byte
    var users []models.User
    var isFound bool
    var userAux models.User

    rDate = utils.GenerateRandomDate(2010, 2020)

    user = models.User{
        FirstName: utils.GenerateRandomString(10),
        LastName:  utils.GenerateRandomString(10),
        DateOfBirth: models.Date{
            Year:  rDate.Year(),
            Month: int(rDate.Month()),
            Day:   rDate.Day(),
        },
    }

    body = fmt.Sprintf(`{
        "first_name":"%s",
        "last_name":"%s",
        "date_of_birth":{
            "year":%d,
            "month":%d,
            "day":%d,
        }
    }`, user.FirstName, user.LastName, user.DateOfBirth.Year, user.DateOfBirth.Month, user.DateOfBirth.Day)

    body = utils.RemoveEscapeSequencesFromString(body, "\t", "\n")

    user, err = datastore.CreateUser(user)

    if err != nil {
        t.Fatalf("Failed to create a new user with %s: %s", body, err.Error())
    }

    bodyBytes, err = json.Marshal(user)

    if err != nil {
        t.Fatalf("Failed to obtain the JSON encoding of the user %+v: %s", user, err.Error())
    }

    t.Logf("User: %s", string(bodyBytes))

    users, err = datastore.GetAllUsers()

    if err != nil {
        t.Fatalf("Failed to get the list of all users: %s", err.Error())
    }

    isFound = false

    for _, userAux = range users {
        // Evaluate the equality of the simulated data with those returned from the associated functionality.
        if cmp.Equal(user, userAux) {
            isFound = true
            break
        }
    }

    if !isFound {
        t.Errorf("Test failed, the user with the id %s wasn't found in the list of all users: %s",
            user.ID.Hex(), string(bodyBytes))
        return
    }

    t.Logf("Test successful, the user found in the list of all users: %s", string(bodyBytes))
}

func TestGetUser(t *testing.T) {
    var rDate time.Time
    var user models.User
    var body string
    var err error
    var bodyBytes []byte
    var userAux models.User
    var bodyBytesAux []byte

    rDate = utils.GenerateRandomDate(2010, 2020)

    user = models.User{
        FirstName: utils.GenerateRandomString(10),
        LastName:  utils.GenerateRandomString(10),
        DateOfBirth: models.Date{
            Year:  rDate.Year(),
            Month: int(rDate.Month()),
            Day:   rDate.Day(),
        },
    }

    body = fmt.Sprintf(`{
        "first_name":"%s",
        "last_name":"%s",
        "date_of_birth":{
            "year":%d,
            "month":%d,
            "day":%d,
        }
    }`, user.FirstName, user.LastName, user.DateOfBirth.Year, user.DateOfBirth.Month, user.DateOfBirth.Day)

    body = utils.RemoveEscapeSequencesFromString(body, "\t", "\n")

    user, err = datastore.CreateUser(user)

    if err != nil {
        t.Fatalf("Failed to create a new user with %s: %s", body, err.Error())
    }

    bodyBytes, err = json.Marshal(user)

    if err != nil {
        t.Fatalf("Failed to obtain the JSON encoding of the user %+v: %s", user, err.Error())
    }

    t.Logf("User: %s", string(bodyBytes))

    userAux, err = datastore.GetUser(user.ID.Hex())

    if err != nil {
        t.Fatalf("Failed to get the user with the id %s: %s", user.ID.Hex(), err.Error())
    }

    // Evaluate the equality of the simulated data with those returned from the associated functionality.
    if !(cmp.Equal(user, userAux)) {
        bodyBytesAux, err = json.Marshal(userAux)

        if err != nil {
            t.Fatalf("Failed to obtain the JSON encoding of the returned user %+v: %s", userAux, err.Error())
        }

        t.Errorf("Test failed, the expected user returned: %s, got: %s", string(bodyBytes), string(bodyBytesAux))
        return
    }

    t.Logf("Test successful, the returned user: %s", string(bodyBytes))
}
