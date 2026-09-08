package users_service

import (
	"context"
	"fmt"

	"github.com/vielchelpykh/golang-todoapp/internal/core/domains"
)

func (s *UsersService) GetUser(
	ctx context.Context,
	id int,
) (domains.User, error) {
	user, err := s.usersRepository.GetUser(ctx, id)
	if err != nil {
		return domains.User{}, fmt.Errorf("get user from repository: %w", err)
	}

	return user, nil
}
