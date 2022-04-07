DROP DATABASE IF EXISTS test;
CREATE DATABASE IF NOT EXISTS test;
USE test;

CREATE TABLE IF NOT EXISTS `person` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `Name` text,
    `Age` int,
    `Location` text,
    PRIMARY KEY (`id`)
);