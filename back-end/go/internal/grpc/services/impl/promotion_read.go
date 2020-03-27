package impl

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/entities"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
	context "golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *PromotionServiceServer) GetAllPromotions(ctx context.Context, 
	e *empty.Empty) (*services.GetAllPromotionsResponse, error) {
		var promotions []models.Promotion
		var err error
		var promotion models.Promotion
		var promotionEntity *entities.Promotion
		var productId string
		var response *services.GetAllPromotionsResponse

		promotions, err = p.ServiceServer.Datastore.GetAllPromotions()

		if err != nil {
			return nil, status.Error(codes.Unknown,
				fmt.Sprintf("Failed to get the list of all promotions: %s", err.Error()))
		}

		response = &services.GetAllPromotionsResponse{}

		for _, promotion = range promotions {
			promotionEntity = &entities.Promotion{
				Id:             promotion.ID.Hex(),
				Code:           promotion.Code,
				Title:          promotion.Title,
				Description:    promotion.Description,
				MaxDiscountPct: float32(promotion.MaxDiscountPct),
			}

			// In case there is no product associated with the promotion,
			// don't display the products field as an empty array.
			if len(promotion.Products) > 0 {
				for _, productId = range promotion.Products {
					promotionEntity.Products = append(promotionEntity.Products, productId)
				}
			}

			response.Promotions = append(response.Promotions, promotionEntity)
		}

		return response, nil
}

func (p *PromotionServiceServer) GetPromotion(ctx context.Context, 
	request *services.GetPromotionRequest) (*entities.Promotion, error) {
		var promotion models.Promotion
		var err error
		var response *entities.Promotion
		var productId string

		promotion, err = p.ServiceServer.Datastore.GetPromotion(request.Id)

		if err != nil {
			return nil, status.Error(codes.Unknown,
				fmt.Sprintf("Failed to get the promotion with the id %s: %s", request.Id, err.Error()))
		}

		if promotion.ID.IsZero() {
			return nil, status.Error(codes.NotFound,
				fmt.Sprintf("Failed to get the promotion with the id %s: the id wasn't found", request.Id))
		}

		response = &entities.Promotion{
			Id:             promotion.ID.Hex(),
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
