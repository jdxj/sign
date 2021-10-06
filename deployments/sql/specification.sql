CREATE TABLE `specification` (
    `spec_id` bigint(20) NOT NULL AUTO_INCREMENT,
    `spec` varchar(100) NOT NULL,
    PRIMARY KEY (`spec_id`),
    UNIQUE KEY `specification_spec_IDX` (`spec`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='已注册的 crontab 表达式'
