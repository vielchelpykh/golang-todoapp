package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/vielchelpykh/golang-todoapp/internal/core/domains"
)

func (r *UsersRepository) CreateUser(
	ctx context.Context,
	user domains.User,
) (domains.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
        INSERT INTO todoapp.users (full_name, phone_number)
        VALUES ($1, $2)
        RETURNING id, version, full_name, phone_number
    `

	var userModel UserModel
	err := r.pool.QueryRow(ctx, query, user.FullName, user.PhoneNumber).Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PhoneNumber,
	)
	if err != nil {
		return domains.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domains.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.FullName,
		userModel.PhoneNumber,
	)

	return userDomain, nil
}
