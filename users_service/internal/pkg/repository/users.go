package repository

import (
	"context"
	"heisenbug/users/internal/pkg/model"
	"heisenbug/users/internal/pkg/store"
)

// Users repo interface
type Users interface {
	Add(ctx context.Context, fio, phone string, age int32) (int32, error)
	Delete(ctx context.Context, id int32) error
	Get(ctx context.Context, id int32) (*model.User, error)
	Upgrade(ctx context.Context, id, level int32) (*model.User, error)
}

type usersRepo struct {
	store *store.Storage
}

func NewUsersRepository(db *store.Storage) Users {
	return &usersRepo{store: db}
}

func (r *usersRepo) Add(ctx context.Context, fio, phone string, age int32) (int32, error) {
	stmt := store.PSQL().
		Insert(model.UsersTable).
		Columns("fio", "phone", "age").
		Values(fio, phone, age).
		Suffix("RETURNING id")

	var id int32
	err := r.store.Getx(ctx, stmt, &id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *usersRepo) Delete(ctx context.Context, id int32) error {
	stmt := store.PSQL().
		Delete(model.UsersTable).
		Where("id = ?", id)

	err := r.store.Execx(ctx, stmt)
	return err
}

func (r *usersRepo) Get(ctx context.Context, id int32) (*model.User, error) {
	stmt := store.PSQL().
		Select("*").
		From(model.UsersTable).
		Where("id = ?", id)

	var result model.User
	err := r.store.Getx(ctx, stmt,
		&result.ID,
		&result.FIO,
		&result.Phone,
		&result.Age,
		&result.Level,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *usersRepo) Upgrade(ctx context.Context, id, level int32) (*model.User, error) {
	stmt := store.PSQL().
		Update(model.UsersTable).
		Set("level", level).
		Where("id = ?", id).
		Suffix("RETURNING *")

	var result model.User
	err := r.store.Getx(ctx, stmt,
		&result.ID,
		&result.FIO,
		&result.Phone,
		&result.Age,
		&result.Level,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
