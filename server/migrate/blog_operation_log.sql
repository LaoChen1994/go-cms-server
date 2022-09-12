Create TABLE IF NOT EXISTS `blog_operation_log`
(
    id             int(10) primary key auto_increment not null comment '主键，日志编号',
    user_id        int(10)                            not null comment '操作的用户id',
    target_type    tinyint                            not null default 1 comment '操作对象类型: 1-文章 2-标签',
    target_id      int(10)                            not null comment '操作对象编号',
    operation_type tinyint                            not null comment '操作类型: 1-创建 2-修改 3-查询 4-删除'
) charset = 'utf8' comment '增删改查日志表';

create index `target_index`
    on `blog_operation_log` (target_id, target_type)
    comment '按素材查询操作记录';

create index `user_index`
    on `blog_operation_log` (user_id, id)
    comment '查询用户对某一篇文章的操作记录';


ALTER TABLE `blog_operation_log`
    ADD created_at timestamp default CURRENT_TIMESTAMP not null;
ALTER TABLE `blog_operation_log`
    ADD updated_at timestamp default CURRENT_TIMESTAMP not null;
ALTER TABLE `blog_operation_log`
    ADD deleted_at timestamp;