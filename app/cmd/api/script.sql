use `go-zero-oj`;
create table category
(
    id         int unsigned auto_increment
        primary key,
    identity   varchar(36)  null,
    name       varchar(100) null comment '分类id',
    parent_id  int          null comment '父级id',
    created_at datetime     null comment '创建时间',
    updated_at datetime     null comment '修改时间',
    deleted_at datetime     null comment '删除时间(软删除)'
);

create table problem
(
    id          int          not null
        primary key,
    identity    varchar(36)  null comment '唯一标识',
    title       varchar(255) null comment '标题',
    max_runtime int          null comment '最大运行时间',
    max_mem     int          null comment '最大运行内存',
    path        varchar(36)  null comment '数据集',
    content     varchar(255) null comment '题目内容',
    created_at  datetime     null comment '创建时间',
    updated_at  datetime     null comment '修改时间',
    deleted_at  datetime     null comment '删除时间(软删除)'
);

create table problem_category
(
    id          int auto_increment
        primary key,
    problem_id  int      not null,
    category_id int      not null,
    created_at  datetime null comment '创建时间',
    updated_at  datetime null comment '修改时间',
    deleted_at  datetime null comment '删除时间(软删除)'
);

create table submit
(
    id               int                not null
        primary key,
    identity         varchar(36)        null,
    problem_identity varchar(36)        null comment '问题的唯一标识',
    user_identity    varchar(36)        null comment '用户的唯一标识',
    status           tinyint default -1 not null comment '-1表示待判断，1表示答案正确，2表示答案错误，3表示运行超时，4表示运行超内存',
    language         tinyint            null comment '1表示go 2表示java 3表示rust',
    run_time         int                null comment '运行时间',
    run_mem          int                null comment '运行内存',
    created_at       datetime           null comment '创建时间',
    updated_at       datetime           null comment '修改时间',
    deleted_at       datetime           null comment '删除时间(软删除)'
);

create table user
(
    id                 int auto_increment
        primary key,
    identity           varchar(36)   null,
    name               varchar(36)   null comment '用户名称',
    password           varchar(128)  null comment '密码',
    phone              varchar(36)   null comment '电话',
    mail               varchar(100)  null comment '邮箱',
    finish_problem_num int default 0 null comment '完成问题的次数',
    submit_num         int default 0 null comment '用户提交次数',
    created_at         datetime      null comment '创建时间',
    updated_at         datetime      null comment '修改时间',
    deleted_at         datetime      null comment '删除时间(软删除)'
);


