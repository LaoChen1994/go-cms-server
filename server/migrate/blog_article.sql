CREATE TABLE if not exists `blog_article`
(
    `id`         int(10) unsigned NOT NULL AUTO_INCREMENT,
    `title`      varchar(100)        DEFAULT '' COMMENT '文章标题',
    `desc`       varchar(255)        DEFAULT '' COMMENT '简述',
    `content`    text,
    `created_at` datetime            DEFAULT current_timestamp COMMENT '创建时间',
    `author`     varchar(100)        DEFAULT '' COMMENT '创建人',
    `updated_at` datetime            DEFAULT current_timestamp COMMENT '修改时间',
    `deleted_at` datetime            DEFAULT '1970-01-01 00:00:00' comment '删除时间',
    `state`      tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='文章管理';