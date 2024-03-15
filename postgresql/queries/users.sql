-- name: AddUser :one
INSERT INTO users (id, role_id, name, password, created_at, updated_at)
VALUES($1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUser :one
SELECT * FROM users WHERE name = $1;
