-- DB 作成
-- 文字コード utf8mb4 で大文字/小文字を区別しない
-- utf8mb4 のとき 1 文字 4 byte
create database if not exists gin_gorm_logrus_basic_db
    character set utf8mb4
    collate utf8mb4_general_ci;

set foreign_key_checks=0;