package users_service

import (
	"context"
	"fmt"

	domain "github.com/vielchelpykh/golang-todoapp/internal/core/domains"
)

func (s *UsersService) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user: %w", err)
	}

	createdUser, err := s.usersRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user in repository: %w", err)
	}

	return createdUser, nil
}
