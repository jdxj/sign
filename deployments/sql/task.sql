CREATE TABLE `task` (
    `task_id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `describe` varchar(100) NOT NULL DEFAULT '',
    `kind` int(11) NOT NULL,
    `spec_id` bigint(20) NOT NULL,
    `secret_id` bigint(20) NOT NULL,
    PRIMARY KEY (`task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
