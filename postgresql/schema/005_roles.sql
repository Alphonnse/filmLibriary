-- +goose Up

CREATE TABLE roles (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL UNIQUE,
	description TEXT NOT NULL UNIQUE
);


-- +goose Down

DROP TABLE roles;
