package storage

import (
	"context"

	"github.com/khanfromasia/hospital/internal/entity"
	"github.com/pkg/errors"
)

const (
	diseaseTypeCreateSQL = `--- name: DiseaseTypeCreate :one
		INSERT INTO 
		    "public.disease_type" 
		    (description) 
		VALUES ($1) 
		RETURNING id`

	diseaseTypeGetAllSQL = `--- name: DiseaseTypeGetAll :many
		SELECT
			id, description
		FROM
			"public.disease_type"`

	diseaseTypeUpdateSQL = `--- name: DiseaseTypeUpdate :one
		UPDATE
			"public.disease_type"
		SET
			description = COALESCE($1, description)
		WHERE
			id = $2
		RETURNING id, description`

	diseaseTypeDeleteSQL = `--- name: DiseaseTypeDelete :one
		DELETE FROM
			"public.disease_type"
		WHERE
			id = $1`

	diseaseCreateSQL = `--- name: DiseaseCreate :one
		INSERT INTO
			"public.disease"
			(id, description, disease_code, pathogen)
		VALUES
			($1, $2, $3, $4)`

	diseaseGetAllSQL = `--- name: DiseaseGetAll :many
		SELECT
			D.id, D.description, D.disease_code, D.pathogen, DT.description
		FROM
			"public.disease" D
		INNER JOIN
			"public.disease_type" DT
		ON
			D.id = DT.id`

	diseaseUpdateSQL = `--- name: DiseaseUpdate :one
		UPDATE
			"public.disease"
		SET
			description = COALESCE($1, description),
			pathogen = COALESCE($3, pathogen),
			id = COALESCE($4, id),
			disease_code = COALESCE($5, disease_code)
		WHERE
			disease_code = $2
		RETURNING
			id, description, disease_code, pathogen`

	diseaseDeleteSQL = `--- name: DiseaseDelete :one
		DELETE FROM
			"public.disease"
		WHERE
			disease_code = $1`
)

// DiseaseTypeCreate creates a new disease type.
func (q *Queries) DiseaseTypeCreate(ctx context.Context, arg entity.DiseaseType) (entity.DiseaseType, error) {
	row := q.db.QueryRow(ctx, diseaseTypeCreateSQL, arg.Description)

	if err := row.Scan(&arg.ID); err != nil {
		return entity.DiseaseType{}, errors.Wrap(err, "could not scan row disease type create")
	}

	return arg, nil
}

// DiseaseTypeGetAll gets all disease types.
func (q *Queries) DiseaseTypeGetAll(ctx context.Context) ([]entity.DiseaseType, error) {
	rows, err := q.db.Query(ctx, diseaseTypeGetAllSQL)

	if err != nil {
		return []entity.DiseaseType{}, err
	}

	defer rows.Close()

	var diseaseTypes []entity.DiseaseType

	for rows.Next() {
		var diseaseType entity.DiseaseType

		err = rows.Scan(
			&diseaseType.ID,
			&diseaseType.Description,
		)

		if err != nil {
			return []entity.DiseaseType{}, errors.Wrap(err, "could not scan row disease type get all")
		}

		diseaseTypes = append(diseaseTypes, diseaseType)
	}

	return diseaseTypes, nil
}

// DiseaseTypeUpdate updates a disease type.
func (q *Queries) DiseaseTypeUpdate(ctx context.Context, id int64, arg entity.DiseaseTypeUpdate) (entity.DiseaseType, error) {
	row := q.db.QueryRow(ctx, diseaseTypeUpdateSQL, arg.Description, id)

	var diseaseType entity.DiseaseType

	if err := row.Scan(&diseaseType.ID, &diseaseType.Description); err != nil {
		return entity.DiseaseType{}, errors.Wrap(err, "could not scan row disease type update")
	}

	return diseaseType, nil
}

// DiseaseTypeDelete deletes a disease type.
func (q *Queries) DiseaseTypeDelete(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, diseaseTypeDeleteSQL, id)

	if err != nil {
		return errors.Wrap(err, "could not delete disease type")
	}

	return nil
}

// DiseaseCreate creates a new disease.
func (q *Queries) DiseaseCreate(ctx context.Context, arg entity.Disease) (entity.Disease, error) {
	_, err := q.db.Exec(ctx, diseaseCreateSQL, arg.DiseaseType.ID, arg.Description, arg.DiseaseCode, arg.Pathogen)

	if err != nil {
		return entity.Disease{}, errors.Wrap(err, "could not create disease")
	}

	return arg, nil
}

// DiseaseGetAll gets all diseases.
func (q *Queries) DiseaseGetAll(ctx context.Context) ([]entity.Disease, error) {
	rows, err := q.db.Query(ctx, diseaseGetAllSQL)

	if err != nil {
		return []entity.Disease{}, errors.Wrap(err, "could not get all diseases")
	}

	defer rows.Close()

	var diseases []entity.Disease

	for rows.Next() {
		var disease entity.Disease

		err = rows.Scan(
			&disease.DiseaseType.ID,
			&disease.Description,
			&disease.DiseaseCode,
			&disease.Pathogen,
			&disease.DiseaseType.Description,
		)

		if err != nil {
			return []entity.Disease{}, errors.Wrap(err, "could not scan row disease get all")
		}

		diseases = append(diseases, disease)
	}

	return diseases, nil
}

// DiseaseUpdate updates a disease.
func (q *Queries) DiseaseUpdate(ctx context.Context, diseaseCode string, arg entity.DiseaseUpdate) (entity.Disease, error) {
	row := q.db.QueryRow(ctx, diseaseUpdateSQL, arg.Description, arg.Pathogen, arg.ID, arg.DiseaseCode, diseaseCode)

	var disease entity.Disease

	if err := row.Scan(&disease.DiseaseType.ID, &disease.Description, &disease.DiseaseCode, &disease.Pathogen); err != nil {
		return entity.Disease{}, errors.Wrap(err, "could not scan row disease update")
	}

	return disease, nil
}

// DiseaseDelete deletes a disease.
func (q *Queries) DiseaseDelete(ctx context.Context, diseaseCode string) error {
	_, err := q.db.Exec(ctx, diseaseDeleteSQL, diseaseCode)

	if err != nil {
		return errors.Wrap(err, "could not delete disease")
	}

	return nil
}
