-- 不删除数据库
CREATE DATABASE IF NOT EXISTS `gpu_sharing_platform`;

USE `gpu_sharing_platform`;

DROP TABLE IF EXISTS `test_instance`;
CREATE TABLE IF NOT EXISTS `test_instance` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(255) NULL,
  `created_at` TIMESTAMP DEFAULT NULL
);