package users_test

import (
	"heisenbug/users/internal/pkg/repository"
	"heisenbug/users/internal/pkg/store"
	"os"
	"testing"
)

var (
	repo repository.Users
)

func TestMain(m *testing.M) {
	db, err := store.ConnectToPostgres()
	if err != nil {
		panic(err)
	}

	storage := store.NewStorage(db)
	repo = repository.NewUsersRepository(storage)

	code := m.Run()

	os.Exit(code)
}
