
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table record
(
    id int not null auto_increment,
    contest_id int,
    user_id varbinary(255) not null,
    problem_id varchar(255) not null,
    language_id int not null,
    judge_result_id int not null default 1,
    compile_info text,
    time_cost int,
    memory_cost int,
    score int,
    
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,

    primary key(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table record;
