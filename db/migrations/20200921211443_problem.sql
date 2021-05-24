
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table problem
(
    id varbinary(255) not null,
    user_id varbinary(255) not null,
    title varchar(255) not null,
    type varchar(255) not null,
    difficulty varchar(255) not null,
    content TEXT not null,
    samples varchar(255) not null,
    time_limit varchar(255) not null,
    memory_limit varchar(255) not null,
    standard_input varchar(255) not null,
    standard_output varchar(255) not null,
    tip varchar(255) not null,
    testdata_number int not null,

    created_at datetime,
    updated_at datetime,
    deleted_at datetime,

    primary key(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table problem;
