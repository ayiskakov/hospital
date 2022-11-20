package storage

import (
	"context"

	"github.com/khanfromasia/hospital/backend/internal/entity"
	"github.com/pkg/errors"
)

const (
	discoveryCreateSQL = `--- name: DiscoveryCreate :one
		INSERT INTO "public.discover" 
			(cname, disease_code, first_enc_date)
		VALUES ($1, $2, $3)`

	discoveryGetAllSQL = `--- name: DiscoveryGetAll :many
		SELECT
			C.cname, C.population, DS.disease_code, D.description, D.pathogen, DS.first_enc_date, D.id, DT.description
		FROM
			"public.discover" DS
		INNER JOIN
			"public.country" C
		ON
			DS.cname = C.cname
		INNER JOIN
			"public.disease" D
		ON
			DS.disease_code = D.disease_code
		INNER JOIN
			"public.disease_type" DT
		ON
			D.id = DT.id`

	discoveryUpdateSQL = `--- name: DiscoveryUpdate :one
		UPDATE
			"public.discover"
		SET
			cname = COALESCE($1, cname),
			disease_code = COALESCE($2, disease_code),
			first_enc_date = COALESCE($3, first_enc_date)
		WHERE
			cname = $4 AND disease_code = $5
		RETURNING cname, disease_code, first_enc_date`

	discoveryDeleteSQL = `--- name: DiscoveryDelete :exec
		DELETE FROM "public.discover"
		WHERE		
			cname = $1 AND disease_code = $2`
)

// DiscoveryCreate creates a new discovery.
func (q *Queries) DiscoveryCreate(ctx context.Context, arg entity.Discovery) (entity.Discovery, error) {
	_, err := q.db.Exec(
		ctx,
		discoveryCreateSQL,
		arg.Country.Cname,
		arg.Disease.DiseaseCode,
		arg.FirstEncDate,
	)

	if err != nil {
		return entity.Discovery{}, errors.Wrap(err, "creating discovery")
	}

	return arg, nil
}

// DiscoveryGetAll gets all discoveries.
func (q *Queries) DiscoveryGetAll(ctx context.Context) ([]entity.Discovery, error) {
	rows, err := q.db.Query(ctx, discoveryGetAllSQL)

	if err != nil {
		return nil, errors.Wrap(err, "getting all discoveries")
	}

	defer rows.Close()

	var discoveries []entity.Discovery

	for rows.Next() {
		var discovery entity.Discovery

		if err = rows.Scan(
			&discovery.Country.Cname,
			&discovery.Country.Population,
			&discovery.Disease.DiseaseCode,
			&discovery.Disease.Description,
			&discovery.Disease.Pathogen,
			&discovery.FirstEncDate,
			&discovery.Disease.DiseaseType.ID,
			&discovery.Disease.DiseaseType.Description,
		); err != nil {
			return nil, errors.Wrap(err, "scanning discovery")
		}

		discoveries = append(discoveries, discovery)
	}

	return discoveries, nil
}

func (q *Queries) DiscoveryUpdate(ctx context.Context, cname, diseaseCode string, arg entity.DiscoveryUpdate) (entity.Discovery, error) {
	row := q.db.QueryRow(ctx, discoveryUpdateSQL, arg.Cname, arg.DiseaseCode, arg.FirstEncDate, cname, diseaseCode)

	var discovery entity.Discovery

	if err := row.Scan(
		&discovery.Country.Cname,
		&discovery.Disease.DiseaseCode,
		&discovery.FirstEncDate,
	); err != nil {
		return entity.Discovery{}, errors.Wrap(err, "updating discovery")
	}

	return discovery, nil
}

func (q *Queries) DiscoveryDelete(ctx context.Context, cname, diseaseCode string) error {
	_, err := q.db.Exec(ctx, discoveryDeleteSQL, cname, diseaseCode)

	if err != nil {
		return errors.Wrap(err, "deleting discovery")
	}

	return nil
}