INSERT INTO `go-zero-oj`.problem (id, identity, title, max_runtime, max_mem, path, content, created_at, updated_at, deleted_at) VALUES (1, 'problem_1', '文章标题', 100, 100, null, '文章正文', '2023-09-11 17:20:25', '2023-09-11 17:20:27', null);
INSERT INTO `go-zero-oj`.problem (id, identity, title, max_runtime, max_mem, path, content, created_at, updated_at, deleted_at) VALUES (2, 'problem_2', '文章标题', 100, 100, null, '文章正文', '2023-09-11 17:20:25', '2023-09-11 17:20:27', null);
INSERT INTO `go-zero-oj`.problem (id, identity, title, max_runtime, max_mem, path, content, created_at, updated_at, deleted_at) VALUES (3, 'problem_3', '文章标题', 100, 100, null, '文章正文', '2023-09-14 09:44:18', '2023-09-14 09:44:19', null);
INSERT INTO `go-zero-oj`.problem_category (id, problem_id, category_id, created_at, updated_at, deleted_at) VALUES (1, 1, 1, '2023-09-12 11:21:11', '2023-09-12 11:21:12', null);
INSERT INTO `go-zero-oj`.problem_category (id, problem_id, category_id, created_at, updated_at, deleted_at) VALUES (2, 1, 2, '2023-09-12 11:21:18', '2023-09-12 11:21:19', null);
INSERT INTO `go-zero-oj`.problem_category (id, problem_id, category_id, created_at, updated_at, deleted_at) VALUES (3, 2, 1, '2023-09-12 11:21:34', '2023-09-12 11:21:35', null);
INSERT INTO `go-zero-oj`.problem_category (id, problem_id, category_id, created_at, updated_at, deleted_at) VALUES (4, 3, 1, '2023-09-14 09:45:09', '2023-09-14 09:45:10', null);
INSERT INTO `go-zero-oj`.submit (id, identity, problem_identity, user_identity, status, language, run_time, run_mem, created_at, updated_at, deleted_at) VALUES (1, 'submit_1', 'problem_1', 'user_1', -1, 1, 80, 12, '2023-09-13 10:20:06', '2023-09-13 10:20:07', null);
INSERT INTO `go-zero-oj`.submit (id, identity, problem_identity, user_identity, status, language, run_time, run_mem, created_at, updated_at, deleted_at) VALUES (2, 'submit_2', 'problem_1', 'user_2', -1, 1, 100, 12, '2023-09-13 10:23:27', '2023-09-13 10:23:27', null);
INSERT INTO `go-zero-oj`.user (id, identity, name, password, phone, mail, finish_problem_num, submit_num, created_at, updated_at, deleted_at) VALUES (1, 'admin', 'admin', '$2a$10$wh6fiea0hn.ypvYwxcA.mu/EXD2sEsfb4SwVk2EHLdbGTLKgSTymm', '', '', 0, 0, '2023-09-11 16:15:06', '2023-09-11 16:15:06', null);
INSERT INTO `go-zero-oj`.user (id, identity, name, password, phone, mail, finish_problem_num, submit_num, created_at, updated_at, deleted_at) VALUES (2, 'user_1', '李四', '$2a$10$wh6fiea0hn.ypvYwxcA.mu/EXD2sEsfb4SwVk2EHLdbGTLKgSTymm', '', '', 0, 0, '2023-09-11 16:15:06', '2023-09-11 16:15:06', null);
INSERT INTO `go-zero-oj`.user (id, identity, name, password, phone, mail, finish_problem_num, submit_num, created_at, updated_at, deleted_at) VALUES (3, 'user_2', '王五', '$2a$10$t9w/Tv4u5fbLjTNsUleGZOdXupNrJLEF/S1g0XApyLf9oF0hN4ZES', '', '', 0, 0, '2023-09-13 09:48:43', '2023-09-13 09:48:43', null);
INSERT INTO `go-zero-oj`.user (id, identity, name, password, phone, mail, finish_problem_num, submit_num, created_at, updated_at, deleted_at) VALUES (4, 'user_3', '老刘', '$2a$10$mWELW.eml8M0bq5umwn57OxvX0iTqeszUY1.9JpKeGrcc3V3sV6Ou', '', '', 0, 0, '2023-09-13 09:51:01', '2023-09-13 09:51:01', null);
INSERT INTO `go-zero-oj`.user (id, identity, name, password, phone, mail, finish_problem_num, submit_num, created_at, updated_at, deleted_at) VALUES (5, 'user_4', '张七', '$2a$10$6PW7uEAVIVZIOAKkfcT1UuYoMlIKKA5SQUcMDML6jxmcjKu9R8jw.', '', '', 0, 0, '2023-09-13 09:53:41', '2023-09-13 09:53:41', null);
INSERT INTO `go-zero-oj`.user (id, identity, name, password, phone, mail, finish_problem_num, submit_num, created_at, updated_at, deleted_at) VALUES (6, 'user_5', '老王', '$2a$10$bOm/CQl2i2LsikUlQrZS..3XfA5EFWS1IgS6Oy2.uIIw4QGugcPAi', '', '', 0, 0, '2023-09-13 09:54:35', '2023-09-13 09:54:35', null);
INSERT INTO `go-zero-oj`.category (id, identity, name, parent_id, created_at, updated_at, deleted_at) VALUES (1, 'category_1', '数组', 0, '2023-09-12 11:19:59', '2023-09-12 11:20:00', null);
INSERT INTO `go-zero-oj`.category (id, identity, name, parent_id, created_at, updated_at, deleted_at) VALUES (2, 'category_2', '字符串', 0, '2023-09-12 11:20:09', '2023-09-12 11:20:10', null);
INSERT INTO `go-zero-oj`.category (id, identity, name, parent_id, created_at, updated_at, deleted_at) VALUES (3, 'category_3', 'map', 0, '2023-09-12 11:20:23', '2023-09-12 11:20:24', null);
