-- create user table
CREATE TABLE IF NOT EXISTS `user` (
	`id` BIGINT NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(255),
	`password` VARCHAR(255),
	`email` VARCHAR(255),
	`username` VARCHAR(255),
	`avatar` VARCHAR(255),
	`status` INT,
	`last_login_ip` VARCHAR(255),
	`last_login_time` DATETIME,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- insert default data
INSERT IGNORE INTO `user` (`id`, `name`, `password`, `email`, `username`, `avatar`, `status`, `last_login_ip`, `last_login_time`) VALUES (1, 'dbland', '$2a$10$kX5f2MW1W9kkia0p4WVjq.bJ6nfOAT0TYU0HQ1S.TKr4OSWg2HQNO', NULL, 'dbland', NULL, 1, NULL, NULL);

-- create config table
CREATE TABLE IF NOT EXISTS `config` (
	`id` BIGINT NOT NULL AUTO_INCREMENT,
	`type` VARCHAR(255),
	`name` VARCHAR(255),
	`host` VARCHAR(255),
	`port` VARCHAR(255),
	`username` VARCHAR(255),
	`password` VARCHAR(255),
	`db_file` VARCHAR(255),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- create execution log table
CREATE TABLE IF NOT EXISTS `execution_log` (
	`id` BIGINT NOT NULL AUTO_INCREMENT,
	`source` VARCHAR(255),
	`database` VARCHAR(255),
	`sql` VARCHAR(255),
	`status` INT,
	`cost` INT,
	`created_time` DATETIME,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;