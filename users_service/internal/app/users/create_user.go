package users

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	desc "heisenbug/users/pkg/api/users"
)

func (i *Implementation) CreateUser(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	if err := validateCreateUser(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	userID, err := i.repo.Add(ctx, req.Fio, req.Phone, req.Age)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.CreateUserResponse{UserId: userID}, nil
}

func validateCreateUser(req *desc.CreateUserRequest) error {
	err := validation.ValidateStruct(req,
		validation.Field(&req.Fio, validation.Required),
		validation.Field(&req.Phone, validation.Required),
		validation.Field(&req.Age, validation.Required),
	)
	return err
}
