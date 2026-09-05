package users_service

import (
	"context"

	"github.com/vielchelpykh/golang-todoapp/internal/core/domains"
)

type UsersService struct {
	usersRepository UsersRepository
}

func NewUsersService(usersRepository UsersRepository) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
	}
}

type UsersRepository interface {
	CreateUser(
		ctx context.Context,
		user domains.User,
	) (domains.User, error)

	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domains.User, error)

	GetUser(
		ctx context.Context,
		id int,
	) (domains.User, error)

	DeleteUser(
		ctx context.Context,
		id int,
	) error

	PatchUser(
		ctx context.Context,
		id int,
		user domains.User,
	) (domains.User, error)
}
