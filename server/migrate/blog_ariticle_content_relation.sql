CREATE TABLE `blog_content_article`
(
    `id`         int(10) not null auto_increment comment '关系表id',
    `content_id` int(10) not null comment '文章id',
    `tag_id`     int(10) not null comment '标签id',
    `created_at` datetime default current_timestamp comment '创建时间',
    `updated_at` datetime default current_timestamp comment '更新时间',
    `deleted_at` datetime default '1970-01-01 00:00:00' comment '删除时间',
    primary key (`id`),
    key content (`content_id`),
    key tag (`tag_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8