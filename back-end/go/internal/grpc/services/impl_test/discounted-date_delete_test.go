package impl_test

import (
    "encoding/json"
    "fmt"
    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/ptypes/wrappers"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/entities"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/utils"
    "google.golang.org/grpc/status"
    "testing"
    "time"
)

func TestDeleteDiscountedDate(t *testing.T) {
    var rDate time.Time
    var discountedDate models.DiscountedDate
    var body string
    var err error
    var bodyBytes []byte
    var discountedDateEntity entities.DiscountedDate
    var request services.DeleteDiscountedDateRequest
    var response *entities.DiscountedDate
    var errStatus *status.Status
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

    body = fmt.Sprintf(`{"title":"%s","description":"%s","discount_pct":%f,"date":{"year":%d,"month":%d,"day":%d}}`,
        discountedDate.Title, discountedDate.Description, discountedDate.DiscountPct,
        discountedDate.Date.Year, discountedDate.Date.Month, discountedDate.Date.Day)

    discountedDate, err = datastore.CreateDiscountedDate(discountedDate)

    if err != nil {
        t.Fatalf("Failed to create a new discounted date with %s: %s", body, err.Error())
    }

    bodyBytes, err = json.Marshal(discountedDate)

    if err != nil {
        t.Fatalf("Failed to obtain the JSON encoding of the discounted date %+v: %s", discountedDate, err.Error())
    }

    t.Logf("Discounted date: %s", string(bodyBytes))

    discountedDateEntity = entities.DiscountedDate{
        Id:          discountedDate.ID.Hex(),
        Title:       discountedDate.Title,
        Description: discountedDate.Description,
        DiscountPct: float32(discountedDate.DiscountPct),
        Date: &entities.Date{
            Year: &wrappers.Int32Value{
                Value: int32(discountedDate.Date.Year),
            },
            Month: &wrappers.Int32Value{
                Value: int32(discountedDate.Date.Month),
            },
            Day: &wrappers.Int32Value{
                Value: int32(discountedDate.Date.Day),
            },
        },
    }

    request = services.DeleteDiscountedDateRequest{
        Id: discountedDateEntity.Id,
    }

    response, err = discountedDateServiceClient.DeleteDiscountedDate(ctx, &request)

    errStatus = status.Convert(err)

    if errStatus != nil {
        t.Errorf("Test failed, response: code=%d and body={\"error\":\"%s\",\"code\":%d,\"message\":\"%s\"}",
            errStatus.Code(), errStatus.Message(), errStatus.Code(), errStatus.Message())
    }

    bodyBytes, err = json.Marshal(discountedDateEntity)

    if err != nil {
        t.Fatalf("Failed to obtain the JSON encoding of the discounted date %+v: %s", discountedDateEntity, err.Error())
    }

    // Evaluate the equality of the simulated data with those returned from the associated functionality.
    if !proto.Equal(&discountedDateEntity, response) {
        bodyBytesAux, err = json.Marshal(response)

        if err != nil {
            t.Fatalf("Failed to obtain the JSON encoding of the returned discounted date %+v: %s",
                response, err.Error())
        }

        t.Errorf("Test failed, the expected discounted date returned: %s, got: %s",
            string(bodyBytes), string(bodyBytesAux))
        return
    }

    t.Logf("Test successful, response: code=%d and body=%s", 0, string(bodyBytes))
}
