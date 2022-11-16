package storage

import (
	"context"

	"github.com/khanfromasia/hospital/internal/entity"
	"github.com/pkg/errors"
)

const (
	countryCreateSQL = `--- name: CountryCreate :one
		INSERT INTO public.country (cname, population)
		VALUES ($1, $2)`

	countryGetAllSQL = `--- name: CountryGetAll :many
		SELECT cname, population
		FROM public.country`

	countryUpdateSQL = `--- name: CountyUpdate :one
		UPDATE public.country
		SET population = $1
		WHERE cname = $2`

	countryDeleteSQL = `--- name: CountyDelete :one
		DELETE FROM public.country
		WHERE cname = $1`
)

// CountryCreate creates a new country.
func (q *Queries) CountryCreate(ctx context.Context, arg entity.Country) (entity.Country, error) {
	_, err := q.db.Exec(
		ctx,
		countryCreateSQL,
		arg.Cname,
		arg.Population,
	)

	if err != nil {
		return entity.Country{}, errors.Wrap(err, "failed to create department")
	}

	return arg, nil
}

// CountryGetAll gets all countries.
func (q *Queries) CountryGetAll(ctx context.Context) ([]entity.Country, error) {
	rows, err := q.db.Query(ctx, countryGetAllSQL)

	if err != nil {
		return []entity.Country{}, errors.Wrap(err, "[queries.DepartmentGetAll] failed to get all departments")
	}

	defer rows.Close()

	var countries []entity.Country

	for rows.Next() {
		var country entity.Country

		err = rows.Scan(
			&country.Cname,
			&country.Population,
		)

		if err != nil {
			return []entity.Country{}, errors.Wrap(err, "failed to scan country")
		}

		countries = append(countries, country)
	}

	return countries, nil
}

// CountryUpdate updates a country.
func (q *Queries) CountryUpdate(ctx context.Context, arg entity.Country) (entity.Country, error) {
	_, err := q.db.Exec(
		ctx,
		countryUpdateSQL,
		arg.Population,
		arg.Cname,
	)

	if err != nil {
		return entity.Country{}, errors.Wrap(err, "failed to create country")
	}

	return arg, nil
}

// CountryDelete deletes a country.
func (q *Queries) CountryDelete(ctx context.Context, cname string) error {
	_, err := q.db.Exec(
		ctx,
		countryDeleteSQL,
		cname,
	)

	if err != nil {
		return errors.Wrap(err, "failed to create country")
	}

	return nil
}
