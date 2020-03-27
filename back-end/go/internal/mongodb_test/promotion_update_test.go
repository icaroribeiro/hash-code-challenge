package mongodb_test

import (
	"encoding/json"
	"fmt"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/utils"
	"testing"
)

func TestUpdatePromotion(t *testing.T) {
	var product models.Product
	var body string
	var err error
	var promotion models.Promotion
	var bodyBytes []byte
	var nMatchedDocs int64
	var nModifiedDocs int64

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

	promotion = models.Promotion{
		ID:             promotion.ID,
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

	t.Logf("Update promotion: %s", body)

	nMatchedDocs, nModifiedDocs, err = datastore.UpdatePromotion(promotion.ID.Hex(), promotion)

	if err != nil {
		t.Fatalf("Failed to update the promotion with the id %s with %s: %s", promotion.ID.Hex(), body, err.Error())
	}

	if nMatchedDocs == 0 {
		t.Errorf("Test failed, the id wasn't found")
		return
	}

	if nModifiedDocs == 0 {
		t.Errorf("Test failed, the data sent is already registered")
	}

	if nModifiedDocs != 1 {
		t.Errorf("Test failed, the expected number of promotions updated: %d, got: %d", 1, nModifiedDocs)
		return
	}

	bodyBytes, err = json.Marshal(promotion)

	if err != nil {
		t.Fatalf("Failed to obtain the JSON encoding of the promotion %+v: %s", promotion, err.Error())
	}

	t.Logf("Test successful, the updated promotion: %s", string(bodyBytes))
}
