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

func TestCreateDiscountedDate(t *testing.T) {
	var rDate time.Time
	var discountedDate models.DiscountedDate
	var body string
	var err error
	var discountedDateAux models.DiscountedDate
	var bodyBytes []byte
	var bodyBytesAux []byte

	rDate = utils.GenerateRandomDate(2010, 2020)

	discountedDate = models.DiscountedDate{
		Title:       utils.GenerateRandomString(10),
		Description: utils.GenerateRandomString(10),
		DiscountPct: float32(utils.GenerateRandomInteger(1, 1000)) / 2.0,
		Date: models.Date{
			Year:  rDate.Year(),
			Month: int(rDate.Month()),
			Day:   rDate.Day(),
		},
	}

	body = fmt.Sprintf(`{
			"title":"%s",
			"description":"%s",
			"discount_pct":%f,
			"date":{
				"year":%d,
				"month":%d,
				"day":%d
			}
		}`, discountedDate.Title, discountedDate.Description, discountedDate.DiscountPct,
		discountedDate.Date.Year, discountedDate.Date.Month, discountedDate.Date.Day)

	body = utils.RemoveEscapeSequencesFromString(body, "\t", "\n")

	t.Logf("Discounted date: %s", body)

	discountedDateAux, err = datastore.CreateDiscountedDate(discountedDate)

	if err != nil {
		t.Fatalf("Failed to create a new discounted date with %s: %s", body, err.Error())
	}

	discountedDate.ID = discountedDateAux.ID

	bodyBytes, err = json.Marshal(discountedDate)

	if err != nil {
		t.Fatalf("Failed to obtain the JSON encoding of the discounted date %+v: %s", discountedDate, err.Error())
	}

	// Evaluate the equality of the simulated data with those returned from the associated functionality.
	if !cmp.Equal(discountedDate, discountedDateAux) {
		bodyBytesAux, err = json.Marshal(discountedDateAux)

		if err != nil {
			t.Fatalf("Failed to obtain the JSON encoding of the returned discounted date %+v: %s", discountedDateAux, err.Error())
		}

		t.Errorf("Test failed, the expected discounted date returned: %s, got: %s", string(bodyBytes), string(bodyBytesAux))
		return
	}

	t.Logf("Test successful, the created discounted date: %s", string(bodyBytes))
}
