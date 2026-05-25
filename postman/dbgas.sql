/*
 Navicat Premium Dump SQL

 Source Server         : localhost3316
 Source Server Type    : MySQL
 Source Server Version : 100502 (10.5.2-MariaDB)
 Source Host           : localhost:3316
 Source Schema         : dbgas

 Target Server Type    : MySQL
 Target Server Version : 100502 (10.5.2-MariaDB)
 File Encoding         : 65001

 Date: 22/05/2026 02:06:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for gas_log
-- ----------------------------
DROP TABLE IF EXISTS `gas_log`;
CREATE TABLE `gas_log`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `status_gas` tinyint(1) NULL DEFAULT NULL,
  `alarm_status` varchar(10) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `wifi_status` tinyint(1) NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `wa_sent` text CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gas_log
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
