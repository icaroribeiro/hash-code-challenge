package impl_test

import (
    "encoding/json"
    "fmt"
    "github.com/golang/protobuf/proto"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/entities"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/utils"
    date "google.golang.org/genproto/googleapis/type/date"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "testing"
    "time"
)

func TestDeleteUser(t *testing.T) {
    var rDate time.Time
    var user models.User
    var body string
    var err error
    var bodyBytes []byte
    var userEntity entities.User
    var request services.DeleteUserRequest
    var response *entities.User
    var errStatus *status.Status
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

    userEntity = entities.User{
        Id:        user.ID.Hex(),
        FirstName: user.FirstName,
        LastName:  user.LastName,
        DateOfBirth: &date.Date{
            Year:  int32(user.DateOfBirth.Year),
            Month: int32(user.DateOfBirth.Month),
            Day:   int32(user.DateOfBirth.Day),
        },
    }

    request = services.DeleteUserRequest{
        Id: userEntity.Id,
    }

    response, err = userServiceClient.DeleteUser(ctx, &request)

    errStatus = status.Convert(err)

    if errStatus != nil {
        t.Errorf(`Test failed, response: code=%d and body={"error":"%s","code":%d,"message":"%s"}`,
            errStatus.Code(), errStatus.Message(), errStatus.Code(), errStatus.Message())
        return
    }

    bodyBytes, err = json.Marshal(userEntity)

    if err != nil {
        t.Fatalf("Failed to obtain the JSON encoding of the user %+v: %s", userEntity, err.Error())
    }

    // Evaluate the equality of the simulated data with those returned from the associated functionality.
    if !proto.Equal(&userEntity, response) {
        bodyBytesAux, err = json.Marshal(response)

        if err != nil {
            t.Fatalf("Failed to obtain the JSON encoding of the returned user %+v: %s", response, err.Error())
        }

        t.Errorf("Test failed, the expected user returned: %s, got: %s", string(bodyBytes), string(bodyBytesAux))
        return
    }

    t.Logf("Test successful, response: code=%d and body=%s", codes.OK, string(bodyBytes))
}
