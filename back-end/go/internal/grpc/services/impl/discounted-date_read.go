package impl

import (
    "fmt"
    "github.com/golang/protobuf/ptypes/empty"
    "github.com/golang/protobuf/ptypes/wrappers"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/entities"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services"
    "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
    context "golang.org/x/net/context"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func (d *DiscountedDateServiceServer) GetAllDiscountedDates(ctx context.Context,
    e *empty.Empty) (*services.GetAllDiscountedDatesResponse, error) {
    var discountedDates []models.DiscountedDate
    var err error
    var discountedDate models.DiscountedDate
    var discountedDateEntity *entities.DiscountedDate
    var response *services.GetAllDiscountedDatesResponse

    discountedDates, err = d.ServiceServer.Datastore.GetAllDiscountedDates()

    if err != nil {
        return nil, status.Error(codes.Internal,
            fmt.Sprintf("Failed to get the list of all discounted dates: %s", err.Error()))
    }

    response = &services.GetAllDiscountedDatesResponse{}

    for _, discountedDate = range discountedDates {
        discountedDateEntity = &entities.DiscountedDate{
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

        response.DiscountedDates = append(response.DiscountedDates, discountedDateEntity)
    }

    return response, nil
}

func (d *DiscountedDateServiceServer) GetDiscountedDate(ctx context.Context,
    request *services.GetDiscountedDateRequest) (*entities.DiscountedDate, error) {
    var discountedDate models.DiscountedDate
    var err error
    var response *entities.DiscountedDate

    if request.Id == "" {
        return nil, status.Error(codes.InvalidArgument,
            "The id is required and must be set to a non-empty value in the request URL")
    }

    discountedDate, err = d.ServiceServer.Datastore.GetDiscountedDate(request.Id)

    if err != nil {
        return nil, status.Error(codes.Internal,
            fmt.Sprintf("Failed to get the discounted date with the id %s: %s", request.Id, err.Error()))
    }

    if discountedDate.ID.IsZero() {
        return nil, status.Error(codes.NotFound,
            fmt.Sprintf("Failed to get the discounted date with the id %s: the discounted date wasn't found", request.Id))
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
