package internal

import (
	"context"
	"e2e/internal/connection"
	"e2e/internal/connection/grpc"
	"e2e/internal/pb/users"
	p "e2e/internal/provider"
	"e2e/internal/steps"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type Suite struct {
	suite.Suite

	User *users.CreateUserResponse

	Steps struct {
		Identification *steps.IdentificationSteps
		Users          *steps.UsersSteps
		Complex        *steps.ComplexSteps
		Stabs          *steps.StubSteps
		Wiremock       *steps.WiremockSteps
	}
}

func (s *Suite) PrepareUserSuite() {
	helper := grpc.NewGrpcHelper()
	s.Steps.Identification = steps.NewIdentificationSteps(connection.NewIdentificationClient(helper))
	s.Steps.Users = steps.NewUsersSteps(connection.NewUsersClient(helper))
	s.Steps.Complex = steps.NewComplexSteps(connection.NewComplexClient(helper))
	s.Steps.Stabs = steps.NewStubSteps()
	s.Steps.Wiremock = steps.NewWiremockSteps(connection.NewWiremockClient())

}

func (s *Suite) AfterEachUserTest(t provider.T) {
	if s.User == nil {
		return
	}
	ctx := p.WithProvider(context.Background(), t)
	s.Steps.Users.RemoveUser(ctx, t, s.User.GetUserId())
}
