use gin_gorm_logrus_basic_db;

-- groups table
create table if not exists `groups`(
    id int not null auto_increment,
    name varchar(80) not null,
    note varchar(400) not null,
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp,
    deleted_at datetime null,
    primary key(id)
);