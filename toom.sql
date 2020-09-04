/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : toom

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 04/09/2020 17:44:51
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for toomhub_square
-- ----------------------------
DROP TABLE IF EXISTS `toomhub_square`;
CREATE TABLE `toomhub_square`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
  `created_by` int(11) UNSIGNED NOT NULL COMMENT '创建时间',
  `created_at` int(11) UNSIGNED NOT NULL COMMENT '创建人',
  `likes_count` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '点赞数',
  `argument_count` int(11) UNSIGNED NOT NULL COMMENT '评论数',
  `collect_count` int(11) UNSIGNED NOT NULL COMMENT '收藏数',
  `share_count` int(11) UNSIGNED NOT NULL COMMENT '分享数',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for toomhub_square_image
-- ----------------------------
DROP TABLE IF EXISTS `toomhub_square_image`;
CREATE TABLE `toomhub_square_image`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `image_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '图片地址',
  `image_extension` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '图片后缀',
  `image_size` int(11) NOT NULL COMMENT '图片大小',
  `square_id` int(11) NOT NULL COMMENT '关联的square',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for toomhub_user_mini
-- ----------------------------
DROP TABLE IF EXISTS `toomhub_user_mini`;
CREATE TABLE `toomhub_user_mini`  (
  `mini_id` int(11) NOT NULL AUTO_INCREMENT,
  `open_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` int(11) NOT NULL,
  PRIMARY KEY (`mini_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 123155 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for toomhub_user_mini_profile
-- ----------------------------
DROP TABLE IF EXISTS `toomhub_user_mini_profile`;
CREATE TABLE `toomhub_user_mini_profile`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `mini_id` int(11) UNSIGNED NOT NULL COMMENT '用户关联id',
  `nick_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `gender` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '性别',
  `city` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '城市',
  `province` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '省',
  `country` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '国家',
  `avatar_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '头像',
  `fans_count` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '粉丝数',
  `likes_count` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '获赞数',
  `follow_count` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '关注数',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for toomhub_user_mini_token
-- ----------------------------
DROP TABLE IF EXISTS `toomhub_user_mini_token`;
CREATE TABLE `toomhub_user_mini_token`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `mini_id` int(11) UNSIGNED NOT NULL COMMENT '关联id',
  `access_token` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'access_token',
  `refresh_token` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '刷新access_token用的token',
  `created_at` int(11) UNSIGNED NOT NULL COMMENT '创建时间',
  `updated_at` int(11) UNSIGNED NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
