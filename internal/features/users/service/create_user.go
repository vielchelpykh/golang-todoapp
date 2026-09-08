package users_service

import (
	"context"
	"fmt"

	"github.com/vielchelpykh/golang-todoapp/internal/core/domains"
)

func (s *UsersService) CreateUser(
	ctx context.Context,
	user domains.User,
) (domains.User, error) {
	if err := user.Validate(); err != nil {
		return domains.User{}, fmt.Errorf("validate user: %w", err)
	}

	createdUser, err := s.usersRepository.CreateUser(ctx, user)
	if err != nil {
		return domains.User{}, fmt.Errorf("create user in repository: %w", err)
	}

	return createdUser, nil
}
