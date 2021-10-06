CREATE TABLE `task` (
    `task_id` bigint(20) NOT NULL AUTO_INCREMENT,
    `describe` varchar(100) NOT NULL DEFAULT '',
    `user_id` bigint(20) NOT NULL,
    `secret_id` bigint(20) NOT NULL,
    `kind` int(11) NOT NULL,
    `spec` varchar(100) NOT NULL,
    PRIMARY KEY (`task_id`),
    KEY `task_user_id_IDX` (`user_id`) USING BTREE,
    KEY `task_spec_IDX` (`spec`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
