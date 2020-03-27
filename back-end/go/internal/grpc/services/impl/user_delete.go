package impl

import (
	"fmt"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/entities"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
	context "golang.org/x/net/context"
	date "google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *UserServiceServer) DeleteUser(ctx context.Context, 
	request *services.DeleteUserRequest) (*entities.User, error) {
		var user models.User
		var nDeletedDocs int64
		var err error
		var response *entities.User

		user, err = u.ServiceServer.Datastore.GetUser(request.Id)

		if err != nil {
			return nil, status.Error(codes.Unknown,
				fmt.Sprintf("Failed to get the user with the id %s: %s", request.Id, err.Error()))
		}

		nDeletedDocs, err = u.ServiceServer.Datastore.DeleteUser(request.Id)

		if err != nil {
			return nil, status.Error(codes.Unknown,
				fmt.Sprintf("Failed to delete the user with the id %s: %s", request.Id, err.Error()))
		}

		if nDeletedDocs == 0 {
			return nil, status.Error(codes.NotFound,
				fmt.Sprintf("Failed to delete the user with the id %s: the id wasn't found", request.Id))
		}

		if nDeletedDocs > 1 {
			return nil, status.Error(codes.Internal,
				fmt.Sprintf("Failed to delete the user with the id %s: the expected number of users deleted: %d, got: %d",
					request.Id, 1, nDeletedDocs))
		}

		response = &entities.User{
			Id:        user.ID.Hex(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			DateOfBirth: &date.Date{
				Year:  int32(user.DateOfBirth.Year),
				Month: int32(user.DateOfBirth.Month),
				Day:   int32(user.DateOfBirth.Day),
			},
		}

		return response, nil
}
