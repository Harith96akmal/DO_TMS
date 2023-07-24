CREATE TABLE IF NOT EXISTS `logsheets` (
  `logsheet_no` varchar(20) DEFAULT NULL,
  `truck_no` longtext DEFAULT NULL,
  `order_no` varchar(20) DEFAULT NULL
  /*`created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`logsheet_no`)
  KEY `idx_logsheets_deleted_at` (`deleted_at`)*/
)