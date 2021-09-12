CREATE TABLE `user` (
    `user_id` bigint(20) NOT NULL AUTO_INCREMENT,
    `nickname` varchar(64) NOT NULL COMMENT '昵称',
    `password` varchar(64) NOT NULL COMMENT '密码',
    PRIMARY KEY (`user_id`),
    UNIQUE KEY `user_nickname_IDX` (`nickname`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表'
