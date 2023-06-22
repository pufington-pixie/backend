-- +migrate Up
DROP TABLE IF EXISTS `services`;
CREATE TABLE services (
    Id INT PRIMARY KEY,
    Name VARCHAR(255)
);

-- +migrate Down
DROP TABLE services;
