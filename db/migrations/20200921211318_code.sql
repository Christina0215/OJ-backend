
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table code
(
    id int not null auto_increment,
    record_id int not null,
    language_id int not null,
    filename VARCHAR(255) not null,

    primary key(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table code;
