package storage

import (
	"context"

	"github.com/khanfromasia/hospital/backend/internal/entity"
	"github.com/pkg/errors"
)

const (
	publicServantCreateSQL = `--- name: PublicServantCreate :one
		INSERT INTO "public.public_servant"
			(email, department)
		VALUES
			($1, $2)`

	publicServantGetAllSQL = `--- name: PublicServantGetAll :many
		SELECT
			P.email, P.department, U.name, U.surname, U.salary, U.phone, C.cname, C.Population
		FROM
			"public.public_servant" P
		INNER JOIN
			"public.users" U
		ON
			P.email = U.email
		INNER JOIN
			"public.country" C
		ON
			U.cname = C.cname`

	publicServantUpdateSQL = `--- name: PublicServantUpdate :one
		UPDATE
			"public.public_servant"
		SET
			department = COALESCE($1, department)
		WHERE
			email = $2
		RETURNING department`

	publicServantDeleteSQL = `--- name: PublicServantDelete :exec
		DELETE FROM
			"public.public_servant"
		WHERE
			email = $1`
)

// PublicServantCreate creates a new public servant
func (q *Queries) PublicServantCreate(ctx context.Context, arg entity.PublicServant) (entity.PublicServant, error) {
	_, err := q.db.Exec(
		ctx,
		publicServantCreateSQL,
		arg.User.Email,
		arg.Department,
	)

	if err != nil {
		return entity.PublicServant{}, errors.Wrap(err, "failed to create public servant")
	}

	return arg, nil
}

// PublicServantGetAll gets all public servants
func (q *Queries) PublicServantGetAll(ctx context.Context) ([]entity.PublicServant, error) {
	rows, err := q.db.Query(ctx, publicServantGetAllSQL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all public servants")
	}

	defer rows.Close()

	var publicServants []entity.PublicServant
	for rows.Next() {
		var publicServant entity.PublicServant
		err = rows.Scan(
			&publicServant.User.Email,
			&publicServant.Department,
			&publicServant.User.Name,
			&publicServant.User.Surname,
			&publicServant.User.Salary,
			&publicServant.User.Phone,
			&publicServant.User.Country.Cname,
			&publicServant.User.Country.Population,
		)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan public servant")
		}

		publicServants = append(publicServants, publicServant)
	}

	return publicServants, nil
}

// PublicServantUpdate updates a public servant
func (q *Queries) PublicServantUpdate(ctx context.Context, email string, arg entity.PublicServantUpdate) (entity.PublicServant, error) {
	row := q.db.QueryRow(
		ctx,
		publicServantUpdateSQL,
		arg.Department,
		email,
	)

	var publicServant entity.PublicServant

	err := row.Scan(
		&publicServant.Department,
	)

	if err != nil {
		return entity.PublicServant{}, errors.Wrap(err, "failed to update public servant")
	}

	publicServant.User.Email = email

	return publicServant, nil
}

// PublicServantDelete deletes a public servant
func (q *Queries) PublicServantDelete(ctx context.Context, email string) error {
	_, err := q.db.Exec(ctx, publicServantDeleteSQL, email)

	if err != nil {
		return errors.Wrap(err, "failed to delete public servant")
	}

	return nil
}
