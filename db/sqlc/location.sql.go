// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: location.sql

package db

import (
	"context"
)

const createLocation = `-- name: CreateLocation :one
INSERT INTO location (
    city,
    state,
    zipcode
) VALUES (
             $1, $2, $3
         ) RETURNING id, city, state, zipcode
`

type CreateLocationParams struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Zipcode string `json:"zipcode"`
}

func (q *Queries) CreateLocation(ctx context.Context, arg CreateLocationParams) (Location, error) {
	row := q.db.QueryRowContext(ctx, createLocation, arg.City, arg.State, arg.Zipcode)
	var i Location
	err := row.Scan(
		&i.ID,
		&i.City,
		&i.State,
		&i.Zipcode,
	)
	return i, err
}

const getAllLocations = `-- name: GetAllLocations :many
SELECT id, city, state, zipcode FROM location
`

func (q *Queries) GetAllLocations(ctx context.Context) ([]Location, error) {
	rows, err := q.db.QueryContext(ctx, getAllLocations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Location{}
	for rows.Next() {
		var i Location
		if err := rows.Scan(
			&i.ID,
			&i.City,
			&i.State,
			&i.Zipcode,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLocation = `-- name: GetLocation :one
SELECT id, city, state, zipcode FROM location
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetLocation(ctx context.Context, id int64) (Location, error) {
	row := q.db.QueryRowContext(ctx, getLocation, id)
	var i Location
	err := row.Scan(
		&i.ID,
		&i.City,
		&i.State,
		&i.Zipcode,
	)
	return i, err
}