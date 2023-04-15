package identification

import (
	"context"

	"e2e/internal/pb/users"
	p "e2e/internal/provider"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func (s *Suite) TestSuccessIdentificationWithWireMock(t provider.T) {
	t.Title("Успешная идентификация пользователя через WireMock")
	ctx := p.WithProvider(context.Background(), t)

	// Повышение идентификации через wiremock и проверка на изменение
	s.Steps.Identification.UpgradeIdentificationLevel(ctx, t, s.User.GetUserId())
	userInfo := s.Steps.Users.GetUserIdentificationLevel(ctx, t, s.User.GetUserId())
	s.Steps.Users.CheckUserIdentificationLevel(ctx, t, users.Level_ANON, userInfo.Level)
}
