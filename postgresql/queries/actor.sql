-- name: AddActorInfo :one
INSERT INTO actors (id, name, sex, birthday, otherInfo, created_at, updated_at)
VALUES($1,$2,$3,$4,$5,$6,$7)
RETURNING *;
