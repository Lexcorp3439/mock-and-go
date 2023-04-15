package public

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"heisenbug/complex/internal/pkg/model"
	desc "heisenbug/complex/pkg/api/public"
)

func TestImplementation_GenerateSuccessResult_Success(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("With existed template", func(t *testing.T) {
		t.Parallel()
		phoneNumber := "70000000001"

		templateName := uuid.NewString()
		protoTemplate := &model.Template{
			Name:        templateName,
			Description: nil,
			Status:      1,
		}
		err := TemplateRepo.CreateTemplate(ctx, protoTemplate)
		require.NoError(t, err)

		err = TemplateRepo.BindPhoneWithTemplateName(ctx, &model.TemplateBinding{
			Name:        templateName,
			PhoneNumber: phoneNumber,
		})
		require.NoError(t, err)

		form, err := PublicService.Identification(ctx, &desc.IdentificationRequest{
			Phone: phoneNumber,
		})
		require.NoError(t, err)
		require.EqualValues(t, form.Status, 1)
	})

	t.Run("Without template", func(t *testing.T) {
		t.Parallel()
		phoneNumber := "70000000002"
		form, err := PublicService.Identification(ctx, &desc.IdentificationRequest{
			Phone: phoneNumber,
		})
		require.NoError(t, err)
		require.EqualValues(t, form.Status, 0)
	})
}
