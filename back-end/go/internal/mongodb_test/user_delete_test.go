package mongodb_test

import (
    "encoding/json"
    "fmt"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/utils"
    "testing"
    "time"
)

func TestDeleteUser(t *testing.T) {
    var rDate time.Time
    var user models.User
    var body string
    var err error
    var bodyBytes []byte
    var nDeletedDocs int64

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

    nDeletedDocs, err = datastore.DeleteUser(user.ID.Hex())

    if err != nil {
        t.Fatalf("Failed to delete the user with the id %s: %s", user.ID.Hex(), err.Error())
    }

    if nDeletedDocs == 0 {
        t.Errorf("Test failed, the user with the id %s wasn't found", user.ID.Hex())
        return
    }

    if nDeletedDocs != 1 {
        t.Errorf("Test failed, the expected number of users deleted: %d, got: %d", 1, nDeletedDocs)
        return
    }

    t.Logf("Test successful, the deleted user: %s", string(bodyBytes))
}
