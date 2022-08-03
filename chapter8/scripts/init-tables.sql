CREATE TABLE `employees` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `full_name` varchar(500) NOT NULL,
    `position` int(11) NOT NULL,
    `salary` decimal(13,4) NOT NULL,
    `joined` datetime NOT NULL,
    `on_probation` bit(1) NOT NULL,
    `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='';