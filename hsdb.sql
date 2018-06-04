/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50631
Source Host           : localhost:3306
Source Database       : hsdb

Target Server Type    : MYSQL
Target Server Version : 50631
File Encoding         : 65001

Date: 2018-06-04 17:28:22
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for f_card_type
-- ----------------------------
DROP TABLE IF EXISTS `f_card_type`;
CREATE TABLE `f_card_type` (
  `f_id` int(11) NOT NULL AUTO_INCREMENT,
  `f_name` varchar(20) NOT NULL COMMENT '名称',
  `f_describe` varchar(20) NOT NULL COMMENT '描述',
  PRIMARY KEY (`f_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='卡牌类型表';

-- ----------------------------
-- Table structure for t_card_base
-- ----------------------------
DROP TABLE IF EXISTS `t_card_base`;
CREATE TABLE `t_card_base` (
  `f_id` int(11) NOT NULL,
  `f_name` varchar(50) NOT NULL COMMENT '卡牌名称',
  `f_profession_id` int(11) NOT NULL COMMENT '职业id',
  `f_rarity` char(1) NOT NULL COMMENT '稀有度（0:普通，1：稀有，2：史诗，3：传说）',
  `f_is_golden` char(1) NOT NULL COMMENT '是否金色卡牌(0:是，1：否',
  `f_type_id` int(11) NOT NULL COMMENT '卡牌类型id',
  `f_is_normal` char(1) NOT NULL DEFAULT '0' COMMENT '是否普通卡（0：是，1：否',
  `f_package_id` int(11) NOT NULL COMMENT '卡包id',
  `f_expend` int(2) NOT NULL COMMENT '消耗尘数',
  `f_describe` varchar(50) NOT NULL COMMENT '描述',
  PRIMARY KEY (`f_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='卡牌基本信息表';

-- ----------------------------
-- Table structure for t_card_package
-- ----------------------------
DROP TABLE IF EXISTS `t_card_package`;
CREATE TABLE `t_card_package` (
  `f_id` int(11) NOT NULL AUTO_INCREMENT,
  `f_name` varchar(50) DEFAULT NULL COMMENT '卡包名称',
  `f_back_image` varchar(50) DEFAULT NULL COMMENT '卡包图片',
  PRIMARY KEY (`f_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='卡包表';

-- ----------------------------
-- Table structure for t_hero
-- ----------------------------
DROP TABLE IF EXISTS `t_hero`;
CREATE TABLE `t_hero` (
  `f_id` int(11) NOT NULL AUTO_INCREMENT,
  `f_name` varchar(100) NOT NULL COMMENT '英雄名称',
  `f_profession_id` int(11) NOT NULL COMMENT '职业id',
  `f_image_src` varchar(50) NOT NULL COMMENT '英雄图片地址',
  `f_skill_id` int(11) NOT NULL COMMENT '技能id',
  PRIMARY KEY (`f_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='英雄表';

-- ----------------------------
-- Table structure for t_profession
-- ----------------------------
DROP TABLE IF EXISTS `t_profession`;
CREATE TABLE `t_profession` (
  `f_id` int(11) NOT NULL AUTO_INCREMENT,
  `f_name` varchar(20) NOT NULL COMMENT '名称',
  PRIMARY KEY (`f_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='职业表';

-- ----------------------------
-- Table structure for t_skill
-- ----------------------------
DROP TABLE IF EXISTS `t_skill`;
CREATE TABLE `t_skill` (
  `f_id` int(11) NOT NULL AUTO_INCREMENT,
  `f_name` varchar(20) NOT NULL COMMENT '名称',
  `f_image_src` varchar(50) NOT NULL COMMENT '技能图片路径',
  `f_expend` int(2) NOT NULL COMMENT '消耗尘数',
  `f_describe` varchar(100) NOT NULL COMMENT '描述',
  PRIMARY KEY (`f_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='技能表';
