package repository_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_UsersRepo(t *testing.T) {
	ctx := context.Background()

	fio := uuid.NewString()
	phone := "79819344567"
	age := int32(18)

	userID, err := repo.Add(ctx, fio, phone, age)
	require.NoError(t, err)
	require.NotZero(t, userID)

	user, err := repo.Get(ctx, userID)
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, fio, user.FIO)
	require.Equal(t, phone, user.Phone)
	require.Equal(t, age, user.Age)
	require.EqualValues(t, 0, user.Level)

	fullUser, err := repo.Upgrade(ctx, userID, 1)
	require.NoError(t, err)
	require.NotNil(t, fullUser)
	require.EqualValues(t, 1, fullUser.Level)

	err = repo.Delete(ctx, userID)
	require.NoError(t, err)

	user, err = repo.Get(ctx, userID)
	require.Error(t, err)
	require.Nil(t, user)
}
