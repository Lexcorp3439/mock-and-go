package users_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"heisenbug/users/internal/app/users"
	"testing"

	desc "heisenbug/users/pkg/api/users"
)

func TestImplementation_RemoveUser(t *testing.T) {
	ctx := context.Background()
	app := users.NewUsersService(repo)

	createResp, err := app.CreateUser(ctx, &desc.CreateUserRequest{
		Fio:   uuid.NewString(),
		Phone: uuid.NewString(),
		Age:   20,
	})
	require.NoError(t, err)
	require.NotNil(t, createResp)

	resp, err := app.RemoveUser(ctx, &desc.RemoveUserRequest{
		UserId: createResp.UserId,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
}
