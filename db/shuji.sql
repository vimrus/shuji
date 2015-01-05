SET NAMES utf8;

DROP TABLE IF EXISTS `user`;
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
