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

func (p *ProductServiceServer) CreateProduct(ctx context.Context, 
	request *services.CreateProductRequest) (*entities.Product, error) {
		var product models.Product
		var body string
		var err error
		var response *entities.Product

		if int(request.Product.PriceInCents) == 0 {
			return nil, status.Error(codes.InvalidArgument, 
				"The price_in_cents field is required and must be set to a non-zero value")
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

		product, err = p.ServiceServer.Datastore.CreateProduct(product)

		if err != nil {
			return nil, status.Error(codes.Unknown, 
				fmt.Sprintf("Failed to create a new product with %s: %s", body, err.Error()))
		}

		response = &entities.Product{
			Id:           product.ID.Hex(),
			PriceInCents: int32(product.PriceInCents),
			Title:        product.Title,
			Description:  product.Description,
		}

		return response, nil
}
