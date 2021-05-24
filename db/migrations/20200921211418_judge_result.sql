
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table judge_result
(
    id int not null auto_increment,
    alias varchar(255) not null,
    en varchar(255) not null,
    zh varchar(255) not null,
    color varchar(255) not null,

    primary key(id)
);
insert into judge_result (alias, en,zh,color) values ("Pending", "Pending", "等待测评","#9e9e9e");
insert into judge_result (alias, en,zh,color) values ("Judging","Judging","正在测评","#2196f3");
insert into judge_result (alias, en,zh,color) values ("AC","Accepted","通过测试","#4caf50");
insert into judge_result (alias, en,zh,color) values ("PE","Presentation Error","输出格式错误","#ff9800");
insert into judge_result (alias, en,zh,color) values ("WA","Wrong Answer","错误答案","#f44336");
insert into judge_result (alias, en,zh,color) values ("OLE","Output Limit Exceeded","超出输出限制","#e91e63");
insert into judge_result (alias, en,zh,color) values ("TLE","Time Limit Exceeded","超出时间限制","#9c27b0");
insert into judge_result (alias, en,zh,color) values ("MLE","Memory Limit Exceeded","超出内存限制","#673ab7");
insert into judge_result (alias, en,zh,color) values ("CE","Compilation Error","编译错误","#ffeb3b");
insert into judge_result (alias, en,zh,color) values ("RE","Runtime Error","运行时错误","#ff5722");
insert into judge_result (alias, en,zh,color) values ("SE","System Error","系统错误","#000000");
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table judge_result;
