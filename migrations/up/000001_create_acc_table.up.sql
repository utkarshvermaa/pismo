CREATE TABLE IF NOT EXISTS `accounts` (
    `account_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `document_number` varchar(100) COLLATE utf8mb4_bin NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation TimeStamp',
    `updated_at` datetime ON UPDATE CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Update TimeStamp',
    PRIMARY KEY (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=COMPRESSED COMMENT='accounts table';
