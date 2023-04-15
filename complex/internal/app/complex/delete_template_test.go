// Code generated by scratch. You must modify it.

package complex

import (
	"context"
	"heisenbug/complex/internal/pkg/model"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	desc "heisenbug/complex/pkg/api/complex"
)

func TestImplementation_DeleteTemplate_Success(t *testing.T) {
	t.Parallel()
	t.Run("", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		templateName := uuid.NewString()

		err := TemplateRepo.CreateTemplate(ctx,
			&model.Template{
				Name:   templateName,
				Status: 0,
			})
		require.NoError(t, err)

		_, err = TemplateRepo.GetTemplate(ctx, templateName)
		require.NoError(t, err)

		resp, err := ComplexService.DeleteTemplate(ctx, &desc.DeleteTemplateRequest{TemplateName: templateName})
		require.NoError(t, err)
		require.Equal(t, &desc.DeleteTemplateResponse{TemplateName: templateName}, resp)

		_, err = TemplateRepo.GetTemplate(ctx, templateName)
		require.Error(t, err)
	})
}