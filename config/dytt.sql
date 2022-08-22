CREATE DATABASE  IF NOT EXISTS `dytt`;
use `dytt`;

DROP TABLE IF EXISTS `comment`;
DROP TABLE IF EXISTS `relation`;
DROP TABLE IF EXISTS `user`;
DROP TABLE IF EXISTS `user_favorite_videos`;
DROP TABLE IF EXISTS `video`;
CREATE TABLE `comment`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL,
  `video_id` bigint UNSIGNED NOT NULL,
  `user_id` bigint UNSIGNED NOT NULL,
  `content` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `deleted_at`(`deleted_at`),
  INDEX `video_id`(`video_id`),
  INDEX `user_id`(`user_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8;

CREATE TABLE `relation`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL,
  `user_id` bigint UNSIGNED NOT NULL,
  `to_user_id` bigint UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `deleted_at`(`deleted_at`),
  INDEX `user_id`(`user_id`),
  INDEX `to_user_id`(`to_user_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8;

CREATE TABLE `user`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL,
  `user_name` varchar(40) NOT NULL,
  `password` varchar(255) NOT NULL,
  `following_count` bigint NULL,
  `follower_count` bigint NULL,
  PRIMARY KEY (`id`),
  INDEX `deleted_at`(`deleted_at`),
  INDEX `user_name`(`user_name`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8;

CREATE TABLE `user_favorite_videos`  (
  `user_id` bigint UNSIGNED NOT NULL,
  `video_id` bigint UNSIGNED NOT NULL,
  PRIMARY KEY (`user_id`, `video_id`)
) ENGINE = InnoDB CHARACTER SET = utf8;

CREATE TABLE `video`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL,
  `author_id` bigint UNSIGNED NOT NULL,
  `play_url` varchar(255) NOT NULL,
  `cover_url` varchar(255) NOT NULL,
  `favorite_count` bigint NULL DEFAULT NULL,
  `comment_count` bigint NULL DEFAULT NULL,
  `title` varchar(255) NULL,
  PRIMARY KEY (`id`),
  INDEX `deleted_at`(`deleted_at`),
  INDEX `author_id`(`author_id`),
  INDEX `updated_at`(`updated_at`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8;

ALTER TABLE `comment` ADD CONSTRAINT `comment_video_id` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`);
ALTER TABLE `comment` ADD CONSTRAINT `comment_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
ALTER TABLE `relation` ADD CONSTRAINT `relation_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
ALTER TABLE `relation` ADD CONSTRAINT `relation_to_user_id` FOREIGN KEY (`to_user_id`) REFERENCES `user` (`id`);
ALTER TABLE `user_favorite_videos` ADD CONSTRAINT `ufv_video_id` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`);
ALTER TABLE `user_favorite_videos` ADD CONSTRAINT `ufv_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
ALTER TABLE `video` ADD CONSTRAINT `video_author_id` FOREIGN KEY (`author_id`) REFERENCES `user` (`id`);