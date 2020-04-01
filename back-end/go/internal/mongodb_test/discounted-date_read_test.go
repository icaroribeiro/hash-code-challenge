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

func TestGetAllDiscountedDates(t *testing.T) {
	var rDate time.Time
	var discountedDate models.DiscountedDate
	var body string
	var err error
	var bodyBytes []byte
	var discountedDates []models.DiscountedDate
	var isFound bool
	var discountedDateAux models.DiscountedDate

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

	discountedDate, err = datastore.CreateDiscountedDate(discountedDate)

	if err != nil {
		t.Fatalf("Failed to create a new discounted date with %s: %s", body, err.Error())
	}

	bodyBytes, err = json.Marshal(discountedDate)

	if err != nil {
		t.Fatalf("Failed to obtain the JSON encoding of the discounted date %+v: %s", discountedDate, err.Error())
	}

	t.Logf("Discounted date: %s", string(bodyBytes))

	discountedDates, err = datastore.GetAllDiscountedDates()

	if err != nil {
		t.Fatalf("Failed to get the list of all discounted dates: %s", err.Error())
	}

	isFound = false

	for _, discountedDateAux = range discountedDates {
		// Evaluate the equality of the simulated data with those returned from the associated functionality.
		if cmp.Equal(discountedDate, discountedDateAux) {
			isFound = true
			break
		}
	}

	if !isFound {
		t.Errorf("Test failed, the discounted date with the id %s wasn't found in the list of all discounted dates: %s", 
			discountedDate.ID.Hex(), string(bodyBytes))
		return
	}

	t.Logf("Test successful, the discounted date found in the list of all discounted dates: %s", string(bodyBytes))
}

func TestGetDiscountedDate(t *testing.T) {
	var rDate time.Time
	var discountedDate models.DiscountedDate
	var body string
	var err error
	var bodyBytes []byte
	var discountedDateAux models.DiscountedDate
	var bodyBytesAux []byte

	rDate = utils.GenerateRandomDate(2019, 2020)

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

	discountedDate, err = datastore.CreateDiscountedDate(discountedDate)

	if err != nil {
		t.Fatalf("Failed to create a new discounted date with %s: %s", body, err.Error())
	}

	bodyBytes, err = json.Marshal(discountedDate)

	if err != nil {
		t.Fatalf("Failed to obtain the JSON encoding of the discounted date %+v: %s", discountedDate, err.Error())
	}

	t.Logf("Discounted date: %s", string(bodyBytes))

	discountedDateAux, err = datastore.GetDiscountedDate(discountedDate.ID.Hex())

	if err != nil {
		t.Fatalf("Failed to get the discounted date with the id %s: %s", discountedDate.ID.Hex(), err.Error())
	}

	// Evaluate the equality of the simulated data with those returned from the associated functionality.
	if !(cmp.Equal(discountedDate, discountedDateAux)) {
		bodyBytesAux, err = json.Marshal(discountedDateAux)

		if err != nil {
			t.Fatalf("Failed to obtain the JSON encoding of the returned discounted date %+v: %s",
				discountedDateAux, err.Error())
		}

		t.Errorf("Test failed, the expected discounted date returned: %s, got: %s", string(bodyBytes), string(bodyBytesAux))
		return
	}

	t.Logf("Test successful, the returned discounted date: %s", string(bodyBytes))
}
