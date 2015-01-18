SET NAMES utf8;

CREATE TABLE `user` (
      `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
      `account` varchar(30) NOT NULL,
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
      `desc` text NOT NULL,
      `type` varchar(10) NOT NULL,
      `createdBy` int(11) NOT NULL,
      `createdAt` datetime NOT NULL,
      PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE TABLE `article` (
      `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
      `title` varchar(255) NOT NULL,
      `content` text NOT NULL,
      `order` smallint NOT NULL,
      `createdAt` datetime NOT NULL,
      `lastEdit` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
