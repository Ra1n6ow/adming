/*
 Navicat Premium Data Transfer

 Source Server         : WSL-MariaDB-10.6.16
 Source Server Type    : MariaDB
 Source Server Version : 100618
 Source Host           : 127.0.0.1:3306
 Source Schema         : adming

 Target Server Type    : MariaDB
 Target Server Version : 100618
 File Encoding         : 65001

 Date: 23/06/2024 23:30:10
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `parent_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '父菜单 ID',
  `type` varchar(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '菜单类型',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由名称',
  `path` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由',
  `component` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '前端组件',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否激活',
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '前端i18n',
  `icon` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单icon',
  `order_no` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `is_show` varchar(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '1' COMMENT '是否显示',
  `permission` varchar(253) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '权限标识',
  `ignore_keepalive` varchar(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '1' COMMENT '是否忽略缓存',
  `is_ext` varchar(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '1' COMMENT '是否外链',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `redirect` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '重定向路由',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_name`(`name`) USING BTREE,
  INDEX `fk_menu_children`(`parent_id`) USING BTREE,
  CONSTRAINT `fk_menu_children` FOREIGN KEY (`parent_id`) REFERENCES `menu` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES (1, NULL, '', 'System', '/system', 'LAYOUT', 0, 'routes.demo.system.moduleName', 'ion:settings-outline', 0, '', '', '', '', '2024-06-23 06:53:48', '2024-06-23 06:53:48', '/system/account');
INSERT INTO `menu` VALUES (2, NULL, '', 'Permission', '/permission', 'LAYOUT', 0, 'routes.demo.permission.permission', 'carbon:user-role', 0, '', '', '', '', '2024-06-23 06:53:48', '2024-06-23 06:53:48', '/permission/front/page');
INSERT INTO `menu` VALUES (3, 1, '', 'AccountManagement', 'account', '/demo/system/account/index', 0, 'routes.demo.system.account', '', 0, '', '', '1', '', '2024-06-23 06:53:48', '2024-06-23 06:53:48', '');
INSERT INTO `menu` VALUES (4, 1, '', 'RoleManagement', 'role', '/demo/system/role/index', 0, 'routes.demo.system.role', '', 0, '', '', '1', '', '2024-06-23 06:53:48', '2024-06-23 06:53:48', '');
INSERT INTO `menu` VALUES (5, 1, '', 'MenuManagement', 'menu', '/demo/system/menu/index', 0, 'routes.demo.system.menu', '', 0, '', '', '1', '', '2024-06-23 06:53:48', '2024-06-23 06:53:48', '');
INSERT INTO `menu` VALUES (6, 1, '', 'changePassword', 'changePassword', '/demo/system/password/index', 0, 'routes.demo.system.password', '', 0, '', '', '1', '', '2024-06-23 06:53:48', '2024-06-23 06:53:48', '');
INSERT INTO `menu` VALUES (7, 2, '', 'PermissionBackDemo', 'back', '', 0, 'routes.demo.permission.back', '', 0, '', '', '', '', '2024-06-23 06:53:48', '2024-06-23 06:53:48', '');
INSERT INTO `menu` VALUES (8, 7, '', 'BackAuthPage', 'page', '/demo/permission/back/index', 0, 'routes.demo.permission.backPage', '', 0, '', '', '', '', '2024-06-23 06:53:48', '2024-06-23 06:53:48', '');
INSERT INTO `menu` VALUES (9, 7, '', 'BackAuthBtn', 'btn', '/demo/permission/back/Btn', 0, 'routes.demo.permission.backBtn', '', 0, '', '', '', '', '2024-06-23 06:53:48', '2024-06-23 06:53:48', '');

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `order_no` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `role_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名',
  `role_value` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色值',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `desc` varchar(253) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '描述',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否激活',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_role_value`(`role_value`) USING BTREE,
  UNIQUE INDEX `idx_role_name`(`role_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (1, 1, '管理员', 'admin', '2024-06-23 06:53:12', '2024-06-23 06:53:12', '', 1);
INSERT INTO `role` VALUES (2, 2, '运维', 'ops', '2024-06-23 06:53:12', '2024-06-23 06:53:12', '', 1);

-- ----------------------------
-- Table structure for role_menu
-- ----------------------------
DROP TABLE IF EXISTS `role_menu`;
CREATE TABLE `role_menu`  (
  `role_id` bigint(20) UNSIGNED NOT NULL,
  `menu_id` bigint(20) UNSIGNED NOT NULL,
  PRIMARY KEY (`role_id`, `menu_id`) USING BTREE,
  INDEX `fk_role_menu_menu`(`menu_id`) USING BTREE,
  CONSTRAINT `fk_role_menu_menu` FOREIGN KEY (`menu_id`) REFERENCES `menu` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_role_menu_role` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role_menu
-- ----------------------------
INSERT INTO `role_menu` VALUES (1, 1);
INSERT INTO `role_menu` VALUES (1, 2);
INSERT INTO `role_menu` VALUES (1, 3);
INSERT INTO `role_menu` VALUES (1, 4);
INSERT INTO `role_menu` VALUES (1, 5);
INSERT INTO `role_menu` VALUES (1, 6);
INSERT INTO `role_menu` VALUES (1, 7);
INSERT INTO `role_menu` VALUES (1, 8);
INSERT INTO `role_menu` VALUES (1, 9);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `role_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `avatar` varchar(253) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `desc` varchar(253) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '描述',
  `home_path` varchar(253) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录默认跳转路由',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否激活',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_username`(`username`) USING HASH,
  INDEX `fk_user_role`(`role_id`) USING BTREE,
  CONSTRAINT `fk_user_role` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 1, 'admin', '$2a$10$zaXttrHl3Rdz4UQPV0M3I.FbNtY.3S7F/UevlWiR5g2T1VsItR5E6', '', '系统管理员', '/system/account', '2024-06-23 06:53:12', '2024-06-23 06:53:12', 0);
INSERT INTO `user` VALUES (2, 2, 'djf', '$2a$10$hXPvO1MSPFx7PZHcVqPYk.P2kEVYOG/Uy6ldc7nK/yI1/eZCJZV/u', '', '杜老二', '/system/account', '2024-06-23 06:53:12', '2024-06-23 06:53:12', 0);

SET FOREIGN_KEY_CHECKS = 1;
