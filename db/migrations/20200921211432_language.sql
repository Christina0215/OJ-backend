
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table language
(
    id int not null auto_increment,
    display_name varchar(255) not null,
    extension varchar(255) not null,

    primary key(id)
);
insert into language(display_name, extension) values('C','c');
insert into language(display_name, extension) values('C++','cc');
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table language;
