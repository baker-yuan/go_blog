```sql
CREATE TABLE `interaction_praise`
(
    `id`          int(20) unsigned    NOT NULL AUTO_INCREMENT COMMENT '点赞ID',
    `module_code` varchar(50)         NOT NULL DEFAULT '' COMMENT '模块标识',
    `uid`         int(10) unsigned    NOT NULL DEFAULT '0' COMMENT '用户ID',
    `object_id`   int(10) unsigned    NOT NULL DEFAULT '0' COMMENT '信息ID',
    `status`      tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '点赞状态 0-未点赞 1-已点赞',
    `update_time` int(10) unsigned    NOT NULL DEFAULT '0' COMMENT '修改时间',
    `create_time` int(10) unsigned    NOT NULL DEFAULT '0' COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY `idx_object` (`object_id`) USING BTREE,
    KEY `idx_createTime` (`create_time`),
    KEY `idx_update_time` (`update_time`),
    UNIQUE KEY `uk_uid_object_id_module_code` (`uid`, `object_id`, `module_code`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户点赞明细表|baker.yuan|2022-04-05'
```

```sql
insert into tb_praise (module_code, uid, object_id)
values ('article', 1, 1),('article', 2, 1);
```





```bash
curl --location --request POST 'http://127.0.0.1:8850/api/praise' \
--header 'Content-Type: application/json' \
--data-raw '{
    "uid": 1,
    "objectId": 1,
    "moduleCode": "article"
}'
```



```bash
curl --location --request DELETE 'http://127.0.0.1:8850/api/praise' \
--header 'Content-Type: application/json' \
--data-raw '{
    "uid": 1,
    "objectId": 1,
    "moduleCode": "article"
}'
```



