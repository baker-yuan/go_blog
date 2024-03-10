drop table blog_article;
drop table blog_tag;
drop table blog_article_tag;
drop table blog_category;
drop table blog_friend_link;
drop table blog_moment;
drop table blog_page;
drop table blog_config;


CREATE TABLE `blog_article`
(
    `id`                 int unsigned     NOT NULL AUTO_INCREMENT COMMENT '主键',
    `category_id`        int              NOT NULL DEFAULT '0' COMMENT '文章分类表ID|biz_type表主键',
    `title`              varchar(255)     NOT NULL DEFAULT '' COMMENT '文章标题',
    `description`        varchar(200)     NOT NULL DEFAULT '' COMMENT '文章简介，最多200字',
    `content`            mediumtext       NOT NULL COMMENT '文章内容',
    `cover_image`        varchar(255)     NOT NULL DEFAULT '' COMMENT '文章封面图片',
    `original_url`       varchar(255)     NOT NULL DEFAULT '' COMMENT '原文链接',
    `password`           varchar(255)     NOT NULL DEFAULT '' COMMENT '密码保护',
    `words`              int unsigned     NOT NULL DEFAULT '0' COMMENT '文章字数',
    `read_time`          int unsigned     NOT NULL DEFAULT '0' COMMENT '文章阅读时长(分钟)',
    `type`               tinyint unsigned NOT NULL DEFAULT '0' COMMENT '文章类型 1-原创 2-转载 3-翻译',
    `status`             tinyint unsigned NOT NULL DEFAULT '0' COMMENT '文章状态 0-草稿 1-已发布',
    `format`             tinyint unsigned NOT NULL DEFAULT '0' COMMENT '文章格式 1-markdown 2-富文本',
    `visibility`         tinyint unsigned NOT NULL DEFAULT '0' COMMENT '文章可见性 1-公开 2-私密 3-密码保护',
    `is_top`             bit(1)           NOT NULL DEFAULT b'0' COMMENT '置顶开关',
    `is_recommend`       bit(1)           NOT NULL DEFAULT b'0' COMMENT '推荐开关',
    `is_appreciation`    bit(1)           NOT NULL DEFAULT b'0' COMMENT '赞赏开关',
    `is_comment_enabled` bit(1)           NOT NULL DEFAULT b'0' COMMENT '评论开关',
    `is_deleted`         bit(1)           NOT NULL DEFAULT b'0' COMMENT '是否删除',
    `create_time`        timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`        timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`),
    KEY `idx_category_id` (`category_id`),
    KEY `idx_create_time` (`create_time`),
    KEY `idx_update_time` (`update_time`)
) ENGINE = InnoDB COMMENT ='文章表';

CREATE TABLE `blog_tag`
(
    `id`          int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`        varchar(50)  NOT NULL DEFAULT '' COMMENT '标签名',
    `description` varchar(100) NOT NULL DEFAULT '' COMMENT '标签描述',
    `color`       varchar(64)  NOT NULL DEFAULT 'red' COMMENT '标签颜色',
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_name` (`name`)
) ENGINE = InnoDB COMMENT ='文章标签表';

CREATE TABLE `blog_article_tag`
(
    `id`          int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `tag_id`      int unsigned NOT NULL DEFAULT '0' COMMENT '标签表ID|blog_tag主键',
    `article_id`  int unsigned NOT NULL DEFAULT '0' COMMENT '文章ID|biz_article主键',
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_tag_id_article_id` (`tag_id`, `article_id`),
    KEY `idx_article_id` (`article_id`),
    KEY `idx_tag_id` (`tag_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='文章标签关联表';


CREATE TABLE `blog_category`
(
    `id`          int unsigned     NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`        varchar(50)      NOT NULL DEFAULT '' COMMENT '文章类型名',
    `description` varchar(200)     NOT NULL DEFAULT '' COMMENT '类型介绍',
    `sort`        bigint           NOT NULL DEFAULT '0' COMMENT '排序',
    `available`   tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否可用 0-不可用 1-可用',
    `create_time` timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_name` (`name`)
) ENGINE = InnoDB COMMENT ='文章分类表';


CREATE TABLE `blog_friend_link`
(
    `id`           int unsigned     NOT NULL AUTO_INCREMENT COMMENT '主键',
    `link_name`    varchar(20)      NOT NULL DEFAULT '' COMMENT '链接名',
    `link_avatar`  varchar(255)     NOT NULL DEFAULT '' COMMENT '链接头像',
    `link_address` varchar(50)      NOT NULL DEFAULT '' COMMENT '链接地址',
    `link_intro`   varchar(100)     NOT NULL DEFAULT '' COMMENT '链接介绍',
    `status`       tinyint unsigned NOT NULL DEFAULT '1' COMMENT '友链状态 1-已发布 2-以下线',
    `sort`         int unsigned     NOT NULL DEFAULT '0' COMMENT '友链排序',
    `is_deleted`   bit(1)           NOT NULL DEFAULT b'0' COMMENT '是否删除',
    `create_time`  timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT ='友链表';

CREATE TABLE `blog_moment`
(
    `id`           int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `content`      longtext     NOT NULL COMMENT '动态内容',
    `likes`        int unsigned NOT NULL DEFAULT '0' COMMENT '点赞数量',
    `is_published` bit(1)       NOT NULL DEFAULT b'0' COMMENT '是否公开',
    `is_deleted`   bit(1)       NOT NULL DEFAULT b'0' COMMENT '是否删除',
    `create_time`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
    COMMENT ='动态表';

CREATE TABLE `blog_page`
(
    `id`          int unsigned NOT NULL AUTO_INCREMENT COMMENT '页面id',
    `page_name`   varchar(10)  NOT NULL DEFAULT '' COMMENT '页面名',
    `page_label`  varchar(20)  NOT NULL DEFAULT '' COMMENT '页面标签',
    `page_cover`  varchar(255) NOT NULL DEFAULT '' COMMENT '页面封面',
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT ='页面表';

CREATE TABLE `blog_config`
(
    `id`          int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `code`        varchar(20)  NOT NULL DEFAULT '' COMMENT '唯一编码',
    `desc`        varchar(20)  NOT NULL DEFAULT '' COMMENT '描述信息',
    `config`      mediumtext   NOT NULL COMMENT '配置信息',
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_code` (`code`)
) ENGINE = InnoDB COMMENT ='网站配置';