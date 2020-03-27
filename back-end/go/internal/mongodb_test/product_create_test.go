package mongodb_test

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/utils"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	var product models.Product
	var body string
	var err error
	var productAux models.Product
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

	body = utils.RemoveEscapeSequences(body, "\t", "\n")

	t.Logf("Product: %s", body)

	productAux, err = datastore.CreateProduct(product)

	if err != nil {
		t.Fatalf("Failed to create a new product with %s: %s", body, err.Error())
	}

	product.ID = productAux.ID

	bodyBytes, err = json.Marshal(product)

	if err != nil {
		t.Fatalf("Failed to obtain the JSON encoding of the product %+v: %s", product, err.Error())
	}

	if !cmp.Equal(product, productAux) {
		bodyBytesAux, err = json.Marshal(productAux)

		if err != nil {
			t.Fatalf("Failed to obtain the JSON encoding of the returned product %+v: %s", productAux, err.Error())
		}

		t.Errorf("Test failed, the expected product returned: %s, got: %s", string(bodyBytes), string(bodyBytesAux))
		return
	}

	t.Logf("Test successful, the created product: %s", string(bodyBytes))
}
