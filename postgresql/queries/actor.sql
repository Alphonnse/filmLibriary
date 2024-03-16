-- name: AddActorInfo :one
INSERT INTO actors (id, name, sex, birthday, otherInfo, created_at, updated_at)
VALUES($1,$2,$3,$4,$5,$6,$7)
RETURNING *;

-- name: ChangeActorInfo :one
UPDATE actors 
SET name = $2,
	sex = $3,
	birthday = $4,
	otherInfo = $5,
	updated_at = $6
WHERE id = $1
RETURNING *;

-- name: RmActorInfo :exec
DELETE FROM actors WHERE id = $1;

-- name: GetActorByID :one
SELECT * FROM actors WHERE id = $1;

-- name: GetActorsList :many
SELECT * FROM actors;

-- name: GetActorsListWithFilms :many
SELECT a.id AS actor_id, a.name AS actor_name, f.id AS film_id, f.title AS film_title
FROM actors a
JOIN actors_films af ON a.id = af.actor_id
JOIN films f ON af.film_id = f.id
ORDER BY a.id, a.name, f.title;
