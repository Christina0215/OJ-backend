
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table contest_x_problem
(
    id int not null auto_increment,
    contest_id varbinary(255),
    problem_id varbinary(255),
    `order` int,
    base_score int,

    created_at datetime,
    updated_at datetime,
    deleted_at datetime,

    primary key(id)
);
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table contest_x_problem;
