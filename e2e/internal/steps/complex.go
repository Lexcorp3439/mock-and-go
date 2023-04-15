package steps

import (
	"context"
	"e2e/internal/pb/complex"
	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/stretchr/testify/require"
)

type ComplexSteps struct {
	client complex.ComplexClient
}

func NewComplexSteps(client complex.ComplexClient) *ComplexSteps {
	return &ComplexSteps{
		client: client,
	}
}

func (s *ComplexSteps) GenerateTemplate(
	ctx context.Context,
	t provider.T,
	status int32,
) string {
	templateName := uuid.NewString()
	req := &complex.TemplateRequest{
		TemplateName: templateName,
		Status:       status,
	}
	resp, err := s.client.CreateOrUpdateTemplate(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	return templateName
}

func (s *ComplexSteps) BindPhone(
	ctx context.Context,
	t provider.T,
	templateName string,
	phone string,
) {
	req := &complex.BindPhoneWithTemplateRequest{
		Name:        templateName,
		PhoneNumber: phone,
	}
	resp, err := s.client.BindPhoneWithTemplate(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
