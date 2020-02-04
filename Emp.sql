CREATE TABLE Employee(
    `id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `firstName` varchar (55) NOT NULL,
    `lastName` varchar (55) NOT NULL,
    `email` varchar (100) NOT NULL,
    `password` varchar (100) NOT NULL,
    `phone` varchar (10) NOT NULL ,
    `registrationAt` TIMESTAMP NULL DEFAULT NULL,
    `updateAt` TIMESTAMP NULL DEFAULT NULL ,
    UNIQUE KEY `email` (`email`)
);