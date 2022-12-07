package repo

import (
	"context"
	"fmt"

	"github.com/rmscoal/Authenticator-API/internal/entity"
	"github.com/rmscoal/Authenticator-API/pkg/postgres"
)

// UserRepo -.
type UserRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

// GetUser -.
func (r *UserRepo) GetUser(ctx context.Context, user entity.User) (entity.User, error) {
	sql, args, err := r.Builder.
		Select("username, firstname, lastname").
		From("users").
		Where("username = ? AND pass = ?", user.Username, user.Password).
		Limit(1).
		ToSql()
	if err != nil {
		return entity.User{}, fmt.Errorf("UserRepo - GetUser - r.Builder: %w", err)
	}

	// QueryRow expects at least a single row in result of the query.
	row := r.Pool.QueryRow(ctx, sql, args...)

	e := entity.User{}

	// If there are no result from the query, row.Scan will produce an
	// error.
	err = row.Scan(&e.Username, &e.FirstName, &e.LastName)
	if err != nil {
		return entity.User{}, fmt.Errorf("UserRepo - GetUser - row.Scan: %w", err)
	}

	return e, nil
}
