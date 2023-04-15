package store

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	pgx "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"heisenbug/complex/internal/config"
)

var (
	// ErrNotFound - запись не найдена
	ErrNotFound = errors.New("no rows in result set")
)

func ConnectToPostgres() (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), config.DatabaseDsn)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

// PSQL Postgres specific squirrel builder
func PSQL() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}

type Sqliser interface {
	ToSql() (string, []interface{}, error)
}
type Storage struct {
	db *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) *Storage {
	return &Storage{db: db}
}

func (s *Storage) Execx(ctx context.Context, sq Sqliser) error {
	query, args, err := sq.ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Getx(ctx context.Context, sq Sqliser, dest ...any) error {
	query, args, err := sq.ToSql()
	if err != nil {
		return err
	}
	err = s.db.QueryRow(ctx, query, args...).Scan(dest...)
	if err != nil {
		return err
	}
	return nil
}

func Select[T any](ctx context.Context, storage *Storage, sq Sqliser) ([]T, error) {
	sql, args, err := sq.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := storage.db.Query(ctx, sql, args...)
	return pgx.CollectRows(rows, pgx.RowToStructByName[T])
}

// DbError - DB error convertor
func DbError(err error) error {
	if err == sql.ErrNoRows {
		return ErrNotFound
	}

	return errors.WithStack(err)
}
