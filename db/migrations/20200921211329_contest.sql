
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table contest
(
    id varbinary(255) not null,
    user_id varbinary(255),
    title varchar(255),

    introduction varchar(255),
    notification varchar(255),
    enabled boolean,

    start_at bigint,
    end_at bigint,

    created_at datetime,
    updated_at datetime,
    deleted_at datetime,

    primary key (id)
);
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table contest;
