
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table solution
(
    id varbinary(255) not null,
    user_id varbinary(255) not null,
    title varchar(255) not null,
    content TEXT not null,
    problem_id varchar(255) not null,
    language varchar(255) not null,

    created_at datetime,
    updated_at datetime,
    deleted_at datetime,

    primary key(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table solution;
