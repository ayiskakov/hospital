package storage

import (
	"context"

	"github.com/khanfromasia/hospital/internal/entity"
)

const (
	specializeCreateSQL = `--- name: SpecializeCreate :one
		INSERT INTO "public.specialize" 
			(id, email)
		VALUES ($1, $2)`

	specializeGetAllSQL = `--- name: SpecializeGetAll :many
		SELECT
			S.id, S.email, D.degree, U.name, U.surname, U.salary, U.phone, C.cname, C.Population
		FROM
			"public.specialize" S
		INNER JOIN
			"public.doctor" D
		ON
			S.email = D.email
		INNER JOIN
			"public.user" U
		ON
			D.email = U.email
		INNER JOIN
			"public.country" C
		ON
			U.cname = C.cname`

	specializeDeleteSQL = `--- name: SpecializeDelete :one
		DELETE FROM
			"public.specialize"
		WHERE
			email = $1 AND id = $2`
)

// SpecializeCreate creates a new specialize
func (q *Queries) SpecializeCreate(ctx context.Context, arg entity.Specialize) (entity.Specialize, error) {
	_, err := q.db.Exec(
		ctx,
		specializeCreateSQL,
		arg.DiseaseType.ID,
		arg.Doctor.User.Email,
	)

	if err != nil {
		return entity.Specialize{}, err
	}

	return arg, nil
}

// SpecializeGetAll gets all specializes
func (q *Queries) SpecializeGetAll(ctx context.Context) ([]entity.Specialize, error) {
	rows, err := q.db.Query(ctx, specializeGetAllSQL)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var specializes []entity.Specialize

	for rows.Next() {
		var specialize entity.Specialize
		err = rows.Scan(
			&specialize.DiseaseType.ID,
			&specialize.Doctor.User.Email,
			&specialize.Doctor.Degree,
			&specialize.Doctor.User.Name,
			&specialize.Doctor.User.Surname,
			&specialize.Doctor.User.Salary,
			&specialize.Doctor.User.Phone,
			&specialize.Doctor.User.Country.Cname,
			&specialize.Doctor.User.Country.Population,
		)

		if err != nil {
			return nil, err
		}

		specializes = append(specializes, specialize)
	}

	return specializes, nil
}

// SpecializeDelete deletes a specialize
func (q *Queries) SpecializeDelete(ctx context.Context, email, id string) error {
	_, err := q.db.Exec(
		ctx,
		specializeDeleteSQL,
		email,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}
