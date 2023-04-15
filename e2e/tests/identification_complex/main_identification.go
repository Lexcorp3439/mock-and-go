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

	// Создание шаблона и привязка к пользователю
	templateName := s.Steps.Complex.GenerateTemplate(ctx, t, 0)
	s.Steps.Complex.BindPhone(ctx, t, templateName, phone)
}

func (s *Suite) AfterEach(t provider.T) {
	s.AfterEachUserTest(t)
}
