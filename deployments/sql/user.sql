CREATE TABLE `user` (
    `user_id` bigint(20) NOT NULL AUTO_INCREMENT,
    `nickname` varchar(64) NOT NULL,
    `password` varchar(64) NOT NULL,
    `salt` varchar(64) NOT NULL,
    PRIMARY KEY (`user_id`),
    UNIQUE KEY `user_nickname_IDX` (`nickname`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='用户表'
