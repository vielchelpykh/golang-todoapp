package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	domain "github.com/vielchelpykh/golang-todoapp/internal/core/domains"
	core_errors "github.com/vielchelpykh/golang-todoapp/internal/core/errors"
	core_postgres_pool "github.com/vielchelpykh/golang-todoapp/internal/core/repository/postgres/pool"
)

// если вызывается метод, то пользователь существует, так как до этого его нужно получить
// если на момент выполнения метода его строку не получили, то был конфликт
func (r *UsersRepository) PatchUser(
	ctx context.Context,
	id int,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE todoapp.users
	SET 
		full_name=$1,
		phone_number=$2,
		version=version+1
	WHERE id=$3 and version=$4
	RETURNING
		id,
		version,
		full_name,
		phone_number;
	`

	row := r.pool.QueryRow(ctx, query, user.FullName, user.PhoneNumber, id, user.Version)

	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user with id='%d' concurrently accessed: %w",
				id,
				core_errors.ErrConflict,
			)
		} else {
			return domain.User{}, fmt.Errorf("scan error: %w", err)
		}
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.FullName,
		userModel.PhoneNumber,
	)

	return userDomain, nil
}
