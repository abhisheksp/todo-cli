-- +goose Up
CREATE TABLE todo(
    id SERIAL PRIMARY KEY,
    list CHARACTER(10),
    item TEXT
);

-- +goose Down
DROP TABLE todo;
