CREATE TABLE comment
(
    id           bigint auto_increment primary key,
    post_id      bigint                              not null comment '帖子id',
    user_name    varchar(64)                         not null comment '用户昵称',
    content      varchar(8192)                       not null comment '评论内容',
    author_id    bigint                              not null comment '评论作者的用户id',
    create_time  timestamp default CURRENT_TIMESTAMP null comment '评论创建时间',
    picture BLOB comment '评论照片',
    picture_url varchar(255) comment 'oss内url',
    constraint fk_comment_post_id
        foreign key (post_id) references post (post_id)
        on delete cascade
)
collate = utf8mb4_general_ci;

CREATE INDEX idx_post_id ON comment (post_id);
CREATE INDEX idx_create_time ON comment (create_time);