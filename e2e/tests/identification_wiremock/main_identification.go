package identification

import (
	"context"
	"e2e/internal"
	p "e2e/internal/provider"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

type Suite struct {
	internal.Suite
}

func (s *Suite) BeforeAll(_ provider.T) {
	s.PrepareUserSuite()
}

func (s *Suite) BeforeEach(t provider.T) {
	ctx := p.WithProvider(context.Background(), t)

	// Генерация пользователя
	fio, phone, age := s.Steps.Stabs.GetStubUsersData()
	s.User = s.Steps.Users.CreateUser(ctx, t, fio, phone, age)

	// Создание сценария в wiremock
	s.Steps.Wiremock.MockIdentification(ctx, t, phone, 3)
}

func (s *Suite) AfterEach(t provider.T) {
	s.AfterEachUserTest(t)
}
