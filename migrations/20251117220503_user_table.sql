-- +goose Up

CREATE TABLE user(
    ID BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY, 
    Firstname VARCHAR(255) NOT NULL,
    LastName VARCHAR(255) NOT NULL,
    Biography VARCHAR(255) NOT NULL
);


-- +goose Down

DROP TABLE users;
