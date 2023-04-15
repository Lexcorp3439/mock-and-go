package app

import (
	"heisenbug/identification/internal/pb/users"
	compl "heisenbug/identification/internal/pkg/complex"
	"heisenbug/identification/internal/pkg/external"
	desc "heisenbug/identification/pkg/api/identification"
)

type Implementation struct {
	desc.UnsafeIdentificationServer
	usersClient    users.UsersClient
	externalClient external.Client
	complexClient  compl.Client
}

// NewIdentificationService return new instance of Implementation.
func NewIdentificationService(
	usersClient users.UsersClient,
	externalClient external.Client,
	complexClient compl.Client,
) *Implementation {
	return &Implementation{
		usersClient:    usersClient,
		externalClient: externalClient,
		complexClient:  complexClient,
	}
}
