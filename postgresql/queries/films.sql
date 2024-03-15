-- name: AddFilmInfo :one
INSERT INTO films (id, title, description, releaseDate, rate, created_at, updated_at)
VALUES($1,$2,$3,$4,$5,$6,$7)
RETURNING *;

-- name: InsertIntoActorToFilme :exec
INSERT INTO actors_films (film_id, actor_id)
VALUES($1, $2);

-- name: ChangeFilmInfo :one
UPDATE films 
SET title = $2,
	description = $3,
	releaseDate = $4,
	rate = $5,
	updated_at = $6
WHERE id = $1
RETURNING *;

-- name: GetFilmByID :one
SELECT * FROM films WHERE id = $1;

-- name: GetActorsByFilmId :many
SELECT * FROM actors_films WHERE film_id = $1;

-- name: RmFilmInfo :exec
DELETE FROM films WHERE id = $1;
