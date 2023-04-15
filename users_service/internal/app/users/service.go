package users

import (
	"heisenbug/users/internal/pkg/repository"
	desc "heisenbug/users/pkg/api/users"
)

type Implementation struct {
	desc.UnsafeUsersServer

	repo repository.Users
}

// NewUsersService return new instance of Implementation.
func NewUsersService(
	repo repository.Users,
) *Implementation {
	return &Implementation{
		repo: repo,
	}
}
