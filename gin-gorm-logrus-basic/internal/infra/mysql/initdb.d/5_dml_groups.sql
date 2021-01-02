use gin_gorm_logrus_basic_db;

delete from `groups`;
insert into `groups` (id, name, note, created_at, updated_at, deleted_at) values (1, '放課後ティータイム', '放課後ティータイムは永遠に放課後です！', '2000-04-02 00:00:00', '2000-09-28 00:00:00', null);
insert into `groups` (id, name, note, created_at, updated_at, deleted_at) values (2, 'ギター陣', 'スカイハイ！', '2001-01-01 00:00:01', '2011-01-01 00:00:01', null);
insert into `groups`(id, name, note, created_at, updated_at, deleted_at) values (3, 'リズム陣', 'ゴリゴリ！', '2002-01-01 00:00:02', '2012-01-01 00:00:02', null);
insert into `groups` (id, name, note, created_at, updated_at, deleted_at) values (4, '作曲陣', '', '2003-01-01 00:00:03', '2013-01-01 00:00:03', null);
insert into `groups` (id, name, note, created_at, updated_at, deleted_at) values (5, '幹部', '', '2004-01-01 00:00:04', '2014-01-01 00:00:04', '2014-01-01 00:00:14');