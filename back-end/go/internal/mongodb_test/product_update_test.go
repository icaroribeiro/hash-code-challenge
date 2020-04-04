package mongodb_test

import (
    "encoding/json"
    "fmt"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/utils"
    "testing"
)

func TestUpdateProduct(t *testing.T) {
    var product models.Product
    var body string
    var err error
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

    product = models.Product{
        ID:           product.ID,
        PriceInCents: utils.GenerateRandomInteger(1, 1000),
        Title:        utils.GenerateRandomString(10),
        Description:  utils.GenerateRandomString(10),
    }

    body = fmt.Sprintf(`{"price_in_cents":%d,"title":"%s","description":"%s"}`,
        product.PriceInCents, product.Title, product.Description)

    t.Logf("Update product: %s", body)

    nMatchedDocs, nModifiedDocs, err = datastore.UpdateProduct(product.ID.Hex(), product)

    if err != nil {
        t.Fatalf("Failed to update the product with the id %s with %s: %s", product.ID.Hex(), body, err.Error())
    }

    if nMatchedDocs == 0 {
        t.Errorf("Test failed, the product with the id %s wasn't found", product.ID.Hex())
        return
    }

    if nModifiedDocs == 0 {
        t.Errorf("Test failed, the data sent are already registered")
    }

    if nModifiedDocs != 1 {
        t.Errorf("Test failed, the expected number of products updated: %d, got: %d", 1, nModifiedDocs)
        return
    }

    bodyBytes, err = json.Marshal(product)

    if err != nil {
        t.Fatalf("Failed to obtain the JSON encoding of the product %+v: %s", product, err.Error())
    }

    t.Logf("Test successful, the updated product: %s", string(bodyBytes))
}
