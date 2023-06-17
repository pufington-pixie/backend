-- +migrate Up
CREATE TABLE services (
    id INT PRIMARY KEY,
    name VARCHAR(255)
);

-- +migrate Down
DROP TABLE services;
