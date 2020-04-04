package mongodb_test

import (
    "encoding/json"
    "fmt"
    "github.com/google/go-cmp/cmp"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/utils"
    "testing"
)

func TestGetAllProducts(t *testing.T) {
    var product models.Product
    var body string
    var err error
    var bodyBytes []byte
    var products []models.Product
    var isFound bool
    var productAux models.Product

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

    products, err = datastore.GetAllProducts()

    if err != nil {
        t.Fatalf("Failed to get the list of all products: %s", err.Error())
    }

    isFound = false

    for _, productAux = range products {
        // Evaluate the equality of the simulated data with those returned from the associated functionality.
        if cmp.Equal(product, productAux) {
            isFound = true
            break
        }
    }

    if !isFound {
        t.Errorf("Test failed, the product with the id %s wasn't found in the list of all products: %s",
            product.ID.Hex(), string(bodyBytes))
        return
    }

    t.Logf("Test succeeded, the product found in the list of all products: %s", string(bodyBytes))
}

func TestGetProduct(t *testing.T) {
    var product models.Product
    var body string
    var err error
    var bodyBytes []byte
    var productAux models.Product
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

    productAux, err = datastore.GetProduct(product.ID.Hex())

    if err != nil {
        t.Fatalf("Failed to get the product with the id %s: %s", product.ID.Hex(), err.Error())
    }

    // Evaluate the equality of the simulated data with those returned from the associated functionality.
    if !(cmp.Equal(product, productAux)) {
        bodyBytesAux, err = json.Marshal(productAux)

        if err != nil {
            t.Fatalf("Failed to obtain the JSON encoding of the returned product %+v: %s", productAux, err.Error())
        }

        t.Errorf("Test failed, the expected product returned: %s, got: %s", string(bodyBytes), string(bodyBytesAux))
        return
    }

    t.Logf("Test successful, the returned product: %s", string(bodyBytes))
}
