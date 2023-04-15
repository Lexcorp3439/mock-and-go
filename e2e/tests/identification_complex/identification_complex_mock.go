package identification

import (
	"context"

	"e2e/internal/pb/users"
	p "e2e/internal/provider"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func (s *Suite) TestSuccessIdentificationWithComplexMock(t provider.T) {
	t.Title("Успешная идентификация пользователя через сложный мок")
	ctx := p.WithProvider(context.Background(), t)

	// Повышение идентификации и проверка на изменение
	s.Steps.Identification.UpgradeV3IdentificationLevel(ctx, t, s.User.GetUserId(), true)
	userInfo := s.Steps.Users.GetUserIdentificationLevel(ctx, t, s.User.GetUserId())
	s.Steps.Users.CheckUserIdentificationLevel(ctx, t, users.Level_FULL, userInfo.Level)
}
