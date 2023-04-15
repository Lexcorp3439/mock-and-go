package app

import (
	"context"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"heisenbug/identification/internal/pb/users"
	"heisenbug/identification/internal/pkg/mocker"
	"heisenbug/identification/internal/pkg/model"
	desc "heisenbug/identification/pkg/api/identification"
)

func (i *Implementation) UpgradeV3(ctx context.Context, req *desc.UpgradeRequest) (*desc.UpgradeResponse, error) {
	if err := validateUpgradeV2(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := i.usersClient.GetInfo(ctx, &users.GetInfoRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var resp *model.IdentificationResponse
	if mocker.UseMock(ctx) {
		resp, err = i.complexClient.Identification(ctx, info.Phone)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	} else {
		resp, err = i.externalClient.Identification(ctx, info.Phone)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	if resp.Status != 0 {
		return nil, status.Error(codes.Aborted, fmt.Sprintf("Identification failed with status %d", resp.Status))
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

func validateUpgradeV3(req *desc.UpgradeRequest) error {
	err := validation.ValidateStruct(req,
		validation.Field(&req.UserId, validation.Required),
	)
	return err
}
