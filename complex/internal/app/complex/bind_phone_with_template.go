// Code generated by scratch. You must modify it.

package complex

import (
	"context"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
	"heisenbug/complex/internal/pkg/model"
	desc "heisenbug/complex/pkg/api/complex"
)

// BindPhoneWithTemplate stub. Please implement it.
func (i *Implementation) BindPhoneWithTemplate(ctx context.Context, req *desc.BindPhoneWithTemplateRequest) (*emptypb.Empty, error) {
	err := i.templateRepo.MarkOldBindsAsExpired(ctx, req.PhoneNumber)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "MarkOldBindsAsExpired err %s", err.Error())
	}

	err = i.templateRepo.BindPhoneWithTemplateName(ctx, &model.TemplateBinding{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "BindPhoneWithTemplateName err %s", err.Error())
	}
	return &emptypb.Empty{}, nil
}
