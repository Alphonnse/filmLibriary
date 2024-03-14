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
