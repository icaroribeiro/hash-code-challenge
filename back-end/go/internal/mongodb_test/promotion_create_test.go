package mongodb_test

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/utils"
	"testing"
)

func TestCreatePromotion(t *testing.T) {
	var product models.Product
	var body string
	var err error
	var promotion models.Promotion
	var promotionAux models.Promotion
	var bodyBytes []byte
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

	body = utils.RemoveEscapeSequencesFromString(body, "\t", "\n")

	product, err = datastore.CreateProduct(product)

	if err != nil {
		t.Fatalf("Failed to create a new product with %s: %s", body, err.Error())
	}

	bodyBytes, err = json.Marshal(product)

	if err != nil {
		t.Fatalf("Failed to obtain the JSON encoding of the product %+v: %s", product, err.Error())
	}

	t.Logf("Product: %s", string(bodyBytes))

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
		"max_discount_pct":%f,
		"products":["%s"]
	}`, promotion.Code, promotion.Title, promotion.Description, promotion.MaxDiscountPct, promotion.Products[0])

	body = utils.RemoveEscapeSequencesFromString(body, "\t", "\n")

	t.Logf("Promotion: %s", body)

	promotionAux, err = datastore.CreatePromotion(promotion)

	if err != nil {
		t.Fatalf("Failed to create a new promotion with %s: %s", body, err.Error())
	}

	promotion.ID = promotionAux.ID

	bodyBytes, err = json.Marshal(promotion)

	if err != nil {
		t.Fatalf("Failed to obtain the JSON encoding of the promotion %+v: %s", promotion, err.Error())
	}

	// Evaluate the equality of the simulated data with those returned from the associated functionality.
	if !cmp.Equal(promotion, promotionAux) {
		bodyBytesAux, err = json.Marshal(promotionAux)

		if err != nil {
			t.Fatalf("Failed to obtain the JSON encoding of the returned promotion %+v: %s", promotionAux, err.Error())
		}

		t.Errorf("Test failed, the expected promotion returned: %s, got: %s", string(bodyBytes), string(bodyBytesAux))
		return
	}

	t.Logf("Test successful, the created promotion: %s", string(bodyBytes))
}
