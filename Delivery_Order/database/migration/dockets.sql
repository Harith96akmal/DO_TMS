CREATE TABLE IF NOT EXISTS `dockets` (
  `order_no` varchar(20) DEFAULT NULL,
  `customer` longtext DEFAULT NULL,
  `pick_up_point` longtext DEFAULT NULL,
  `delivery_point` longtext DEFAULT NULL,
  `quantity` double DEFAULT NULL,
  `volume` double DEFAULT NULL,
  `status` longtext DEFAULT NULL,
  `truck_no` longtext DEFAULT NULL,
  `logsheet_no` longtext DEFAULT NULL,
  /*`created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,*/
  PRIMARY KEY (`order_no`)
  /*KEY `idx_dockets_deleted_at` (`deleted_at`)*/
)