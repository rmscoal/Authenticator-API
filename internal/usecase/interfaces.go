package usecase

import (
	"context"

	"github.com/rmscoal/Authenticator-API/internal/entity"
)

type (
	// User
	User interface {
		Find(context.Context, entity.User) (entity.User, error)
		Store(context.Context, entity.User) error
	}

	// UserRepo
	UserRepo interface {
		GetUser(context.Context, entity.User) (entity.User, error)
		StoreUser(context.Context, entity.User) error
	}
)
