package steps

import (
	"context"
	"e2e/internal/pb/users"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/stretchr/testify/require"
)

type UsersSteps struct {
	client users.UsersClient
}

func NewUsersSteps(client users.UsersClient) *UsersSteps {
	return &UsersSteps{
		client: client,
	}
}

func (s *UsersSteps) GetUserIdentificationLevel(ctx context.Context, t provider.T, userID int32) *users.GetInfoResponse {
	var response *users.GetInfoResponse
	var err error

	t.WithNewStep("Получаем идентификационный уровень пользователя", func(sCtx provider.StepCtx) {
		response, err = s.client.GetInfo(ctx, &users.GetInfoRequest{
			UserId: userID,
		})
		require.NoError(t.RealT(), err)
		require.NotNil(t.RealT(), response)
	})
	return response
}

func (s *UsersSteps) CheckUserIdentificationLevel(ctx context.Context, t provider.T, expectedLevel, level users.Level) {
	t.WithNewStep("Проверяем уровень идентификации пользователя", func(sCtx provider.StepCtx) {
		require.Equal(t.RealT(), expectedLevel, level)
	})
}

func (s *UsersSteps) CreateUser(ctx context.Context, t provider.T, fio string, phone string, age int32) *users.CreateUserResponse {
	var user *users.CreateUserResponse

	t.WithNewStep("Создаем нового пользователя для теста", func(sCtx provider.StepCtx) {
		response, err := s.client.CreateUser(ctx, &users.CreateUserRequest{
			Fio:   fio,
			Phone: phone,
			Age:   age,
		})
		require.NoError(t.RealT(), err)
		require.NotNil(t.RealT(), response)
		require.NotEmpty(t.RealT(), response.UserId)
		user = response
	})
	return user
}

func (s *UsersSteps) RemoveUser(ctx context.Context, t provider.T, userID int32) {
	t.WithNewStep("Удаляем пользователя", func(sCtx provider.StepCtx) {
		response, err := s.client.RemoveUser(ctx, &users.RemoveUserRequest{
			UserId: userID,
		})
		require.NoError(t.RealT(), err)
		require.NotNil(t.RealT(), response)
	})
}
