CREATE TABLE IF NOT EXISTS `blog_tag`
(
    `id`         int(10) unsigned not null AUTO_INCREMENT,
    `name`       varchar(100) default '' comment '标签名称',
    `parent_id`  int(10)      default null comment '父标签ID',
    `created_at` datetime     default current_timestamp comment '创建时间',
    `updated_at` datetime     default current_timestamp comment '更新时间',
    `deleted_at` datetime     default '1970-01-01 00:00:00' comment '删除时间',
    primary key (`id`)
) engine = InnoDB
  default charset = utf8 comment ='文章标签';

ALTER TABLE `blog_tag`
    add author int(10) unsigned;
ALTER TABLE `blog_tag`
    add state tinyint not null;

ALTER TABLE `blog_tag`
    CHANGE "author" "created_id" int unsigned not null;