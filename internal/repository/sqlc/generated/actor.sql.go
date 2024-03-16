// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: actor.sql

package generated

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const addActorInfo = `-- name: AddActorInfo :one
INSERT INTO actors (id, name, sex, birthday, otherInfo, created_at, updated_at)
VALUES($1,$2,$3,$4,$5,$6,$7)
RETURNING id, name, sex, birthday, otherinfo, created_at, updated_at
`

type AddActorInfoParams struct {
	ID        uuid.UUID
	Name      string
	Sex       string
	Birthday  time.Time
	Otherinfo sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) AddActorInfo(ctx context.Context, arg AddActorInfoParams) (Actor, error) {
	row := q.db.QueryRowContext(ctx, addActorInfo,
		arg.ID,
		arg.Name,
		arg.Sex,
		arg.Birthday,
		arg.Otherinfo,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Actor
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Sex,
		&i.Birthday,
		&i.Otherinfo,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const changeActorInfo = `-- name: ChangeActorInfo :one
UPDATE actors 
SET name = $2,
	sex = $3,
	birthday = $4,
	otherInfo = $5,
	updated_at = $6
WHERE id = $1
RETURNING id, name, sex, birthday, otherinfo, created_at, updated_at
`

type ChangeActorInfoParams struct {
	ID        uuid.UUID
	Name      string
	Sex       string
	Birthday  time.Time
	Otherinfo sql.NullString
	UpdatedAt time.Time
}

func (q *Queries) ChangeActorInfo(ctx context.Context, arg ChangeActorInfoParams) (Actor, error) {
	row := q.db.QueryRowContext(ctx, changeActorInfo,
		arg.ID,
		arg.Name,
		arg.Sex,
		arg.Birthday,
		arg.Otherinfo,
		arg.UpdatedAt,
	)
	var i Actor
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Sex,
		&i.Birthday,
		&i.Otherinfo,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getActorByID = `-- name: GetActorByID :one
SELECT id, name, sex, birthday, otherinfo, created_at, updated_at FROM actors WHERE id = $1
`

func (q *Queries) GetActorByID(ctx context.Context, id uuid.UUID) (Actor, error) {
	row := q.db.QueryRowContext(ctx, getActorByID, id)
	var i Actor
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Sex,
		&i.Birthday,
		&i.Otherinfo,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getActorsList = `-- name: GetActorsList :many
SELECT id, name, sex, birthday, otherinfo, created_at, updated_at FROM actors
`

func (q *Queries) GetActorsList(ctx context.Context) ([]Actor, error) {
	rows, err := q.db.QueryContext(ctx, getActorsList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Actor
	for rows.Next() {
		var i Actor
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Sex,
			&i.Birthday,
			&i.Otherinfo,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getActorsListWithFilms = `-- name: GetActorsListWithFilms :many
SELECT a.id AS actor_id, a.name AS actor_name, f.id AS film_id, f.title AS film_title
FROM actors a
JOIN actors_films af ON a.id = af.actor_id
JOIN films f ON af.film_id = f.id
ORDER BY a.id, a.name, f.title
`

type GetActorsListWithFilmsRow struct {
	ActorID   uuid.UUID
	ActorName string
	FilmID    uuid.UUID
	FilmTitle string
}

func (q *Queries) GetActorsListWithFilms(ctx context.Context) ([]GetActorsListWithFilmsRow, error) {
	rows, err := q.db.QueryContext(ctx, getActorsListWithFilms)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetActorsListWithFilmsRow
	for rows.Next() {
		var i GetActorsListWithFilmsRow
		if err := rows.Scan(
			&i.ActorID,
			&i.ActorName,
			&i.FilmID,
			&i.FilmTitle,
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

const rmActorInfo = `-- name: RmActorInfo :exec
DELETE FROM actors WHERE id = $1
`

func (q *Queries) RmActorInfo(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, rmActorInfo, id)
	return err
}
