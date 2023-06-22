-- +migrate Up
DROP TABLE IF EXISTS `projects`;
CREATE TABLE IF NOT EXISTS `projects` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(128) NOT NULL,
  `Title` varchar(45) DEFAULT NULL,
  `Date` datetime DEFAULT NULL,
  `SAPNumber` varchar(45) DEFAULT NULL,
  `StatusId` int(11) DEFAULT NULL,
  `Notes` varchar(500) DEFAULT NULL,
  `BranchId` int(11) DEFAULT NULL,
  `ServiceId` int(11),
  PRIMARY KEY (`Id`),
   CONSTRAINT projects_relation_1 FOREIGN KEY (ServiceId) REFERENCES services(Id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE projects;
