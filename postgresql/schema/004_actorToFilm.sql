-- +goose Up

CREATE TABLE actors_films (
	film_id UUID NOT NULL REFERENCES films(id) ON DELETE CASCADE,
	actor_id UUID NOT NULL REFERENCES actors(id) ON DELETE CASCADE,
	UNIQUE(film_id, actor_id)
);


-- +goose Down

DROP TABLE actors_films;
