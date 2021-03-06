CREATE DATABASE sy_vendor
  DEFAULT CHARACTER SET utf8
  DEFAULT COLLATE utf8_general_ci;
USE sy_vendor;
SET NAMES utf8;

DROP TABLE IF EXISTS user;

CREATE TABLE IF NOT EXISTS `user` 
(
    `user_id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户编号',
    `user_name` VARCHAR(45) NOT NULL COMMENT '用户名称',
    `user_age` TINYINT(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户年龄',
    `user_sex` TINYINT(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户性别',
    PRIMARY KEY (`user_id`)
)
  ENGINE =InnoDB
  DEFAULT CHARACTER SET =utf8
  COLLATE =utf8_unicode_ci
  COMMENT ='用户表';

Insert into user values(1,"sy-vendor",1,1);
