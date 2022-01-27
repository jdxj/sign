CREATE TABLE `secret`
(
    `secret_id` bigint(20) NOT NULL AUTO_INCREMENT,
    `describe`  varchar(100) NOT NULL DEFAULT '',
    `user_id`   bigint(20) NOT NULL,
    `domain`    int(11) NOT NULL,
    `key`       text         NOT NULL,
    PRIMARY KEY (`secret_id`),
    KEY         `secret_user_id_IDX` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `specification`
(
    `spec_id` bigint(20) NOT NULL AUTO_INCREMENT,
    `spec`    varchar(100) NOT NULL,
    PRIMARY KEY (`spec_id`),
    UNIQUE KEY `specification_spec_IDX` (`spec`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='已注册的 crontab 表达式';

CREATE TABLE `task`
(
    `task_id`    bigint       NOT NULL AUTO_INCREMENT,
    `describe`   varchar(100) NOT NULL DEFAULT '',
    `user_id`    bigint       NOT NULL,
    `kind`       varchar(100) NOT NULL COMMENT '任务类型',
    `spec`       varchar(100) NOT NULL,
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `param`      blob COMMENT '任务需要的参数或配置',
    PRIMARY KEY (`task_id`),
    KEY          `task_user_id_IDX` (`user_id`) USING BTREE,
    KEY          `task_spec_IDX` (`spec`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `user`
(
    `user_id`    bigint       NOT NULL AUTO_INCREMENT,
    `nickname`   varchar(64)  NOT NULL,
    `password`   varchar(64)  NOT NULL,
    `salt`       varchar(64)  NOT NULL,
    `mail`       varchar(100) NOT NULL DEFAULT '',
    `telegram`   varchar(100) NOT NULL DEFAULT '' COMMENT 'todo: 字段名可能不准确, 需要修改',
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`user_id`),
    UNIQUE KEY `user_nickname_IDX` (`nickname`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表'

