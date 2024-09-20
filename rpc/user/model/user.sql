CREATE SCHEMA IF NOT EXISTS user_info;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- 账号表:记录登录账号信息
CREATE TABLE IF NOT EXISTS user_info.account(
    id serial PRIMARY KEY,
    user_id uuid DEFAULT uuid_generate_v4() NOT NULL,
    user_account varchar(255) UNIQUE NOT NULL,
    category smallint NOT NULL,
    deleted_at bool DEFAULT NULL,
    created_at timestamp,
    updated_at timestamp
);
-- 注释
COMMENT ON TABLE user_info.account IS '账号表:记录登录账号信息';

COMMENT ON COLUMN user_info.account.id IS '账号ID';
COMMENT ON COLUMN user_info.account.user_id IS '用户唯一标识';
COMMENT ON COLUMN user_info.account.user_account IS '登录账号';
COMMENT ON COLUMN user_info.account.category IS '账号类别,1=用户名，2=邮箱，3=手机号';
-- 索引
CREATE INDEX account_id_index ON user_info.account(id);
CREATE INDEX account_user_id_index ON user_info.account(user_id);


-- 用户表:记录用户基本信息和密码
CREATE TABLE IF NOT EXISTS user_info.user(
    id serial PRIMARY KEY,
    user_id uuid NOT NULL,
    name varchar(255),
    avatar bytea,
    phone_number varchar(11),
    salt varchar(64),
    password varchar(64),
    deleted_at bool DEFAULT NULL,
    created_at timestamp,
    updated_at timestamp
);
-- 注释
COMMENT ON TABLE user_info.user IS '用户表:记录用户基本信息和密码';

COMMENT ON COLUMN user_info.user.id IS '用户ID';
COMMENT ON COLUMN user_info.user.user_id IS '用户唯一标识';
COMMENT ON COLUMN user_info.user.name IS '姓名';
COMMENT ON COLUMN user_info.user.avatar IS '用户头像图片';
COMMENT ON COLUMN user_info.user.phone_number IS '手机号码';
COMMENT ON COLUMN user_info.user.salt IS '密码加盐';
COMMENT ON COLUMN user_info.user.password IS '登录密码';

-- 索引
CREATE INDEX user_id_index ON user_info.user(id);
CREATE INDEX user_uuid_index ON user_info.user(user_id);

-- 创建一个用户 原始的账号和密码：admin admin123
INSERT INTO user_info.user(user_id,name,phone_number,salt,password)
VALUES ('cc9de34f-731f-4c45-9ecd-ef84a07b7427','张三','15600755813','','JAvlGPq9JyTdtvBO6x2llnRI1+gxwIyPqCKAn3THIKk=');

INSERT INTO user_info.account(user_id,user_account,category)
VALUES ('cc9de34f-731f-4c45-9ecd-ef84a07b7427','admin',1);
