package app

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"heisenbug/identification/internal/pb/users"
	desc "heisenbug/identification/pkg/api/identification"
)

func (i *Implementation) Upgrade(ctx context.Context, req *desc.UpgradeRequest) (*desc.UpgradeResponse, error) {
	if err := validateUpgrade(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := i.usersClient.GetInfo(ctx, &users.GetInfoRequest{
		UserId: req.UserId,
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if info.Level == users.Level_FULL {
		return nil, status.Error(codes.Unavailable, "User identification level already FULL")
	}

	resp, err := i.externalClient.Identification(ctx, info.Phone)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if resp.Status != 0 {
		return &desc.UpgradeResponse{
			IdentificationId: "empty",
		}, nil
	}

	_, err = i.usersClient.UpdateUserLevel(ctx, &users.UpdateUserLevelRequest{
		UserId: req.UserId,
		Level:  users.Level_FULL,
	})
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	return &desc.UpgradeResponse{
		IdentificationId: resp.Id,
	}, nil
}

func validateUpgrade(req *desc.UpgradeRequest) error {
	err := validation.ValidateStruct(req,
		validation.Field(&req.UserId, validation.Required),
	)
	return err
}
