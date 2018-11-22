/*
Navicat MySQL Data Transfer

Source Server         : 111
Source Server Version : 50506
Source Host           : localhost:3306
Source Database       : changed

Target Server Type    : MYSQL
Target Server Version : 50506
File Encoding         : 65001

Date: 2018-11-22 17:23:43
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `comments`
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `id` varchar(64) NOT NULL DEFAULT '',
  `video_id` varchar(64) DEFAULT NULL,
  `author_id` int(10) DEFAULT NULL,
  `content` text,
  `time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of comments
-- ----------------------------

-- ----------------------------
-- Table structure for `sessions`
-- ----------------------------
DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions` (
  `session_id` tinytext NOT NULL,
  `TTL` tinytext,
  `login_name` text
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of sessions
-- ----------------------------

-- ----------------------------
-- Table structure for `users`
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `login_name` varchar(64) DEFAULT NULL,
  `pwd` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of users
-- ----------------------------

-- ----------------------------
-- Table structure for `video_info`
-- ----------------------------
DROP TABLE IF EXISTS `video_info`;
CREATE TABLE `video_info` (
  `id` varchar(64) NOT NULL,
  `author_id` int(10) DEFAULT NULL,
  `name` text,
  `display_time` text,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of video_info
-- ----------------------------
