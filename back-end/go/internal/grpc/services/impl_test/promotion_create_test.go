package impl_test

import (
    "encoding/json"
    "fmt"
    "github.com/golang/protobuf/proto"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/entities"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/utils"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "testing"
)

func TestCreatePromotion(t *testing.T) {
    var product models.Product
    var body string
    var bodyBytes []byte
    var promotion models.Promotion
    var promotionEntity entities.Promotion
    var request services.CreatePromotionRequest
    var response *entities.Promotion
    var err error
    var errStatus *status.Status
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

    t.Logf("Promotion: %s", body)

    promotionEntity = entities.Promotion{
        Code:           promotion.Code,
        Title:          promotion.Title,
        Description:    promotion.Description,
        MaxDiscountPct: float32(promotion.MaxDiscountPct),
        Products:       []string{product.ID.Hex()},
    }

    request = services.CreatePromotionRequest{
        Promotion: &promotionEntity,
    }

    response, err = promotionServiceClient.CreatePromotion(ctx, &request)

    errStatus = status.Convert(err)

    if errStatus != nil {
        t.Errorf(`Test failed, response: code=%d and body={"error":"%s","code":%d,"message":"%s"}`,
            errStatus.Code(), errStatus.Message(), errStatus.Code(), errStatus.Message())
        return
    }

    promotionEntity.Id = response.Id

    bodyBytes, err = json.Marshal(promotionEntity)

    if err != nil {
        t.Fatalf("Failed to obtain the JSON encoding of the promotion %+v: %s", promotionEntity, err.Error())
    }

    // Evaluate the equality of the simulated data with those returned from the associated functionality.
    if !proto.Equal(&promotionEntity, response) {
        bodyBytesAux, err = json.Marshal(response)

        if err != nil {
            t.Fatalf("Failed to obtain the JSON encoding of the returned promotion %+v: %s", response, err.Error())
        }

        t.Errorf("Test failed, the expected promotion returned: %s, got: %s", string(bodyBytes), string(bodyBytesAux))
        return
    }

    t.Logf("Test successful, response: code=%d and body=%s", codes.OK, string(bodyBytes))
}
