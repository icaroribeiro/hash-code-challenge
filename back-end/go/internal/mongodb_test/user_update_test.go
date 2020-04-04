package mongodb_test

import (
    "encoding/json"
    "fmt"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/utils"
    "testing"
    "time"
)

func TestUpdateUser(t *testing.T) {
    var rDate time.Time
    var user models.User
    var body string
    var err error
    var bodyBytes []byte
    var nMatchedDocs int64
    var nModifiedDocs int64

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

    body = fmt.Sprintf(`{"first_name":"%s","last_name":"%s","date_of_birth":{"year":%d,"month":%d,"day":%d}}`,
        user.FirstName, user.LastName, user.DateOfBirth.Year, user.DateOfBirth.Month, user.DateOfBirth.Day)

    user, err = datastore.CreateUser(user)

    if err != nil {
        t.Fatalf("Failed to create a new user with %s: %s", body, err.Error())
    }

    bodyBytes, err = json.Marshal(user)

    if err != nil {
        t.Fatalf("Failed to obtain the JSON encoding of the user %+v: %s", user, err.Error())
    }

    t.Logf("User: %s", string(bodyBytes))

    rDate = utils.GenerateRandomDate(2010, 2020)

    user = models.User{
        ID:        user.ID,
        FirstName: utils.GenerateRandomString(10),
        LastName:  utils.GenerateRandomString(10),
        DateOfBirth: models.Date{
            Year:  rDate.Year(),
            Month: int(rDate.Month()),
            Day:   rDate.Day(),
        },
    }

    body = fmt.Sprintf(`{"first_name":"%s","last_name":"%s","date_of_birth":{"year":%d,"month":%d,"day":%d}}`,
        user.FirstName, user.LastName, user.DateOfBirth.Year, user.DateOfBirth.Month, user.DateOfBirth.Day)

    t.Logf("New user data: %s", body)

    nMatchedDocs, nModifiedDocs, err = datastore.UpdateUser(user.ID.Hex(), user)

    if err != nil {
        t.Fatalf("Failed to update the user with the id %s with %s: %s", user.ID.Hex(), body, err.Error())
    }

    if nMatchedDocs == 0 {
        t.Errorf("Test failed, the user with the id %s wasn't found", user.ID.Hex())
        return
    }

    if nModifiedDocs == 0 {
        t.Errorf("Test failed, the data sent are already registered")
    }

    if nModifiedDocs != 1 {
        t.Errorf("Test failed, the expected number of users updated: %d, got: %d", 1, nModifiedDocs)
        return
    }

    bodyBytes, err = json.Marshal(user)

    if err != nil {
        t.Fatalf("Failed to obtain the JSON encoding of the user %+v: %s", user, err.Error())
    }

    t.Logf("Test successful, the updated user: %s", string(bodyBytes))
}
