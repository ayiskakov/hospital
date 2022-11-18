package storage

import (
	"context"

	"github.com/khanfromasia/hospital/backend/internal/entity"
	"github.com/pkg/errors"
)

const (
	userCreateSQL = `--- name: UserCreate :one
		INSERT INTO "public.user" (email, name, surname, salary, phone, cname)
		VALUES ($1, $2, $3, $4, $5, $6)`

	userUpdateSQL = `--- name: UserUpdate :one
		UPDATE "public.user"
		SET
			email = COALESCE($1, email),
			name = COALESCE($2, name),
			surname = COALESCE($3, surname),
			salary = COALESCE($4, salary),
			phone = COALESCE($5, phone),
			cname = COALESCE($6, cname)
		WHERE
			email = $7
		RETURNING email, name, surname, salary, phone, cname`

	userDeleteSQL = `--- name: UserDelete :one
		DELETE FROM "public.user"
		WHERE email = $1`

	userGetAllSQL = `--- name: UserGetAll :many
		SELECT
			U.email, U.name, U.surname, U.salary, U.phone, C.cname, C.Population
		FROM
			"public.user" U
		INNER JOIN
			"public.country" C
		ON
			U.cname = C.cname`
)

// UserCreate creates a new user.
func (q *Queries) UserCreate(ctx context.Context, arg entity.User) (entity.User, error) {
	_, err := q.db.Exec(
		ctx,
		userCreateSQL,
		arg.Email,
		arg.Name,
		arg.Surname,
		arg.Salary,
		arg.Phone,
		arg.Country.Cname,
	)

	if err != nil {
		return entity.User{}, errors.Wrap(err, "failed to create user")
	}

	return arg, nil
}

// UserUpdate updates a user.
func (q *Queries) UserUpdate(ctx context.Context, email string, arg entity.UserUpdate) (entity.User, error) {
	row := q.db.QueryRow(
		ctx,
		userUpdateSQL,
		arg.Email,
		arg.Name,
		arg.Surname,
		arg.Salary,
		arg.Phone,
		arg.Cname,
		email,
	)

	var user entity.User
	err := row.Scan(
		&user.Email,
		&user.Name,
		&user.Surname,
		&user.Salary,
		&user.Phone,
		&user.Country.Cname,
	)

	if err != nil {
		return entity.User{}, errors.Wrap(err, "failed to update user")
	}

	return user, nil
}

// UserDelete deletes a user.
func (q *Queries) UserDelete(ctx context.Context, email string) error {
	_, err := q.db.Exec(ctx, userDeleteSQL, email)

	if err != nil {
		return errors.Wrap(err, "failed to delete user")
	}

	return nil
}

// UserGetAll returns all users.
func (q *Queries) UserGetAll(ctx context.Context) ([]entity.User, error) {
	rows, err := q.db.Query(ctx, userGetAllSQL)

	if err != nil {
		return nil, errors.Wrap(err, "failed to get all users")
	}

	defer rows.Close()

	var users []entity.User

	for rows.Next() {
		var user entity.User

		err = rows.Scan(
			&user.Email,
			&user.Name,
			&user.Surname,
			&user.Salary,
			&user.Phone,
			&user.Country.Cname,
			&user.Country.Population,
		)

		if err != nil {
			return nil, errors.Wrap(err, "failed to scan user")
		}

		users = append(users, user)
	}

	return users, nil
}
