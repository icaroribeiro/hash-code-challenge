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

func TestCreateProduct(t *testing.T) {
    var product models.Product
    var body string
    var productEntity entities.Product
    var request services.CreateProductRequest
    var response *entities.Product
    var err error
    var errStatus *status.Status
    var bodyBytes []byte
    var bodyBytesAux []byte

    product = models.Product{
        PriceInCents: utils.GenerateRandomInteger(1, 1000),
        Title:        utils.GenerateRandomString(10),
        Description:  utils.GenerateRandomString(10),
    }

    body = fmt.Sprintf(`{"price_in_cents":%d,"title":"%s","description":"%s"}`,
        product.PriceInCents, product.Title, product.Description)

    t.Logf("Product: %s", body)

    productEntity = entities.Product{
        PriceInCents: int32(product.PriceInCents),
        Title:        product.Title,
        Description:  product.Description,
    }

    request = services.CreateProductRequest{
        Product: &productEntity,
    }

    response, err = productServiceClient.CreateProduct(ctx, &request)

    errStatus = status.Convert(err)

    if errStatus != nil {
        t.Errorf("Test failed, response: code=%d and body={\"error\":\"%s\",\"code\":%d,\"message\":\"%s\"}",
            errStatus.Code(), errStatus.Message(), errStatus.Code(), errStatus.Message())
    }

    productEntity.Id = response.Id

    bodyBytes, err = json.Marshal(productEntity)

    if err != nil {
        t.Fatalf("Failed to obtain the JSON encoding of the product %+v: %s", productEntity, err.Error())
    }

    // Evaluate the equality of the simulated data with those returned from the associated functionality.
    if !proto.Equal(&productEntity, response) {
        bodyBytesAux, err = json.Marshal(response)

        if err != nil {
            t.Fatalf("Failed to obtain the JSON encoding of the returned product %+v: %s", response, err.Error())
        }

        t.Errorf("Test failed, the expected product returned: %s, got: %s", string(bodyBytes), string(bodyBytesAux))
        return
    }

    t.Logf("Test successful, response: code=%d and body=%s", codes.OK, string(bodyBytes))
}
