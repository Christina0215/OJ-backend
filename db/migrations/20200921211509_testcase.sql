
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table testcase
(
    id int not null auto_increment,
    record_id int not null,
    testdata_index int not null,
    judge_result_id int not null,
    time_cost int not null,
    memory_cost int not null,
    diff text,

    primary key(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table testcase;
