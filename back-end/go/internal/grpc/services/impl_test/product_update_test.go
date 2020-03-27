package impl_test

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/entities"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/utils"
	"google.golang.org/grpc/status"
	"testing"
)

func TestUpdateProduct(t *testing.T) {
	var product models.Product
	var body string
	var err error
	var bodyBytes []byte
	var productEntity entities.Product
	var request services.UpdateProductRequest
	var response *entities.Product
	var errStatus *status.Status
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
		ID:			  product.ID,
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

	t.Logf("Update product: %s", body)

	productEntity = entities.Product{
		Id:           product.ID.Hex(),
		PriceInCents: int32(product.PriceInCents),
		Title:        product.Title,
		Description:  product.Description,
	}

	request = services.UpdateProductRequest{
		Id:      productEntity.Id,
		Product: &productEntity,
	}

	response, err = productServiceClient.UpdateProduct(ctx, &request)

	errStatus = status.Convert(err)

	if errStatus != nil {
		t.Errorf("Test failed, response: code=%d and body={\"error\":\"%s\",\"code\":%d,\"message\":\"%s\"}",
			errStatus.Code(), errStatus.Message(), errStatus.Code(), errStatus.Message())
	}

	bodyBytes, err = json.Marshal(productEntity)

	if err != nil {
		t.Fatalf("Failed to obtain the JSON encoding of the product %+v: %s", productEntity, err.Error())
	}

	if !proto.Equal(&productEntity, response) {
		bodyBytesAux, err = json.Marshal(response)

		if err != nil {
			t.Fatalf("Failed to obtain the JSON encoding of the returned product %+v: %s", response, err.Error())
		}

		t.Errorf("Test failed, the expected product returned: %s, got: %s", string(bodyBytes), string(bodyBytesAux))
		return
	}

	t.Logf("Test successful, response: code=%d and body=%s", 0, string(bodyBytes))
}
