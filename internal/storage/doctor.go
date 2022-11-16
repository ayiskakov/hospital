package storage

import (
	"context"

	"github.com/khanfromasia/hospital/internal/entity"
	"github.com/pkg/errors"
)

const (
	doctorCreateSQL = `--- name: DoctorCreate :one
		INSERT INTO
			public.doctor (email, degree)
		VALUES
			($1, $2)`

	doctorGetAllSQL = `--- name: DoctorGetAll :many
		SELECT
			D.email, D.degree, U.name, U.surname, U.salary, U.phone, C.cname, C.Population
		FROM
			public.doctor D
		INNER JOIN
			public.user U
		ON
			D.email = U.email
		INNER JOIN
			public.country C
		ON
			U.cname = C.cname`

	doctorUpdateSQL = `--- name: DoctorUpdate :one
		UPDATE
			public.doctor
		SET
			degree = COALESCE($1, degree)
		WHERE
			email = $2
		RETURNING degree`

	doctorDeleteSQL = `--- name: DoctorDelete :one
		DELETE FROM
			public.doctor
		WHERE
			email = $1`
)

// DoctorCreate creates a new doctor
func (q *Queries) DoctorCreate(ctx context.Context, arg entity.Doctor) (entity.Doctor, error) {
	_, err := q.db.Exec(
		ctx,
		doctorCreateSQL,
		arg.User.Email,
		arg.Degree,
	)

	if err != nil {
		return entity.Doctor{}, errors.Wrap(err, "failed to create doctor")
	}

	return arg, nil
}

// DoctorGetAll returns all doctors
func (q *Queries) DoctorGetAll(ctx context.Context) ([]entity.Doctor, error) {
	rows, err := q.db.Query(ctx, doctorGetAllSQL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all doctors")
	}
	defer rows.Close()

	var doctors []entity.Doctor

	for rows.Next() {
		var d entity.Doctor
		if err = rows.Scan(
			&d.User.Email,
			&d.Degree,
			&d.User.Name,
			&d.User.Surname,
			&d.User.Salary,
			&d.User.Phone,
			&d.User.Country.Cname,
			&d.User.Country.Population,
		); err != nil {
			return nil, errors.Wrap(err, "failed to scan doctor")
		}
		doctors = append(doctors, d)
	}

	return doctors, nil
}

// DoctorUpdate updates a doctor
func (q *Queries) DoctorUpdate(ctx context.Context, email string, arg entity.DoctorUpdate) (entity.Doctor, error) {
	row := q.db.QueryRow(
		ctx,
		doctorUpdateSQL,
		arg.Degree,
		email,
	)

	var d entity.Doctor

	if err := row.Scan(
		&d.Degree,
	); err != nil {
		return entity.Doctor{}, errors.Wrap(err, "failed to update doctor")
	}

	return d, nil
}

// DoctorDelete deletes a doctor
func (q *Queries) DoctorDelete(ctx context.Context, email string) error {
	_, err := q.db.Exec(ctx, doctorDeleteSQL, email)

	if err != nil {
		return errors.Wrap(err, "failed to delete doctor")
	}
	
	return nil
}
