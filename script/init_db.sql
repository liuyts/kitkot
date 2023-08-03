
create schema douin_user;
create table douyin_user.user
(
    id                bigint       not null
        primary key,
    username          varchar(30)  not null,
    password          varchar(30)  not null,
    avatar            varchar(255) not null,
    back_ground_image varchar(255) not null,
    signature         varchar(255) not null,
    constraint user_username_uindex
        unique (username)
);


create schema douyin_chat;
create table douyin_chat.message
(
    id           bigint auto_increment
        primary key,
    from_user_id bigint        not null,
    to_user_id   bigint        not null,
    content      varchar(255)  not null,
    create_time  int           not null
);



