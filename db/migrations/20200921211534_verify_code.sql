
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table verify_code
(
    id int not null auto_increment,
    code varchar(255) not null,
    email varchar(255) not null,

    expired_at datetime not null,

    created_at datetime,
    updated_at datetime,
    deleted_at datetime,

    primary key(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table verify_code;
