package impl

import (
	"fmt"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/entities"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/services"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/models"
	"github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/utils"
	context "golang.org/x/net/context"
	date "google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *UserServiceServer) CreateUser(ctx context.Context, 
	request *services.CreateUserRequest) (*entities.User, error) {
		var user models.User
		var body string
		var err error
		var response *entities.User

		if request.User.FirstName == "" {
			return nil, status.Error(codes.InvalidArgument,
				"The first_name field is required and must be set to a non-empty value")
		}

		if request.User.LastName == "" {
			return nil, status.Error(codes.InvalidArgument,
				"The last_name field is required and must be set to a non-empty value")
		}

		if request.User.DateOfBirth == nil {
			return nil, status.Error(codes.InvalidArgument,
				"The date_of_birth field is required along with the year, month and day fields")
		}

		if request.User.DateOfBirth.Year < 1 || request.User.DateOfBirth.Year > 9999 {
			return nil, status.Error(codes.InvalidArgument,
				"The year field of date_of_birth field is required and must be set to a value in the range from 1 to 9999")
		}

		if request.User.DateOfBirth.Month < 1 || request.User.DateOfBirth.Month > 12 {
			return nil, status.Error(codes.InvalidArgument,
				"The month field of date_of_birth field is required and must be set to a value in the range from 1 to 12")
		}

		if request.User.DateOfBirth.Day < 1 || request.User.DateOfBirth.Day > 31 {
			return nil, status.Error(codes.InvalidArgument,
				"The day field of date_of_birth field is required and must be set to a value in the range from 1 to 31")
		}

		user = models.User{
			FirstName: request.User.FirstName,
			LastName:  request.User.LastName,
			DateOfBirth: models.Date{
				Year:  int(request.User.DateOfBirth.Year),
				Month: int(request.User.DateOfBirth.Month),
				Day:   int(request.User.DateOfBirth.Day),
			},
		}

		body = fmt.Sprintf(`{
			"first_name":"%s",
			"last_name":"%s",
			"date_of_birth":{
				"year":%d,
				"month":%d,
				"day":%d
			}
		}`, user.FirstName, user.LastName, user.DateOfBirth.Year, user.DateOfBirth.Month, user.DateOfBirth.Day)

		body = utils.RemoveEscapeSequencesFromString(body, "\t", "\n")

		user, err = u.ServiceServer.Datastore.CreateUser(user)

		if err != nil {
			return nil, status.Error(codes.Unknown,
				fmt.Sprintf("Failed to create a new user with %s: %s", body, err.Error()))
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
