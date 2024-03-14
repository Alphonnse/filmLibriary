-- +goose Up

CREATE TABLE users (
	id UUID PRIMARY KEY,
	role_id INTEGER NOT NULL REFERENCES roles(id) ON DELETE CASCADE DEFAULT 2,
	name TEXT NOT NULL,
	password TEXT NOT NULL UNIQUE,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);


-- +goose Down
DROP TABLE users;
