package mongodb_test

import (
    "encoding/json"
    "fmt"
    "github.com/google/go-cmp/cmp"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/utils"
    "testing"
)

func TestGetAllPromotions(t *testing.T) {
    var product models.Product
    var body string
    var err error
    var bodyBytes []byte
    var promotion models.Promotion
    var promotions []models.Promotion
    var isFound bool
    var promotionAux models.Promotion

    product = models.Product{
        PriceInCents: utils.GenerateRandomInteger(1, 1000),
        Title:        utils.GenerateRandomString(10),
        Description:  utils.GenerateRandomString(10),
    }

    body = fmt.Sprintf(`{"price_in_cents":%d,"title":"%s","description":"%s"}`,
        product.PriceInCents, product.Title, product.Description)

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

    body = fmt.Sprintf(`{"code":"%s","title":"%s","description":"%s","max_discount_pct":%f,"products":["%s"]}`,
        promotion.Code, promotion.Title, promotion.Description, promotion.MaxDiscountPct, promotion.Products[0])

    promotion, err = datastore.CreatePromotion(promotion)

    if err != nil {
        t.Fatalf("Failed to create a new promotion with %s: %s", body, err.Error())
    }

    bodyBytes, err = json.Marshal(promotion)

    if err != nil {
        t.Fatalf("Failed to obtain the JSON encoding of the promotion %+v: %s", promotion, err.Error())
    }

    t.Logf("Promotion: %s", string(bodyBytes))

    promotions, err = datastore.GetAllPromotions()

    if err != nil {
        t.Fatalf("Failed to get the list of all promotions: %s", err.Error())
    }

    isFound = false

    for _, promotionAux = range promotions {
        // Evaluate the equality of the simulated data with those returned from the associated functionality.
        if cmp.Equal(promotion, promotionAux) {
            isFound = true
            break
        }
    }

    if !isFound {
        t.Errorf("Test failed, the promotion not found in the list of all promotions: %s", string(bodyBytes))
        return
    }

    t.Logf("Test successful, the promotion found in the list of all promotions: %s", string(bodyBytes))
}

func TestGetPromotion(t *testing.T) {
    var product models.Product
    var body string
    var err error
    var bodyBytes []byte
    var promotion models.Promotion
    var promotionAux models.Promotion
    var bodyBytesAux []byte

    product = models.Product{
        PriceInCents: utils.GenerateRandomInteger(1, 1000),
        Title:        utils.GenerateRandomString(10),
        Description:  utils.GenerateRandomString(10),
    }

    body = fmt.Sprintf(`{"price_in_cents":%d,"title":"%s","description":"%s"}`,
        product.PriceInCents, product.Title, product.Description)

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

    body = fmt.Sprintf(`{"code":"%s","title":"%s","description":"%s","max_discount_pct":%f,"products":["%s"]}`,
        promotion.Code, promotion.Title, promotion.Description, promotion.MaxDiscountPct, promotion.Products[0])

    promotion, err = datastore.CreatePromotion(promotion)

    if err != nil {
        t.Fatalf("Failed to create a new promotion with %s: %s", body, err.Error())
    }

    bodyBytes, err = json.Marshal(promotion)

    if err != nil {
        t.Fatalf("Failed to obtain the JSON encoding of the promotion %+v: %s", promotion, err.Error())
    }

    t.Logf("Promotion: %s", string(bodyBytes))

    promotionAux, err = datastore.GetPromotion(promotion.ID.Hex())

    if err != nil {
        t.Fatalf("Failed to get the promotion with the id %s: %s", promotion.ID.Hex(), err.Error())
    }

    // Evaluate the equality of the simulated data with those returned from the associated functionality.
    if !(cmp.Equal(promotion, promotionAux)) {
        bodyBytesAux, err = json.Marshal(promotionAux)

        if err != nil {
            t.Fatalf("Failed to obtain the JSON encoding of the returned promotion %+v: %s", promotionAux, err.Error())
        }

        t.Errorf("Test failed, the expected promotion returned: %s, got: %s", string(bodyBytes), string(bodyBytesAux))
        return
    }

    t.Logf("Test successful, the returned promotion: %s", string(bodyBytes))
}
