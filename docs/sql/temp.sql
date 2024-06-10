
alter table `user`
MODIFY COLUMN  `name`        VARCHAR(255)        DEFAULT '' NOT NULL COMMENT '成品编号',
MODIFY COLUMN  `password`    VARCHAR(255)         DEFAULT '' NOT NULL COMMENT '密码',
MODIFY COLUMN  `phone`       VARCHAR(255)         DEFAULT '' NOT NULL COMMENT '手机号码';
