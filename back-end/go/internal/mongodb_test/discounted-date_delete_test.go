package mongodb_test

import (
    "encoding/json"
    "fmt"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/utils"
    "testing"
    "time"
)

func TestDeleteDiscountedDate(t *testing.T) {
    var rDate time.Time
    var discountedDate models.DiscountedDate
    var body string
    var err error
    var bodyBytes []byte
    var nDeletedDocs int64

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

    nDeletedDocs, err = datastore.DeleteDiscountedDate(discountedDate.ID.Hex())

    if err != nil {
        t.Fatalf("Failed to delete the discounted date with the id %s: %s", discountedDate.ID.Hex(), err.Error())
    }

    if nDeletedDocs == 0 {
        t.Errorf("Test failed, the discounted date with the id %s wasn't found", discountedDate.ID.Hex())
        return
    }

    if nDeletedDocs != 1 {
        t.Errorf("Test failed, the expected number of discounted dates deleted: %d, got: %d", 1, nDeletedDocs)
        return
    }

    t.Logf("Test successful, the deleted discounted date: %s", string(bodyBytes))
}
