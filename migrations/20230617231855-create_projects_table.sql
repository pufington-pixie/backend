-- +migrate Up
CREATE TABLE projects (
    id INT AUTO_INCREMENT PRIMARY KEY,
    projectid INT DEFAULT 0,
    title VARCHAR(255),
    date DATE,
    sap_number VARCHAR(255),
    notes VARCHAR(255),
    branchId INT,
    statusId INT,
    serviceId INT,
    CONSTRAINT projects_relation_1 FOREIGN KEY (serviceId) REFERENCES services(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE projects;
