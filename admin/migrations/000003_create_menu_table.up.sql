CREATE TABLE IF NOT EXISTS `menu` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `rule_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `status` tinyint(1) NOT NULL DEFAULT 1,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `name` varchar(100) NOT NULL,
  `desc` varchar(100) DEFAULT NULL,
  `frontpath` varchar(255) DEFAULT NULL,
  `front_file_path` varchar(255) DEFAULT NULL,
  `condition` varchar(255) DEFAULT NULL,
  `menu` tinyint(1) NOT NULL DEFAULT 1,
  `order` int(11) NOT NULL DEFAULT 0,
  `icon` varchar(100) DEFAULT NULL,
  `method` varchar(10) DEFAULT 'GET',
  PRIMARY KEY (`id`),
  KEY `idx_menu_rule_id` (`rule_id`),
  KEY `idx_menu_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci; 