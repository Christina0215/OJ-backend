
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table token
(
    id int not null auto_increment,
    user_id varbinary(255) not null,
    token varchar(255) not null,
    expired_at datetime,

    created_at datetime,
    updated_at datetime,
    deleted_at datetime,

    primary key(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table token;
