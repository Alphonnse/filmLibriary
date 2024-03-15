-- +goose Up

CREATE TABLE roles (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL UNIQUE,
	description TEXT NOT NULL UNIQUE
);

INSERT INTO roles(name, description) 
	VALUES 
	('admin', 'can do anything'),
	('Regular user', 'can only get info from api')
);


-- +goose Down

DROP TABLE roles;
