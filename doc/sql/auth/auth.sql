
CREATE TABLE `baker_auth_tree_node`
(
    `id`          int unsigned                           NOT NULL AUTO_INCREMENT COMMENT '节点ID',
    `parent_id`   int                                    NOT NULL DEFAULT '0' COMMENT '父节点ID',
    `name`        varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '节点名称',
    `biz_type`    tinyint                                NOT NULL DEFAULT '0' COMMENT '业务类型',
    `create_user` int unsigned                           NOT NULL DEFAULT '0' COMMENT '创建人id',
    `update_user` int unsigned                           NOT NULL DEFAULT '0' COMMENT '更新人id',
    `create_time` timestamp                              NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp                              NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='树形结构表';

