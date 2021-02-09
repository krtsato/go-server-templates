use gglb_db;

-- 中間テーブル: groups は accounts を所有する
create table if not exists groups_accounts(
    id int not null auto_increment,
    group_id int not null,
    account_id int not null,
    primary key(id),
    unique(group_id, account_id),
    constraint fk_groups_accounts_group_id
        foreign key(group_id)
        references `groups`(id)
        on update cascade
        on delete cascade,
    constraint fk_groups_accounts_account_id
        foreign key(account_id)
        references accounts(id)
        on update cascade
        on delete cascade
);