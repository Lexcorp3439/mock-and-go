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

func TestImplementation_GetTemplateByName_Success(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	templateName := uuid.NewString()

	_, err := ComplexService.CreateOrUpdateTemplate(ctx, &desc.TemplateRequest{
		TemplateName: templateName,
		Status:       0,
	})
	require.NoError(t, err)

	response, err := ComplexService.GetTemplateByName(ctx, &desc.GetTemplateByNameRequest{TemplateName: templateName})
	require.NoError(t, err)
	require.EqualValues(t, 0, response.Template.Status)

}

func TestImplementation_GetTemplateByName_Negative(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name              string
		req               *desc.GetTemplateByNameRequest
		wantErrStatusCode codes.Code
	}
	tests := []testCase{
		{
			name: "not exist template name -> err not found",
			req: &desc.GetTemplateByNameRequest{
				TemplateName: uuid.NewString(),
			},
			wantErrStatusCode: codes.Aborted,
		},
		{
			name: "broken template name -> internal err",
			req: &desc.GetTemplateByNameRequest{
				TemplateName: "x\000",
			},
			wantErrStatusCode: codes.Aborted,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			_, err := ComplexService.GetTemplateByName(ctx, tt.req)
			require.Error(t, err)
			require.Equal(t, status.Code(err), tt.wantErrStatusCode)
		})
	}
}
