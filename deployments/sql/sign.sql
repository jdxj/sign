CREATE TABLE `specification`
(
    `spec_id` bigint(20) NOT NULL AUTO_INCREMENT,
    `spec`    varchar(100) NOT NULL,
    PRIMARY KEY (`spec_id`),
    UNIQUE KEY `specification_spec_IDX` (`spec`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='已注册的 crontab 表达式';

CREATE TABLE `task`
(
    `task_id`     bigint(20) NOT NULL AUTO_INCREMENT,
    `description` varchar(100) NOT NULL DEFAULT '',
    `user_id`     bigint(20) NOT NULL,
    `kind`        varchar(100) NOT NULL COMMENT '任务类型',
    `spec`        varchar(100) NOT NULL,
    `param`       blob COMMENT '任务需要的参数或配置',
    `created_at`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`  timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`task_id`),
    KEY           `task_user_id_IDX` (`user_id`) USING BTREE,
    KEY           `task_spec_IDX` (`spec`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4

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

-- 设计中
CREATE TABLE `contact`
(
    `contact_id` bigint       NOT NULL AUTO_INCREMENT,
    `user_id`    bigint       NOT NULL,
    `type`       varchar(100) NOT NULL COMMENT '联系方式类型',
    `contact`    varchar(100) NOT NULL COMMENT '联系方式取值',
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`contact_id`),
    KEY          `user_id_idx` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 设计中
CREATE TABLE `notice`
(
    `user_id`    bigint NOT NULL,
    `contact_id` bigint NOT NULL,
    PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
