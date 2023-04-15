package store

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"heisenbug/users/internal/config"
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
	sql, args, err := sq.ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Getx(ctx context.Context, sq Sqliser, dest ...any) error {
	sql, args, err := sq.ToSql()
	if err != nil {
		return err
	}

	err = s.db.QueryRow(ctx, sql, args...).Scan(dest...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Close() {
	s.db.Close()
}
