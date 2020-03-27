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

func TestCreateDiscountedDate(t *testing.T) {
	var rDate time.Time
	var discountedDate models.DiscountedDate
	var body string
	var discountedDateEntity entities.DiscountedDate
	var request services.CreateDiscountedDateRequest
	var response *entities.DiscountedDate
	var err error
	var errStatus *status.Status
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

	body = utils.RemoveEscapeSequences(body, "\t", "\n")

	t.Logf("Discounted date: %s", body)

	discountedDateEntity = entities.DiscountedDate{
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

	request = services.CreateDiscountedDateRequest{
		DiscountedDate: &discountedDateEntity,
	}

	response, err = discountedDateServiceClient.CreateDiscountedDate(ctx, &request)

	errStatus = status.Convert(err)

	if errStatus != nil {
		t.Errorf("Test failed, response: code=%d and body={\"error\": \"%s\", \"code\": %d, \"message\": \"%s\"}",
			errStatus.Code(), errStatus.Message(), errStatus.Code(), errStatus.Message())
	}

	discountedDateEntity.Id = response.Id

	bodyBytes, err = json.Marshal(discountedDateEntity)

	if err != nil {
		t.Fatalf("Failed to obtain the JSON encoding of the discounted date %+v: %s", discountedDateEntity, err.Error())
	}

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
