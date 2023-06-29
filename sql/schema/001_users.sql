-- +goose Up
CREATE TABLE jwtUSER (
    Email VARCHAR(255) NOT NULL UNIQUE,
    Password VARCHAR(255) NOT NULL,
    PRIMARY KEY (Email)
);

-- +goose Down
DROP TABLE jwtUSER;
