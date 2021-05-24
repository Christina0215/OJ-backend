
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table role
(
    id int not null auto_increment,
    alias varchar(255) not null,
    name varchar(255) not null,

    primary key(id)
);
insert into role(alias,name) values ('admin','管理员');
insert into role(alias,name) values ('user','用户');
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table role;
