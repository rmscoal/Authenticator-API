package usecase

import (
	"context"

	"github.com/rmscoal/Authenticator-API/internal/entity"
)

// UserUseCase -.
type UserUseCase struct {
	repo UserRepo
}

// New -.
func New(r UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

// Find - getting user from username and password
func (uc *UserUseCase) Find(ctx context.Context, user entity.User) (entity.User, error) {
	user, err := uc.repo.GetUser(ctx, user)
	if err != nil {
		// return entity.User{}, fmt.Errorf("UserUseCase - Find - s.repo.GetUser: %w", err)
		return entity.User{}, err
	}

	return user, nil
}
