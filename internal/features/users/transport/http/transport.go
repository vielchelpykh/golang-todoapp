package users_transport_http

import (
	"context"
	"net/http"

	"github.com/vielchelpykh/golang-todoapp/internal/core/domains"
	core_http_server "github.com/vielchelpykh/golang-todoapp/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
	usersService UsersService
}

type UsersService interface {
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
		patch domains.UserPatch,
	) (domains.User, error)
}

func NewUsersHTTPHandler(usersService UsersService) *UsersHTTPHandler {
	return &UsersHTTPHandler{usersService: usersService}
}

func (h *UsersHTTPHandler) Routers() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: h.GetUsers,
			// Middleware: []core_http_middleware.Middleware{
			// 	core_http_middleware.Dummy("get users middleware"),
			// },
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: h.GetUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/{id}",
			Handler: h.DeleteUser,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users/{id}",
			Handler: h.PatchUser,
		},
	}
}
