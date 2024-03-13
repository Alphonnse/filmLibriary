-- +goose Up

CREATE TABLE actors (
	id UUID PRIMARY KEY,
	name TEXT NOT NULL,
	sex VARCHAR(10) NOT NULL,
	birthday TIMESTAMP NOT NULL,
	otherInfo VARCHAR(1000),
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE actors;
