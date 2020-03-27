package impl

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/entities"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
	context "golang.org/x/net/context"
	date "google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *UserServiceServer) GetAllUsers(ctx context.Context, 
	e *empty.Empty) (*services.GetAllUsersResponse, error) {
		var users []models.User
		var err error
		var user models.User
		var userEntity *entities.User
		var response *services.GetAllUsersResponse

		users, err = u.ServiceServer.Datastore.GetAllUsers()

		if err != nil {
			return nil, status.Error(codes.Unknown,
				fmt.Sprintf("Failed to get the list of all users: %s", err.Error()))
		}

		response = &services.GetAllUsersResponse{}

		for _, user = range users {
			userEntity = &entities.User{
				Id:        user.ID.Hex(),
				FirstName: user.FirstName,
				LastName:  user.LastName,
				DateOfBirth: &date.Date{
					Year:  int32(user.DateOfBirth.Year),
					Month: int32(user.DateOfBirth.Month),
					Day:   int32(user.DateOfBirth.Day),
				},
			}

			response.Users = append(response.Users, userEntity)
		}

		return response, nil
}

func (u *UserServiceServer) GetUser(ctx context.Context, 
	request *services.GetUserRequest) (*entities.User, error) {
		var user models.User
		var err error
		var response *entities.User

		user, err = u.ServiceServer.Datastore.GetUser(request.Id)

		if err != nil {
			return nil, status.Error(codes.Unknown,
				fmt.Sprintf("Failed to get the user with the id %s: %s", request.Id, err.Error()))
		}

		if user.ID.IsZero() {
			return nil, status.Error(codes.NotFound,
				fmt.Sprintf("Failed to get the user with the id %s: the id wasn't found", request.Id))
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
