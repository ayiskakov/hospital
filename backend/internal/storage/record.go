package storage

import (
	"context"

	"github.com/khanfromasia/hospital/backend/internal/entity"
)

const (
	recordCreateSQL = `--- name: RecordCreate :one
		INSERT INTO "public.record"
			(email, cname, disease_code, total_deaths, total_patients)
		VALUES
			($1, $2, $3, $4, $5)`

	recordGetAllSQL = `--- name: RecordGetAll :many
		SELECT
			email, cname, disease_code, total_deaths, total_patients
		FROM
			"public.record"`

	recordDeleteSQL = `--- name: RecordDelete :one
		DELETE FROM "public.record"
		WHERE
			email = $1 AND cname = $2 AND disease_code = $3`

	recordUpdateSQL = `--- name: RecordUpdate :one
		UPDATE "public.record"
		SET
			total_deaths = COALESCE($1, total_deaths),
			total_patients = COALESCE($2, total_patients)
		WHERE
			email = $3 AND cname = $4 AND disease_code = $5
		RETURNING email, cname, disease_code, total_deaths, total_patients`
)

// RecordCreate creates a new record.
func (q *Queries) RecordCreate(ctx context.Context, arg entity.Record) (entity.Record, error) {
	_, err := q.db.Exec(
		ctx,
		recordCreateSQL,
		arg.PublicServant.User.Email,
		arg.Country.Cname,
		arg.Disease.DiseaseCode,
		arg.TotalDeaths,
		arg.TotalPatients,
	)

	if err != nil {
		return entity.Record{}, err
	}

	return arg, nil
}

// RecordGetAll gets all records.
func (q *Queries) RecordGetAll(ctx context.Context) ([]entity.Record, error) {
	rows, err := q.db.Query(ctx, recordGetAllSQL)

	if err != nil {
		return []entity.Record{}, err
	}

	defer rows.Close()

	var records []entity.Record

	for rows.Next() {
		var record entity.Record

		err = rows.Scan(
			&record.PublicServant.User.Email,
			&record.Country.Cname,
			&record.Disease.DiseaseCode,
			&record.TotalDeaths,
			&record.TotalPatients,
		)

		if err != nil {
			return []entity.Record{}, err
		}

		records = append(records, record)
	}

	return records, nil
}

// RecordDelete deletes a record.
func (q *Queries) RecordDelete(ctx context.Context, cname, email, diseaseCode string) error {
	_, err := q.db.Exec(
		ctx,
		recordDeleteSQL,
		email,
		cname,
		diseaseCode,
	)

	if err != nil {
		return err
	}

	return nil
}

// RecordUpdate updates a record.
func (q *Queries) RecordUpdate(ctx context.Context, cname, email, diseaseCode string, arg entity.RecordUpdate) (entity.Record, error) {
	row := q.db.QueryRow(
		ctx,
		recordUpdateSQL,
		arg.TotalDeaths,
		arg.TotalPatients,
		email,
		cname,
		diseaseCode,
	)

	var record entity.Record

	err := row.Scan(
		&record.PublicServant.User.Email,
		&record.Country.Cname,
		&record.Disease.DiseaseCode,
		&record.TotalDeaths,
		&record.TotalPatients,
	)

	if err != nil {
		return entity.Record{}, err
	}

	return record, nil
}
