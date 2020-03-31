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

func TestCreateUser(t *testing.T) {
	var rDate time.Time
	var user models.User
	var body string
	var err error
	var userAux models.User
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
			"day":%d,
		}
	}`, user.FirstName, user.LastName, user.DateOfBirth.Year, user.DateOfBirth.Month, user.DateOfBirth.Day)

	body = utils.RemoveEscapeSequences(body, "\t", "\n")

	t.Logf("User: %s", body)

	userAux, err = datastore.CreateUser(user)

	if err != nil {
		t.Fatalf("Failed to create a new user with %s: %s", body, err.Error())
	}

	user.ID = userAux.ID

	bodyBytes, err = json.Marshal(user)

	if err != nil {
		t.Fatalf("Failed to obtain the JSON encoding of the user %+v: %s", user, err.Error())
	}

	// Evaluate the equality of the simulated data with those returned from the associated functionality.
	if !cmp.Equal(user, userAux) {
		bodyBytesAux, err = json.Marshal(userAux)

		if err != nil {
			t.Fatalf("Failed to obtain the JSON encoding of the returned user %+v: %s", userAux, err.Error())
		}

		t.Errorf("Test failed, the expected user returned: %s, got: %s", string(bodyBytes), string(bodyBytesAux))
		return
	}

	t.Logf("Test successful, the created user: %s", string(bodyBytes))
}
