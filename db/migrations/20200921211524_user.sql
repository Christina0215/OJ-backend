
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table user
(
    id varbinary(255) not null,
    role_id boolean not null default 2,
    email varchar(255) not null,
    username varchar(255) not null,
    password varchar(255) not null,
    gender boolean not null default 0,
    introduction text not null,
    avatar varchar(255),
    school varchar(255),
    company varchar(255),
    github varchar(255),
    solved int default 0,

    created_at datetime,
    updated_at datetime,
    deleted_at datetime
);
insert into user(id, role_id, email,username, password, gender,introduction) values('b4342ace-34a5-4f9e-9690-b1c97876667a',1,'1532706870@qq.com','admin','ca9e680399decb9dd10d0cc4acda282c05e905174ab331bd9503e9f2e3b59f07',true,'admin');
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table user;
