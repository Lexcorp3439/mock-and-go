// Code generated by scratch. You must modify it.

package public

import (
	"context"
	"github.com/google/uuid"
	desc "heisenbug/external/pkg/api/public"
	"strings"
)

// Identification stub. Please implement it.
func (i *Implementation) Identification(ctx context.Context, req *desc.IdentificationRequest) (*desc.IdentificationResponse, error) {
	resp := &desc.IdentificationResponse{
		Id: uuid.NewString(),
	}

	phone := req.GetPhone()

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

	return resp, nil
}