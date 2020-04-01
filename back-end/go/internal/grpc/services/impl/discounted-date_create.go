package impl

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/entities"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/utils"
	context "golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (d *DiscountedDateServiceServer) CreateDiscountedDate(ctx context.Context, 
	request *services.CreateDiscountedDateRequest) (*entities.DiscountedDate, error) {
		var discountedDate models.DiscountedDate
		var body string
		var err error
		var response *entities.DiscountedDate

		if request.DiscountedDate.Title == "" {
			return nil, status.Error(codes.InvalidArgument, 
				"The title field is required and must be set to a non-empty value")
		}

		if request.DiscountedDate.Description == "" {
			return nil, status.Error(codes.InvalidArgument, 
				"The description field is required and must be set to a non-empty value")
		}

		if float32(request.DiscountedDate.DiscountPct) <= 0 {
			return nil, status.Error(codes.InvalidArgument, 
				"The discount_pct field is required and must be set to a value greater than 0")
		}

		if request.DiscountedDate.Date == nil {
			return nil, status.Error(codes.InvalidArgument,
				"The date field is required along with the year, month and day fields")
		}

		if request.DiscountedDate.Date.Year == nil {
			return nil, status.Error(codes.InvalidArgument,
				"The year field of the date field is required and must be set to a value in the range from 0 to 9999")
		} else {
			if request.DiscountedDate.Date.Year.Value < 0 || request.DiscountedDate.Date.Year.Value > 9999 {
				return nil, status.Error(codes.InvalidArgument,
					"The year field of the date field is required and must be set to a value in the range from 0 to 9999")
			}
		}

		if request.DiscountedDate.Date.Month == nil {
			return nil, status.Error(codes.InvalidArgument,
				"The month field of the date field is required and must be set to a value in the range from 0 to 12")
		} else {
			if request.DiscountedDate.Date.Month.Value < 0 || request.DiscountedDate.Date.Month.Value > 12 {
				return nil, status.Error(codes.InvalidArgument,
					"The month field of the date field is required and must be set to a value in the range from 0 to 12")
			}
		}

		if request.DiscountedDate.Date.Day == nil {
			return nil, status.Error(codes.InvalidArgument,
				"The day field of the date field is required and must be set to a value in the range from 0 to 31")
		} else {
			if request.DiscountedDate.Date.Day.Value < 0 || request.DiscountedDate.Date.Day.Value > 31 {
				return nil, status.Error(codes.InvalidArgument,
					"The day field of the date field is required and must be set to a value in the range from 0 to 31")
			}
		}

		discountedDate = models.DiscountedDate{
			Title:       request.DiscountedDate.Title,
			Description: request.DiscountedDate.Description,
			DiscountPct: float32(request.DiscountedDate.DiscountPct),
			Date: models.Date{
				Year:  int(request.DiscountedDate.Date.Year.Value),
				Month: int(request.DiscountedDate.Date.Month.Value),
				Day:   int(request.DiscountedDate.Date.Day.Value),
			},
		}

		body = fmt.Sprintf(`{
			"title":"%s",
			"description":"%s",
			"discount_pct":%f,
			"date":{
				"year":%d,
				"month":%d,
				"day":%d
			}
		}`, discountedDate.Title, discountedDate.Description, discountedDate.DiscountPct,
			discountedDate.Date.Year, discountedDate.Date.Month, discountedDate.Date.Day)

		body = utils.RemoveEscapeSequencesFromString(body, "\t", "\n")

		discountedDate, err = d.ServiceServer.Datastore.CreateDiscountedDate(discountedDate)

		if err != nil {
			return nil, status.Error(codes.Internal, 
				fmt.Sprintf("Failed to create a new discounted date with %s: %s", body, err.Error()))
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
