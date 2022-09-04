create table if not exists `blog_user`
(
    `id`         int(10)      not null auto_increment,
    `username`   varchar(255) not null comment '用户名',
    `mobile`     varchar(11) default '' comment '手机号',
    `password`   varchar(30)  not null comment '密码',
    `deleted_at` datetime    default '1970-01-01 00:00:00' comment '删除时间',
    `created_at` datetime    default current_timestamp comment '用户创建时间',
    primary key (`id`)
) engine = InnoDB
  default charset = utf8 comment ='博客用户表';

ALTER TABLE `blog_user`
    ADD `updated_at` datetime default current_timestamp comment '用户更新时间';