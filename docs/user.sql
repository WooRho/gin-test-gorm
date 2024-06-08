CREATE TABLE IF NOT EXISTS `user`
(
    `name`        VARCHAR(255)        DEFAULT '' NOT NULL COMMENT '成品编号',
    `password`    VARCHAR(64)         DEFAULT '' NOT NULL COMMENT '密码',
    `phone`       VARCHAR(64)         DEFAULT '' NOT NULL COMMENT '手机号码',
    `id`          bigint(20) unsigned            NOT NULL AUTO_INCREMENT,
    `create_time` datetime                       NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
    `creator_id`  bigint(20) unsigned DEFAULT 0  NOT NULL COMMENT '创建人ID （关联user.id）',
    `update_time` datetime                       NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
    `updater_id`  bigint(20) unsigned DEFAULT 0  NOT NULL COMMENT '更新人ID （关联user.id）',
    `delete_time` datetime                       NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '删除时间',
    `deleter_id`  bigint(20) unsigned DEFAULT 0  NOT NULL COMMENT '删除人ID （关联user.id）',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='用户表';