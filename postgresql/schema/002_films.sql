-- +goose Up

CREATE TABLE films (
	id UUID PRIMARY KEY,
	title VARCHAR(150) NOT NULL CHECK (LENGTH(title) BETWEEN 1 AND 150),
	description VARCHAR(1000),
	releaseDate TIMESTAMP,
	rate INTEGER CHECK (rate BETWEEN 1 AND 10),
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE films;
