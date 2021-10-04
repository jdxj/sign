CREATE TABLE `secret` (
    `secret_id` bigint(20) NOT NULL AUTO_INCREMENT,
    `describe` varchar(100) NOT NULL DEFAULT '',
    `user_id` bigint(20) NOT NULL,
    `domain` int(11) NOT NULL,
    `key` text NOT NULL,
    PRIMARY KEY (`secret_id`),
    KEY `secret_user_id_IDX` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
