create table `user` (
	`id` bigint(64) not null primary key AUTO_INCREMENT,
	`name` varchar(24) not null,
	`age` int(11) unsigned not null,
	`create_time` DATETIME not null,
	`update_time` DATETIME not null
) ENGINE=InnoDB DEFAULT CHARSET=utf8;