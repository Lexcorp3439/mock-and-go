package users

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	desc "heisenbug/users/pkg/api/users"
)

func (i *Implementation) UpdateUserLevel(ctx context.Context, req *desc.UpdateUserLevelRequest) (*desc.UpdateUserLevelResponse, error) {
	if err := validateUpdateUserLevel(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := i.repo.Upgrade(ctx, req.UserId, int32(req.Level))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if user.Level != int32(req.Level) {
		return nil, status.Error(codes.Internal, "failed to upgrade level")
	}

	return &desc.UpdateUserLevelResponse{}, nil
}

func validateUpdateUserLevel(req *desc.UpdateUserLevelRequest) error {
	err := validation.ValidateStruct(req,
		validation.Field(&req.UserId, validation.Required),
		validation.Field(&req.Level, validation.Required),
	)
	return err
}
