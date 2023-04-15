package app_test

import (
	"context"
	"heisenbug/identification/internal/pkg/model"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	serv "heisenbug/identification/internal/app"
	"heisenbug/identification/internal/pb/users"
	desc "heisenbug/identification/pkg/api/identification"
	"heisenbug/identification/pkg/mocks/external_mock"
	"heisenbug/identification/pkg/mocks/users_mock"
)

func TestImplementation_Upgrade(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	usersClient := users_mock.NewMockUsersClient(ctrl)
	externalClient := external_mock.NewMockClient(ctrl)

	app := serv.NewIdentificationService(usersClient, externalClient, nil)

	userID := int32(100)
	identificationId := uuid.NewString()

	getInfoCall := usersClient.EXPECT().GetInfo(ctx, gomock.Any()).Return(&users.GetInfoResponse{
		UserId: userID,
		Fio:    "Petrov Petr Petrovich",
		Phone:  defaultPhone,
		Level:  users.Level_ANON,
	}, nil)

	identificationCall := externalClient.EXPECT().Identification(ctx, defaultPhone).
		Return(&model.IdentificationResponse{
			Id:     identificationId,
			Status: 0,
		}, nil)

	UpdateUserLevelCall := usersClient.EXPECT().UpdateUserLevel(ctx, gomock.Any()).Return(&users.UpdateUserLevelResponse{}, nil)

	gomock.InOrder(getInfoCall, identificationCall, UpdateUserLevelCall)

	resp, err := app.Upgrade(ctx, &desc.UpgradeRequest{
		UserId: userID,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, identificationId, resp.IdentificationId)
}
