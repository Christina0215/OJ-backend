
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

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table user;
