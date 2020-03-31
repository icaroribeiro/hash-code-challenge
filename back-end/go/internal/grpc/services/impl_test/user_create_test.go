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
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	var rDate time.Time
	var user models.User
	var body string
	var userEntity entities.User
	var request services.CreateUserRequest
	var response *entities.User
	var err error
	var errStatus *status.Status
	var bodyBytes []byte
	var bodyBytesAux []byte

	rDate = utils.GenerateRandomDate(2019, 2020)

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
			"day":%d
		}
	}`, user.FirstName, user.LastName, user.DateOfBirth.Year, user.DateOfBirth.Month, user.DateOfBirth.Day)

	body = utils.RemoveEscapeSequences(body, "\t", "\n")

	t.Logf("User: %s", body)

	userEntity = entities.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		DateOfBirth: &date.Date{
			Year:  int32(user.DateOfBirth.Year),
			Month: int32(user.DateOfBirth.Month),
			Day:   int32(user.DateOfBirth.Day),
		},
	}

	request = services.CreateUserRequest{
		User: &userEntity,
	}

	response, err = userServiceClient.CreateUser(ctx, &request)

	errStatus = status.Convert(err)

	if errStatus != nil {
		t.Errorf("Test failed, response: code=%d and body={\"error\":\"%s\",\"code\":%d,\"message\":\"%s\"}",
			errStatus.Code(), errStatus.Message(), errStatus.Code(), errStatus.Message())
	}

	userEntity.Id = response.Id

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

	t.Logf("Test successful, response: code=%d and body=%s", 0, string(bodyBytes))
}
