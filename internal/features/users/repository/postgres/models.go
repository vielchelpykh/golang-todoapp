package users_postgres_repository

import "github.com/vielchelpykh/golang-todoapp/internal/core/domains"

type UserModel struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func userDomainsFromModels(users []UserModel) []domains.User {
	userDomains := make([]domains.User, len(users))

	for i, user := range users {
		userDomains[i] = domains.NewUser(
			user.ID,
			user.Version,
			user.FullName,
			user.PhoneNumber,
		)
	}

	return userDomains
}
