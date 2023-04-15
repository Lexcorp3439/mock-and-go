package users

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	desc "heisenbug/users/pkg/api/users"
)

func (i *Implementation) RemoveUser(ctx context.Context, req *desc.RemoveUserRequest) (*desc.RemoveUserResponse, error) {
	if err := validateRemoveUser(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := i.repo.Delete(ctx, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.RemoveUserResponse{}, nil
}

func validateRemoveUser(req *desc.RemoveUserRequest) error {
	err := validation.ValidateStruct(req,
		validation.Field(&req.UserId, validation.Required),
	)
	return err
}
