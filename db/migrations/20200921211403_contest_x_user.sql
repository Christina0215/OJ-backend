
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table contest_user
(
    id int not null auto_increment,
    contest_id varbinary(255),
    user_id varbinary(255),

    primary key(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table contest_user;
