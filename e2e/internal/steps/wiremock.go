package steps

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/stretchr/testify/require"
	wire "github.com/walkerus/go-wiremock"
)

type WiremockSteps struct {
	client *wire.Client
}

func NewWiremockSteps(client *wire.Client) *WiremockSteps {
	return &WiremockSteps{
		client: client,
	}
}

func (w *WiremockSteps) MockIdentification(ctx context.Context, t provider.T, phone string, status int) {
	bodyPattern := fmt.Sprintf("{ \"Phone\": \"%s\"}", phone)
	t.WithNewStep("Подготовка мока идентификации через Озон", func(sCtx provider.StepCtx) {
		err := w.client.StubFor(wire.Post(wire.URLPathEqualTo("/identification")).
			WithBodyPattern(wire.EqualToJson(bodyPattern)).
			WillReturnJSON(
				map[string]interface{}{
					"id":     uuid.NewString(),
					"status": status,
				},
				map[string]string{"Content-Type": "application/json"},
				200,
			).
			AtPriority(1))
		require.NoError(t.RealT(), err)
	})

}
