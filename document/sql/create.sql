CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `passport` varchar(45) NOT NULL COMMENT '用户账号',
  `password` varchar(45) NOT NULL COMMENT '用户密码',
  `nickname` varchar(45) NOT NULL COMMENT '用户昵称',
  `create_at` datetime DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `account` (
  `account` varchar(45) NOT NULL COMMENT '用户账号',
  `balance` int(30) NOT NULL COMMENT '用户余额',
  `interest` float(30) NOT NULL COMMENT '用户利息',
  `transaction_time` datetime NOT NULL COMMENT '交易时间',
  `create_at` datetime DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;