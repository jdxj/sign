SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for secret
-- ----------------------------
DROP TABLE IF EXISTS `secret`;
CREATE TABLE `secret` (
  `secret_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `describe` varchar(100) NOT NULL DEFAULT '',
  `user_id` bigint(20) NOT NULL,
  `domain` int(11) NOT NULL,
  `key` text NOT NULL,
  PRIMARY KEY (`secret_id`),
  KEY `secret_user_id_IDX` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for specification
-- ----------------------------
DROP TABLE IF EXISTS `specification`;
CREATE TABLE `specification` (
  `spec_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `spec` varchar(100) NOT NULL,
  PRIMARY KEY (`spec_id`),
  UNIQUE KEY `specification_spec_IDX` (`spec`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='已注册的 crontab 表达式';

-- ----------------------------
-- Table structure for task
-- ----------------------------
DROP TABLE IF EXISTS `task`;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nickname` varchar(64) NOT NULL,
  `password` varchar(64) NOT NULL,
  `salt` varchar(64) NOT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `user_nickname_IDX` (`nickname`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

SET FOREIGN_KEY_CHECKS = 1;
