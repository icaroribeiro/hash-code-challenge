package impl

import (
    "fmt"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/entities"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/utils"
    context "golang.org/x/net/context"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func (p *ProductServiceServer) UpdateProduct(ctx context.Context,
    request *services.UpdateProductRequest) (*entities.Product, error) {
    var product models.Product
    var body string
    var nMatchedDocs int64
    var nModifiedDocs int64
    var err error
    var response *entities.Product

    if request.Id == "" {
        return nil, status.Error(codes.InvalidArgument,
            "The id is required and must be set to a non-empty value in the request URL")
    }

    if request.Product.PriceInCents <= 0 {
        return nil, status.Error(codes.InvalidArgument,
            "The price_in_cents field is required and must be set to a value greater than 0")
    }

    if request.Product.Title == "" {
        return nil, status.Error(codes.InvalidArgument,
            "The title field is required and must be set to a non-empty value")
    }

    if request.Product.Description == "" {
        return nil, status.Error(codes.InvalidArgument,
            "The description field is required and must be set to a non-empty value")
    }

    product = models.Product{
        PriceInCents: int(request.Product.PriceInCents),
        Title:        request.Product.Title,
        Description:  request.Product.Description,
    }

    body = fmt.Sprintf(`{
            "price_in_cents":%d,
            "title":"%s",
            "description":"%s"
        }`, product.PriceInCents, product.Title, product.Description)

    body = utils.RemoveEscapeSequencesFromString(body, "\t", "\n")

    nMatchedDocs, nModifiedDocs, err = p.ServiceServer.Datastore.UpdateProduct(request.Id, product)

    if err != nil {
        return nil, status.Error(codes.Internal,
            fmt.Sprintf("Failed to update the product with the id %s with %s: %s", request.Id, body, err.Error()))
    }

    if nMatchedDocs == 0 {
        return nil, status.Error(codes.NotFound,
            fmt.Sprintf("Failed to update the product with the id %s with %s: the product wasn't found", request.Id, body))
    }

    if nModifiedDocs == 0 {
        return nil, status.Error(codes.AlreadyExists,
            fmt.Sprintf("Failed to update the product with the id %s with %s: the data sent are already registered",
                request.Id, body))
    }

    if nModifiedDocs != 1 {
        return nil, status.Error(codes.Internal,
            fmt.Sprintf("Failed to update the product with the id %s with %s: the expected number of "+
                "products updated: %d, got: %d", request.Id, body, 1, nModifiedDocs))
    }

    response = &entities.Product{
        Id:           request.Id,
        PriceInCents: int32(product.PriceInCents),
        Title:        product.Title,
        Description:  product.Description,
    }

    return response, nil
}
