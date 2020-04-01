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

func (p *PromotionServiceServer) DeletePromotion(ctx context.Context, 
	request *services.DeletePromotionRequest) (*entities.Promotion, error) {
		var promotion models.Promotion
		var nDeletedDocs int64
		var err error
		var response *entities.Promotion
		var productId string

		if request.Id == "" {
			return nil, status.Error(codes.InvalidArgument,
				"The id is required and must be set to a non-empty value in the request URL")
		}

		promotion, err = p.ServiceServer.Datastore.GetPromotion(request.Id)

		if err != nil {
			return nil, status.Error(codes.Internal,
				fmt.Sprintf("Failed to get the promotion with the id %s: %s", request.Id, err.Error()))
		}

		nDeletedDocs, err = p.ServiceServer.Datastore.DeletePromotion(request.Id)

		if err != nil {
			return nil, status.Error(codes.Internal,
				fmt.Sprintf("Failed to delete the promotion with the id %s: %s", request.Id, err.Error()))
		}

		if nDeletedDocs == 0 {
			return nil, status.Error(codes.NotFound,
				fmt.Sprintf("Failed to delete the promotion with the id %s: the promotion wasn't found", request.Id))
		}

		if nDeletedDocs > 1 {
			return nil, status.Error(codes.Internal,
				fmt.Sprintf("Failed to delete the promotion with the id %s: the expected number of "+
					"promotions deleted: %d, got: %d", request.Id, 1, nDeletedDocs))
		}

		response = &entities.Promotion{
			Id:             promotion.ID.Hex(),
			Code:           promotion.Code,
			Title:          promotion.Title,
			Description:    promotion.Description,
			MaxDiscountPct: float32(promotion.MaxDiscountPct),
		}

		// In case of there is no product associated with the promotion,
		// don't display the products field as an empty array.
		if len(promotion.Products) > 0 {
			for _, productId = range promotion.Products {
				response.Products = append(response.Products, productId)
			}
		}

		return response, nil
}
