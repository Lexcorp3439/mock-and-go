// Code generated by scratch. You must modify it.

package public

import (
	desc "heisenbug/external/pkg/api/public"
)

type Implementation struct {
	desc.UnimplementedPublicServer
}

// NewPublicService return new instance of Implementation.
func NewPublicService() *Implementation {
	return &Implementation{}
}
