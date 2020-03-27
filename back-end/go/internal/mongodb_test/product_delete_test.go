package mongodb_test

import (
	"encoding/json"
	"fmt"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/utils"
	"testing"
)

func TestDeleteProduct(t *testing.T) {
	var product models.Product
	var err error
	var body string
	var bodyBytes []byte
	var nDeletedDocs int64

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

	bodyBytes, err = json.Marshal(product)

	if err != nil {
		t.Fatalf("Failed to obtain the JSON encoding of the product %+v: %s", product, err.Error())
	}

	t.Logf("Product: %s", string(bodyBytes))

	nDeletedDocs, err = datastore.DeleteProduct(product.ID.Hex())

	if err != nil {
		t.Fatalf("Failed to delete the product with the id %s: %s", product.ID.Hex(), err.Error())
	}

	if nDeletedDocs != 1 {
		t.Errorf("Test failed, the expected number of products deleted: %d, got: %d", 1, nDeletedDocs)
		return
	}

	t.Logf("Test successful, the deleted product: %s", string(bodyBytes))
}
