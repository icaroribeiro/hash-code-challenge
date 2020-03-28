package impl_test

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/entities"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/utils"
	"google.golang.org/grpc/status"
	"testing"
)

func TestGetAllPromotions(t *testing.T) {
	var product models.Product
	var body string
	var err error
	var promotion models.Promotion
	var bodyBytes []byte
	var promotionEntity entities.Promotion
	var request empty.Empty
	var response *services.GetAllPromotionsResponse
	var errStatus *status.Status
	var isFound bool
	var promotionEntityAux *entities.Promotion

	product = models.Product{
		PriceInCents: utils.GenerateRandomInteger(1, 1000),
		Title:        utils.GenerateRandomString(10),
		Description:  utils.GenerateRandomString(10),
	}

	body = fmt.Sprintf(`{
		"price_in_cents":%d,
		"title":"%s",
		"description":"%s"
	}`, product.PriceInCents, product.Title, product.Description)

	body = utils.RemoveEscapeSequences(body, "\t", "\n")

	t.Logf("Product: %s", body)

	product, err = datastore.CreateProduct(product)

	if err != nil {
		t.Fatalf("Failed to create a new product with %s: %s", body, err.Error())
	}

	promotion = models.Promotion{
		Code:           utils.GenerateRandomString(10),
		Title:          utils.GenerateRandomString(10),
		Description:    utils.GenerateRandomString(10),
		MaxDiscountPct: float32(utils.GenerateRandomInteger(1, 1000)) / 2.0,
		Products:       []string{product.ID.Hex()},
	}

	body = fmt.Sprintf(`{
		"code":"%s",
		"title":"%s",
		"description":"%s",
		"max_discount_pct":%f
		"products":["%s"]
	}`, promotion.Code, promotion.Title, promotion.Description, promotion.MaxDiscountPct, promotion.Products[0])

	body = utils.RemoveEscapeSequences(body, "\t", "\n")

	promotion, err = datastore.CreatePromotion(promotion)

	if err != nil {
		t.Fatalf("Failed to create a new promotion with %s: %s", body, err.Error())
	}

	bodyBytes, err = json.Marshal(promotion)

	if err != nil {
		t.Fatalf("Failed to obtain the JSON encoding of the promotion %+v: %s", promotion, err.Error())
	}

	t.Logf("Promotion: %s", string(bodyBytes))

	promotionEntity = entities.Promotion{
		Id:				promotion.ID.Hex(),
		Code:           promotion.Code,
		Title:          promotion.Title,
		Description:    promotion.Description,
		MaxDiscountPct: float32(promotion.MaxDiscountPct),
		Products: []string{product.ID.Hex()},
	}

	request = empty.Empty{}

	response, err = promotionServiceClient.GetAllPromotions(ctx, &request)

	errStatus = status.Convert(err)

	if errStatus != nil {
		t.Errorf("Test failed, response: code=%d and body={\"error\":\"%s\",\"code\":%d,\"message\":\"%s\"}",
			errStatus.Code(), errStatus.Message(), errStatus.Code(), errStatus.Message())
	}

	bodyBytes, err = json.Marshal(promotionEntity)

	if err != nil {
		t.Fatalf("Failed to obtain the JSON encoding of the promotion %+v: %s", promotionEntity, err.Error())
	}

	isFound = false

	for _, promotionEntityAux = range response.Promotions {
		if proto.Equal(&promotionEntity, promotionEntityAux) {
			isFound = true
			break
		}
	}

	if !isFound {
		t.Errorf("Test failed, the promotion wasn't found: %s", string(bodyBytes))
		return
	}

	t.Logf("Test successful, the promotion was found in the response body: code=%d and body=%s", 0, string(bodyBytes))
}

func TestGetPromotion(t *testing.T) {
	var product models.Product
	var body string
	var err error
	var promotion models.Promotion
	var bodyBytes []byte
	var promotionEntity entities.Promotion
	var request services.GetPromotionRequest
	var response *entities.Promotion
	var errStatus *status.Status
	var bodyBytesAux []byte

	product = models.Product{
		PriceInCents: utils.GenerateRandomInteger(1, 1000),
		Title:        utils.GenerateRandomString(10),
		Description:  utils.GenerateRandomString(10),
	}

	body = fmt.Sprintf(`{
		"price_in_cents":%d,
		"title":"%s",
		"description":"%s"
	}`, product.PriceInCents, product.Title, product.Description)

	body = utils.RemoveEscapeSequences(body, "\t", "\n")

	t.Logf("Product: %s", body)

	product, err = datastore.CreateProduct(product)

	if err != nil {
		t.Fatalf("Failed to create a new product with %s: %s", body, err.Error())
	}

	promotion = models.Promotion{
		Code:           utils.GenerateRandomString(10),
		Title:          utils.GenerateRandomString(10),
		Description:    utils.GenerateRandomString(10),
		MaxDiscountPct: float32(utils.GenerateRandomInteger(1, 1000)) / 2.0,
		Products:       []string{product.ID.Hex()},
	}

	body = fmt.Sprintf(`{
		"code":"%s",
		"title":"%s",
		"description":"%s",
		"max_discount_pct":%f
		"products":["%s"]
	}`, promotion.Code, promotion.Title, promotion.Description, promotion.MaxDiscountPct, promotion.Products[0])

	body = utils.RemoveEscapeSequences(body, "\t", "\n")

	promotion, err = datastore.CreatePromotion(promotion)

	if err != nil {
		t.Fatalf("Failed to create a new promotion with %s: %s", body, err.Error())
	}

	bodyBytes, err = json.Marshal(promotion)

	if err != nil {
		t.Fatalf("Failed to obtain the JSON encoding of the promotion %+v: %s", promotion, err.Error())
	}

	t.Logf("Promotion: %s", string(bodyBytes))

	promotionEntity = entities.Promotion{
		Id:				promotion.ID.Hex(),
		Code:           promotion.Code,
		Title:          promotion.Title,
		Description:    promotion.Description,
		MaxDiscountPct: float32(promotion.MaxDiscountPct),
		Products: []string{product.ID.Hex()},
	}

	request = services.GetPromotionRequest{
		Id: promotionEntity.Id,
	}

	response, err = promotionServiceClient.GetPromotion(ctx, &request)

	errStatus = status.Convert(err)

	if errStatus != nil {
		t.Fatalf("Test failed, response: code=%d and body={\"error\":\"%s\", \"code\": %d, \"message\": \"%s\"}",
			errStatus.Code(), errStatus.Message(), errStatus.Code(), errStatus.Message())
	}

	bodyBytes, err = json.Marshal(promotionEntity)

	if err != nil {
		t.Fatalf("Failed to obtain the JSON encoding of the promotion %+v: %s", promotionEntity, err.Error())
	}

	if !proto.Equal(&promotionEntity, response) {
		bodyBytesAux, err = json.Marshal(response)

		if err != nil {
			t.Fatalf("Failed to obtain the JSON encoding of the returned promotion %+v: %s", response, err.Error())
		}

		t.Errorf("Test failed, the expected promotion returned: %s, got: %s", string(bodyBytes), string(bodyBytesAux))
		return
	}

	t.Logf("Test successful, response: code=%d and body=%s", 0, string(bodyBytes))
}
