SET NAMES utf8;

CREATE TABLE `user` (
      `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
      `name` varchar(30) NOT NULL,
      `password` varchar(255) NOT NULL,
      `email` varchar(60) NOT NULL,
      `avatar` varchar(255) NOT NULL,
      `createdat` datetime NOT NULL,
      `lastlogin` datetime NOT NULL,
      PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE TABLE `book` (
      `id` int(11) NOT NULL AUTO_INCREMENT,
      `name` varchar(100) NOT NULL,
      `desc` int(11) NOT NULL,
      `type` varchar(10) NOT NULL,
      `createdBy` int(11) NOT NULL,
      `createdAt` datetime NOT NULL,
      PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
