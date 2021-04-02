/*
 Navicat Premium Data Transfer

 Source Server         : 开发环境mysql：172.20.12.79
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : 172.20.232.81:60792
 Source Schema         : database

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 31/03/2021 14:55:50
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;


-- ----------------------------
-- Records of tb_menu
-- ----------------------------
INSERT INTO `tb_menu` VALUES (1, '2021-03-24 16:09:32', '2021-03-24 16:09:35', NULL, 1, 0, 'SysManage', 0, '系统管理', 'SysManage', 0, 'sys', 'sys', 'el-icon-s-tools', 0);
INSERT INTO `tb_menu` VALUES (2, '2021-03-26 17:44:53', '2021-03-26 17:44:56', NULL, 1, 1, 'UserManage', 1, '用户管理', 'UserManage', 0, 'users', 'user', 'el-icon-user-solid', 0);
INSERT INTO `tb_menu` VALUES (3, '2021-03-26 17:45:41', '2021-03-26 17:45:47', NULL, 1, 1, 'RoleManage', 1, '角色管理', 'RoleManage', 0, 'roles', 'role', 'el-icon-user', 1);
INSERT INTO `tb_menu` VALUES (4, '2021-03-26 17:46:21', '2021-03-26 17:46:24', NULL, 1, 1, 'MeneManage', 1, '菜单管理', 'MenuManage', 0, 'menus', 'menu', 'el-icon-menu', 2);
INSERT INTO `tb_menu` VALUES (5, '2021-03-30 15:34:26', '2021-03-30 15:34:29', NULL, 1, 4, 'Retrieve', 2, '', NULL, 0, 'menus', '', NULL, 0);

-- ----------------------------
-- Records of tb_role
-- ----------------------------
INSERT INTO `tb_role` VALUES (1, '2021-03-24 16:08:49', '2021-03-24 16:08:51', NULL, 0, 'admin', 'admin');

-- ----------------------------
-- Records of tb_role_menu
-- ----------------------------
INSERT INTO `tb_role_menu` VALUES (1, 1);
INSERT INTO `tb_role_menu` VALUES (1, 2);
INSERT INTO `tb_role_menu` VALUES (1, 3);
INSERT INTO `tb_role_menu` VALUES (1, 4);
INSERT INTO `tb_role_menu` VALUES (1, 5);

-- ----------------------------
-- Records of tb_user
-- ----------------------------
INSERT INTO `tb_user` VALUES (1, '2021-03-24 15:36:46', '2021-03-24 15:36:49', NULL, 0, 'admin', '973a4208b1d6ce49f26aa6f9d078b84dc7c30921b4df40465a79f758887a8a1b', '421826500@qq.com', '18810517357', 1);


-- ----------------------------
-- Records of tb_user_role
-- ----------------------------
INSERT INTO `tb_user_role` VALUES (1, 1);

SET FOREIGN_KEY_CHECKS = 1;
