create table user
(
    id          bigint auto_increment
        primary key,
    user_id     bigint                              not null,
    username    varchar(64)                         not null,
    phone       varchar(64)                         not null,
    password    varchar(64)                         not null,
    email       varchar(64)                         not null,
    gender      tinyint   default 0                 not null,
    avatar BLOB,
    avatar_url varchar(255),
    create_time timestamp default CURRENT_TIMESTAMP null,
    update_time timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    constraint idx_user_id
        unique (user_id),
    constraint idx_username
        unique (username)
)
    collate = utf8mb4_general_ci;

CREATE INDEX idx_phone ON user (phone);
CREATE INDEX idx_email ON user (email);
CREATE INDEX idx_create_time ON user (create_time);

INSERT INTO go_bbs.user (id, user_id, username, password, email, gender, create_time, update_time) VALUES (1, 28018727488323585, '小明', '313233343536639a9119599647d841b1bef6ce5ea293', null, 0, '2020-07-12 07:01:03', '2020-07-12 07:01:03');
INSERT INTO go_bbs.user (id, user_id, username, password, email, gender, create_time, update_time) VALUES (2, 4183532125556736, '小李', '313233639a9119599647d841b1bef6ce5ea293', null, 0, '2020-07-12 13:03:51', '2020-07-12 13:03:51');