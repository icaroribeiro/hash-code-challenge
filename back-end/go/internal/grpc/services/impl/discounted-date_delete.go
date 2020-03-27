package impl

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/entities"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
	context "golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (d *DiscountedDateServiceServer) DeleteDiscountedDate(ctx context.Context, 
	request *services.DeleteDiscountedDateRequest) (*entities.DiscountedDate, error) {
		var discountedDate models.DiscountedDate
		var nDeletedDocs int64
		var err error
		var response *entities.DiscountedDate

		discountedDate, err = d.ServiceServer.Datastore.GetDiscountedDate(request.Id)

		if err != nil {
			return nil, status.Error(codes.Unknown,
				fmt.Sprintf("Failed to get the discounted date with the id %s: %s", request.Id, err.Error()))
		}

		nDeletedDocs, err = d.ServiceServer.Datastore.DeleteDiscountedDate(request.Id)

		if err != nil {
			return nil, status.Error(codes.Unknown,
				fmt.Sprintf("Failed to delete the discounted date with the id %s: %s", request.Id, err.Error()))
		}

		if nDeletedDocs == 0 {
			return nil, status.Error(codes.NotFound,
				fmt.Sprintf("Failed to delete the discounted date with the id %s: the id wasn't found", request.Id))
		}

		if nDeletedDocs > 1 {
			return nil, status.Error(codes.Internal,
				fmt.Sprintf("Failed to delete the discounted date with id %s: the expected number of "+
					"discounted dates deleted: %d, got: %d",
					request.Id, 1, nDeletedDocs))
		}

		response = &entities.DiscountedDate{
			Id:          discountedDate.ID.Hex(),
			Title:       discountedDate.Title,
			Description: discountedDate.Description,
			DiscountPct: float32(discountedDate.DiscountPct),
			Date: &entities.Date{
				Year: &wrappers.Int32Value{
					Value: int32(discountedDate.Date.Year),
				},
				Month: &wrappers.Int32Value{
					Value: int32(discountedDate.Date.Month),
				},
				Day: &wrappers.Int32Value{
					Value: int32(discountedDate.Date.Day),
				},
			},
		}

		return response, nil
}
