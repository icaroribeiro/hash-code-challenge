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

func (p *PromotionServiceServer) UpdatePromotion(ctx context.Context, 
	request *services.UpdatePromotionRequest) (*entities.Promotion, error) {
		var promotion models.Promotion
		var body string
		var productsMap map[string]bool
		var i int
		var productId string
		var err error
		var product models.Product
		var nMatchedDocs int64
		var nModifiedDocs int64
		var response *entities.Promotion

		if request.Id == "" {
			return nil, status.Error(codes.InvalidArgument,
				"The id is required and must be set to a non-empty value in the request URL")
		}

		if request.Promotion.Code == "" {
			return nil, status.Error(codes.InvalidArgument,
				"The code field is required and must be set to a non-empty value")
		}

		if request.Promotion.Title == "" {
			return nil, status.Error(codes.InvalidArgument,
				"The title field is required and must be set to a non-empty value")
		}

		if request.Promotion.Description == "" {
			return nil, status.Error(codes.InvalidArgument,
				"The description field is required and must be set to a non-empty value")
		}

		if request.Promotion.MaxDiscountPct <= 0 {
			return nil, status.Error(codes.InvalidArgument,
				"The max_discount_pct field is required and must be set to a value greater than 0")
		}

		promotion = models.Promotion{
			Code:           request.Promotion.Code,
			Title:          request.Promotion.Title,
			Description:    request.Promotion.Description,
			MaxDiscountPct: float32(request.Promotion.MaxDiscountPct),
		}

		body = fmt.Sprintf(`{
			"code":"%s",
			"title":"%s",
			"description":"%s",
			"max_discount_pct":%f
		`, promotion.Code, promotion.Title, promotion.Description, promotion.MaxDiscountPct)

		// Verify if all the ids of the products associated with the promotion are valid.
		// Additionally, checks if there are no duplicated ids of the products.
		productsMap = make(map[string]bool)

		if len(request.Promotion.Products) > 0 {
			for i, productId = range request.Promotion.Products {
				if productId == "" {
					return nil, status.Error(codes.InvalidArgument,
						"There is an empty value in the list of all products")
				}

				product, err = p.ServiceServer.Datastore.GetProduct(productId)

				if err != nil {
					return nil, status.Error(codes.Internal,
						fmt.Sprintf("Failed to add the product with the id %s: %s", productId, err.Error()))
				}

				if product.ID.IsZero() {
					return nil, status.Error(codes.NotFound,
						fmt.Sprintf("Failed to add the product with the id %s: the product wasn't found", productId))
				}

				if !(productsMap[product.ID.Hex()]) {
					productsMap[product.ID.Hex()] = true
				} else {
					return nil, status.Error(codes.Internal,
						fmt.Sprintf("Failed to add the product with the id %s: the id is duplicated", product.ID.Hex()))
				}

				promotion.Products = append(promotion.Products, product.ID.Hex())

				if i == 0 {
					body += fmt.Sprintf(`,"products":["%s"`, product.ID.Hex())
				} else {
					body += fmt.Sprintf(`,"%s"`, product.ID.Hex())
				}
			}

			if len(promotion.Products) > 0 {
				body += `]}`
			} else {
				body += `}`
			}

		} else {
			body += `}`
		}

		body = utils.RemoveEscapeSequencesFromString(body, "\t", "\n")

		nMatchedDocs, nModifiedDocs, err = p.ServiceServer.Datastore.UpdatePromotion(request.Id, promotion)

		if err != nil {
			return nil, status.Error(codes.Internal,
				fmt.Sprintf("Failed to update the promotion with the id %s with %s: %s", request.Id, body, err.Error()))
		}

		if nMatchedDocs == 0 {
			return nil, status.Error(codes.NotFound,
				fmt.Sprintf("Failed to update the promotion with the id %s with %s: the id wasn't found", request.Id, body))
		}

		if nModifiedDocs == 0 {
			return nil, status.Error(codes.Internal,
				fmt.Sprintf("Failed to update the promotion with the id %s with %s: the data sent is already registered",
					request.Id, body))
		}

		if nModifiedDocs != 1 {
			return nil, status.Error(codes.NotFound,
				fmt.Sprintf("Failed to update the promotion with the id %s with %s: the expected number of "+
					"promotions updated: %d, got: %d", request.Id, body, 1, nModifiedDocs))
		}

		response = &entities.Promotion{
			Id:             request.Id,
			Code:           promotion.Code,
			Title:          promotion.Title,
			Description:    promotion.Description,
			MaxDiscountPct: float32(promotion.MaxDiscountPct),
		}

		// In case there is no product associated with the promotion,
		// don't display the products field as an empty array.
		if len(promotion.Products) > 0 {
			for _, productId = range promotion.Products {
				response.Products = append(response.Products, productId)
			}
		}

		return response, nil
}
