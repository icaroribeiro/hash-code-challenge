package mongodb_test

import (
    "encoding/json"
    "fmt"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/utils"
    "testing"
)

func TestDeletePromotion(t *testing.T) {
    var product models.Product
    var body string
    var err error
    var bodyBytes []byte
    var promotion models.Promotion
    var nDeletedDocs int64

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

    nDeletedDocs, err = datastore.DeletePromotion(promotion.ID.Hex())

    if err != nil {
        t.Fatalf("Failed to delete the promotion with the id %s: %s", promotion.ID.Hex(), err.Error())
    }

    if nDeletedDocs == 0 {
        t.Errorf("Test failed, the promotion with the id %s wasn't found", promotion.ID.Hex())
        return
    }

    if nDeletedDocs != 1 {
        t.Errorf("Test failed, the expected number of promotions deleted: %d, got: %d", 1, nDeletedDocs)
        return
    }

    t.Logf("Test successful, the deleted promotion: %s", string(bodyBytes))
}
