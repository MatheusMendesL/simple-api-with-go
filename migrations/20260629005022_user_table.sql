-- +goose Up

CREATE TABLE "user"
(
    ID        BIGSERIAL    NOT NULL PRIMARY KEY,
    Firstname VARCHAR(255) NOT NULL,
    LastName  VARCHAR(255) NOT NULL,
    Biography VARCHAR(255) NOT NULL
);

-- +goose Down

DROP TABLE IF EXISTS "user";