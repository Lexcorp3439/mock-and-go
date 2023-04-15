package users

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	desc "heisenbug/users/pkg/api/users"
)

func (i *Implementation) GetInfo(ctx context.Context, req *desc.GetInfoRequest) (*desc.GetInfoResponse, error) {
	if err := validateGetInfo(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := i.repo.Get(ctx, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.GetInfoResponse{
		UserId: user.ID,
		Fio:    user.FIO,
		Phone:  user.Phone,
		Age:    user.Age,
		Level:  desc.Level(user.Level),
	}, nil
}

func validateGetInfo(req *desc.GetInfoRequest) error {
	err := validation.ValidateStruct(req,
		validation.Field(&req.UserId, validation.Required),
	)
	return err
}
