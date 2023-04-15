package steps

import (
	"context"
	"e2e/internal/pb/identification"
	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

type IdentificationSteps struct {
	client identification.IdentificationClient
}

func NewIdentificationSteps(client identification.IdentificationClient) *IdentificationSteps {
	return &IdentificationSteps{
		client: client,
	}
}

func (s *IdentificationSteps) UpgradeIdentificationLevel(ctx context.Context, t provider.T, userID int32) {
	t.WithNewStep("Повышаем уровень идентификации пользователя", func(sCtx provider.StepCtx) {
		resp, err := s.client.Upgrade(ctx, &identification.UpgradeRequest{
			UserId: userID,
		})
		require.NoError(t.RealT(), err)
		require.NotNil(t.RealT(), resp)
		require.NotEmpty(t.RealT(), resp.GetIdentificationId())
	})
}

func (s *IdentificationSteps) UpgradeV2IdentificationLevel(ctx context.Context, t provider.T, userID int32, mockExternal bool) {
	t.WithNewStep("Повышаем уровень идентификации пользователя v2", func(sCtx provider.StepCtx) {
		if mockExternal {
			md := metadata.MD{}
			md.Set("mock", uuid.NewString())
			ctx = metadata.NewOutgoingContext(ctx, md)
		}
		resp, err := s.client.UpgradeV2(ctx, &identification.UpgradeRequest{
			UserId: userID,
		})
		require.NoError(t.RealT(), err)
		require.NotNil(t.RealT(), resp)
		require.NotEmpty(t.RealT(), resp.GetIdentificationId())
	})
}

func (s *IdentificationSteps) UpgradeV3IdentificationLevel(ctx context.Context, t provider.T, userID int32, mockExternal bool) {
	t.WithNewStep("Повышаем уровень идентификации пользователя v2", func(sCtx provider.StepCtx) {
		if mockExternal {
			md := metadata.MD{}
			md.Set("mock", uuid.NewString())
			ctx = metadata.NewOutgoingContext(ctx, md)
		}
		resp, err := s.client.UpgradeV3(ctx, &identification.UpgradeRequest{
			UserId: userID,
		})
		require.NoError(t.RealT(), err)
		require.NotNil(t.RealT(), resp)
		require.NotEmpty(t.RealT(), resp.GetIdentificationId())
	})
}
