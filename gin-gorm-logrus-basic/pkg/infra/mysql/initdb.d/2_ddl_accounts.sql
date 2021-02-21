use gglb_db;

-- accounts table
create table if not exists accounts(
    id int not null auto_increment,
    name varchar(80) not null,
    note varchar(400) null,
    thx_count int unsigned not null,
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp,
    deleted_at datetime null,
    primary key(id)
);