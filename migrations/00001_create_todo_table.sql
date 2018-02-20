-- +goose Up
CREATE TABLE todo(
    id CHARACTER(10) PRIMARY KEY,
    item TEXT
);

-- +goose Down
DROP TABLE todo;
