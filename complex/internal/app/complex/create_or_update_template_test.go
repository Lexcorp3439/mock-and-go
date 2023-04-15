package complex

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	desc "heisenbug/complex/pkg/api/complex"
)

type testItem struct {
	name               string
	req                *desc.TemplateRequest
	initialDescription string
	otherDescription   string
}

func TestImplementation_CreateOrUpdateTemplate_Validation(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	templateName := uuid.NewString()

	tests := []testItem{
		{
			name: "Status: -1",
			req: &desc.TemplateRequest{
				TemplateName: templateName,
				Status:       -1,
			},
		},
		{
			name: "Without TemplateName",
			req: &desc.TemplateRequest{
				Status: 1,
			},
		},
	}
	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, err := ComplexService.CreateOrUpdateTemplate(ctx, tc.req)
			require.Error(t, err)
			require.Equal(t, codes.InvalidArgument, status.Code(err))
		})
	}
}

func TestImplementation_CreateOrUpdateTemplate_Create_Success(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	templateName := uuid.NewString()

	tests := []testItem{
		{
			name: "With hex message & description",
			req: &desc.TemplateRequest{
				TemplateName: templateName,
				Status:       0,
			},
			initialDescription: "description",
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.req.TemplateName = uuid.NewString()
			checkTemplateNotExist(ctx, t, tc.req.TemplateName)

			tc.req.Description = &tc.initialDescription

			response, err := ComplexService.CreateOrUpdateTemplate(ctx, tc.req)

			require.NoError(t, err)
			require.Equal(t, tc.req.TemplateName, response.Template.TemplateName)
			require.Equal(t, tc.initialDescription, response.Template.Description)

			tmp, err := TemplateRepo.GetTemplate(ctx, tc.req.TemplateName)
			require.NoError(t, err)
			require.Equal(t, response.Template.TemplateName, tmp.Name)
			require.EqualValues(t, tc.req.Status, tmp.Status)
			require.Equal(t, response.Template.Description, *tmp.Description)
		})
	}
}

func checkTemplateNotExist(ctx context.Context, t *testing.T, templateName string) {
	t.Helper()
	_, err := TemplateRepo.GetTemplate(ctx, templateName)
	require.Error(t, err)
}
