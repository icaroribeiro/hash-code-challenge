package impl

import (
    "fmt"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/entities"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
    context "golang.org/x/net/context"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func (p *ProductServiceServer) DeleteProduct(ctx context.Context,
    request *services.DeleteProductRequest) (*entities.Product, error) {
    var product models.Product
    var nDeletedDocs int64
    var err error
    var response *entities.Product

    if request.Id == "" {
        return nil, status.Error(codes.InvalidArgument,
            "The id is required and must be set to a non-empty value in the request URL")
    }

    product, err = p.ServiceServer.Datastore.GetProduct(request.Id)

    if err != nil {
        return nil, status.Error(codes.Internal,
            fmt.Sprintf("Failed to get the product with the id %s: %s", request.Id, err.Error()))
    }

    nDeletedDocs, err = p.ServiceServer.Datastore.DeleteProduct(request.Id)

    if err != nil {
        return nil, status.Error(codes.Internal,
            fmt.Sprintf("Failed to delete the product with the id %s: %s", request.Id, err.Error()))
    }

    if nDeletedDocs == 0 {
        return nil, status.Error(codes.NotFound,
            fmt.Sprintf("Failed to delete the product with the id %s: the product wasn't found", request.Id))
    }

    if nDeletedDocs > 1 {
        return nil, status.Error(codes.Internal,
            fmt.Sprintf("Failed to delete the product with the id %s: the expected number of "+
                "products deleted: %d, got: %d", request.Id, 1, nDeletedDocs))
    }

    response = &entities.Product{
        Id:           product.ID.Hex(),
        PriceInCents: int32(product.PriceInCents),
        Title:        product.Title,
        Description:  product.Description,
    }

    return response, nil
}
