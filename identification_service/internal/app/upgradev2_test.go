package app_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
	"heisenbug/identification/internal/pkg/mocker"
	"heisenbug/identification/internal/pkg/model"
	"heisenbug/identification/pkg/mocks/external_mock"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	serv "heisenbug/identification/internal/app"
	"heisenbug/identification/internal/pb/users"
	desc "heisenbug/identification/pkg/api/identification"
	"heisenbug/identification/pkg/mocks/users_mock"
)

func TestImplementation_UpgradeV2_Stg(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	err := os.Setenv("env", mocker.MockEnv)
	require.NoError(t, err)

	usersClient := users_mock.NewMockUsersClient(ctrl)
	externalClient := external_mock.NewMockClient(ctrl)

	app := serv.NewIdentificationService(usersClient, externalClient, nil)

	userID := int32(10010101)

	usersClient.EXPECT().GetInfo(gomock.Any(), gomock.Any()).Return(&users.GetInfoResponse{
		UserId: userID,
		Fio:    "Petrov Petr Petrovich",
		Phone:  defaultPhone,
	}, nil)

	externalClient.EXPECT().Identification(gomock.Any(), gomock.Any()).Times(0)

	usersClient.EXPECT().UpdateUserLevel(gomock.Any(), gomock.Any()).Return(&users.UpdateUserLevelResponse{}, nil)

	md := make(metadata.MD)
	md.Set(mocker.MockKey, uuid.NewString())
	ctx = metadata.NewIncomingContext(ctx, md)
	resp, err := app.UpgradeV2(ctx, &desc.UpgradeRequest{
		UserId: userID,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotEmpty(t, resp.IdentificationId)
}

func TestImplementation_UpgradeV2_Prod(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	err := os.Setenv("env", "prod")
	require.NoError(t, err)

	usersClient := users_mock.NewMockUsersClient(ctrl)
	externalClient := external_mock.NewMockClient(ctrl)

	app := serv.NewIdentificationService(usersClient, externalClient, nil)

	userID := int32(1919)
	identificationId := uuid.NewString()

	usersClient.EXPECT().GetInfo(ctx, gomock.Any()).Return(&users.GetInfoResponse{
		UserId: userID,
		Fio:    "Petrov Petr Petrovich",
		Phone:  defaultPhone,
	}, nil)

	externalClient.EXPECT().Identification(ctx, defaultPhone).Return(&model.IdentificationResponse{
		Id:     identificationId,
		Status: 0,
	}, nil)

	usersClient.EXPECT().UpdateUserLevel(ctx, gomock.Any()).Return(&users.UpdateUserLevelResponse{}, nil)

	resp, err := app.UpgradeV2(ctx, &desc.UpgradeRequest{
		UserId: userID,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, identificationId, resp.IdentificationId)
}
