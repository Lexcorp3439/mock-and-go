package mocker

import (
	"github.com/google/uuid"
	"heisenbug/identification/internal/pkg/model"
	"strings"
)

func Mocker(phone string) *model.IdentificationResponse {
	resp := &model.IdentificationResponse{
		Id: uuid.NewString(),
	}

	switch {
	case strings.HasSuffix(phone, "987"):
		resp.Status = 1
	case strings.HasSuffix(phone, "999"):
		resp.Status = 2
	case strings.HasSuffix(phone, "111"):
		resp.Status = 3
	default:
		resp.Status = 0
	}
	return resp
}
