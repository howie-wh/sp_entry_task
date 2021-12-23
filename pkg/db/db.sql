/*
 Data Transfer
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_tab
-- ----------------------------
DROP TABLE IF EXISTS `user_tab`;
CREATE TABLE `user_tab` (
    `user_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_name` varchar(128) NOT NULL COMMENT 'username',
    `password` varchar(128) NOT NULL COMMENT 'password',
    `nick_name` varchar(128) NOT NULL DEFAULT '' COMMENT 'nickname',
    `email` varchar(128) NOT NULL DEFAULT '' COMMENT 'email',
    `avatar` varchar(128) NOT NULL DEFAULT '' COMMENT 'avatar',
    `del_flag` char(1) NOT NULL DEFAULT '0' COMMENT 'del flag（0-normal 1-delete)',
    `create_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'create time',
    `update_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'update time',
    PRIMARY KEY (`user_id`),
    UNIQUE KEY `uniq_username` (`user_name`) USING BTREE COMMENT 'uniq_username',
    KEY `idx_ctime` (`create_time`) USING BTREE COMMENT 'idx_ctime'
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='user table';

-- ----------------------------
-- Records of user_tab
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_user_tab
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_tab`;
CREATE TABLE `sys_user_tab` (
    `user_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'sys user id',
    `user_name` varchar(128) NOT NULL COMMENT 'username',
    `password` varchar(128) NOT NULL COMMENT 'password',
    `del_flag` char(1) NOT NULL DEFAULT '0' COMMENT 'del flag（0-normal 1-delete)',
    `create_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'create time',
    `update_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'update time',
    PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='system user table';

-- ----------------------------
-- Records of sys_user_tab
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_tab` VALUES ('11','admin','$2a$10$.J/16dp2rYJdQ0A0cYA0P.2UiyIIbwemvuJ52pTetVqYmH6cHURbe','0', '1633774561', '1633774561');
COMMIT;

-- ----------------------------
-- Table structure for activity_tab
-- ----------------------------
DROP TABLE IF EXISTS `activity_tab`;
CREATE TABLE `activity_tab` (
    `activity_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `type_id` bigint unsigned NOT NULL COMMENT 'type id',
    `title` varchar(128) NOT NULL DEFAULT '' COMMENT 'title',
    `location` varchar(128) NOT NULL DEFAULT '' COMMENT 'activity location',
    `content` varchar(1024) NOT NULL DEFAULT '' COMMENT 'activity content',
    `start_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'start time',
    `end_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'end time',
    `del_flag` char(1) NOT NULL DEFAULT '0' COMMENT 'del flag（0-normal 1-delete)',
    `create_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'create time',
    `update_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'update time',
    PRIMARY KEY (`activity_id`),
    UNIQUE KEY `uniq_title` (`title`) USING BTREE COMMENT 'uniq_title',
    KEY `idx_tid_stime` (`type_id`,`start_time`) USING BTREE COMMENT 'idx_tid_stime',
    KEY `idx_stime` (`start_time`) USING BTREE COMMENT 'idx_etime'
    KEY `idx_etime` (`end_time`) USING BTREE COMMENT 'idx_etime'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='activity table';

-- ----------------------------
-- Records of activity_tab
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for activity_type_tab
-- ----------------------------
DROP TABLE IF EXISTS `activity_type_tab`;
CREATE TABLE `activity_type_tab` (
    `type_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `type_name` varchar(64) NOT NULL DEFAULT '' COMMENT 'name',
    `del_flag` char(1) NOT NULL DEFAULT '0' COMMENT 'del flag（0-normal 1-delete)',
    `create_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'create time',
    `update_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'update time',
    PRIMARY KEY (`type_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='activity type table';

-- ----------------------------
-- Records of activity_type_tab
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for activity_join_tab
-- ----------------------------
DROP TABLE IF EXISTS `activity_join_tab`;
CREATE TABLE `activity_join_tab` (
    `join_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'user id',
    `activity_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'activity id',
    `del_flag` char(1) NOT NULL DEFAULT '0' COMMENT 'del flag（0-normal 1-delete)',
    `create_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'create time',
    `update_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'update time',
    PRIMARY KEY (`join_id`),
    KEY `idx_uid_aid` (`user_id`,`activity_id`) USING BTREE COMMENT 'idx_uid_aid',
    UNIQUE KEY `uniq_aid_uid` (`activity_id`,`user_id`) USING BTREE COMMENT 'uniq_aid_uid'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='activity type table';

-- ----------------------------
-- Records of activity_join_tab
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for comment_tab
-- ----------------------------
DROP TABLE IF EXISTS `comment_tab`;
CREATE TABLE `comment_tab` (
    `comment_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'user id',
    `activity_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'activity id',
    `content` varchar(255) NOT NULL DEFAULT '' COMMENT 'comment content',
    `del_flag` char(1) NOT NULL DEFAULT '0' COMMENT 'del flag（0-normal 1-delete)',
    `create_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'create time',
    `update_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'update time',
    PRIMARY KEY (`comment_id`),
    UNIQUE KEY `uniq_aid_uid` (`activity_id`,`user_id`) USING BTREE COMMENT 'uniq_aid_uid'
    KEY `idx_ctime` (`create_time`) USING BTREE COMMENT 'idx_ctime'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='comment table';

-- ----------------------------
-- Records of comment_tab
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;