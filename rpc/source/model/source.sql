CREATE SCHEMA IF NOT EXISTS source;

CREATE TABLE IF NOT EXISTS source.info(
    id serial PRIMARY KEY,
    key uuid NOT NULL UNIQUE,
    name varchar(255) NOT NULL,
    source_category integer NOT NULL,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp DEFAULT NULL
);

COMMENT ON TABLE source.info IS '存储数据源的基础信息表';

COMMENT ON COLUMN source.info.id IS '自增ID';
COMMENT ON COLUMN source.info.key IS '数据源的唯一标识';
COMMENT ON COLUMN source.info.name IS '数据源的名称';
COMMENT ON COLUMN source.info.source_category IS '数据源的数据类型：1-矢量，2-影像';


CREATE INDEX origin_info_id_index ON source.info(id);
CREATE INDEX origin_info_uuid_index ON source.info(key);
CREATE INDEX origin_info_category_index ON source.info(source_category);

CREATE TABLE IF NOT EXISTS source.original(
    id serial PRIMARY KEY,
    key uuid NOT NULL UNIQUE,
    parent_key uuid,
    name varchar(255),
    storage_category int2 NOT NULL,
    size bigint default 0,
    modified_time timestamp,
    path varchar(255) NOT NULL,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp DEFAULT NULL
);

COMMENT ON TABLE source.original IS '存储所有源数据的表';

COMMENT ON COLUMN source.original.id IS '唯一标识id';
COMMENT ON COLUMN source.original.name IS '数据名称';
COMMENT ON COLUMN source.original.storage_category IS '类型：文件-1和文件夹-2';
COMMENT ON COLUMN source.original.key IS '数据的唯一标识';
COMMENT ON COLUMN source.original.size IS '文件大小，单位kb';
COMMENT ON COLUMN source.original.modified_time IS '修改时间';
COMMENT ON COLUMN source.original.parent_key IS '所属的父文件夹的key,可以为空';
COMMENT ON COLUMN source.original.path IS '数据源的minio的存储路径';

CREATE INDEX original_id_index ON source.original(id);
CREATE INDEX original_uuid_index ON source.original(key);
CREATE INDEX original_parent_uuid_index ON source.original(parent_key);

-- cloud_optimized_source
CREATE TABLE IF NOT EXISTS source.cloud_optimized(
    id serial PRIMARY KEY,
    key uuid NOT NULL UNIQUE,
    source_key uuid NOT NULL UNIQUE,
    source_category integer NOT NULL,
    size bigint default 0,
    path varchar(255) NOT NULL,
    modified_time timestamp,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp DEFAULT NULL
);

COMMENT ON TABLE source.cloud_optimized IS '统一数据存储表';

COMMENT ON COLUMN source.cloud_optimized.id IS '自增ID';
COMMENT ON COLUMN source.cloud_optimized.key IS '数据的唯一标识';
COMMENT ON COLUMN source.cloud_optimized.source_key IS '数据源的唯一标识';
COMMENT ON COLUMN source.cloud_optimized.source_category IS '数据源的数据类型：1-矢量，2-影像';
COMMENT ON COLUMN source.cloud_optimized.size IS '文件大小，单位kb';
COMMENT ON COLUMN source.cloud_optimized.path IS '存储路径';
COMMENT ON COLUMN source.cloud_optimized.modified_time IS '修改时间';
COMMENT ON COLUMN source.cloud_optimized.created_at IS '创建时间';
COMMENT ON COLUMN source.cloud_optimized.updated_at IS '更新时间';
COMMENT ON COLUMN source.cloud_optimized.deleted_at IS '删除时间';

CREATE INDEX cloud_optimized_id_index ON source.cloud_optimized(id);
CREATE INDEX cloud_optimized_key_index ON source.cloud_optimized(key);
CREATE INDEX cloud_optimized_source_key_index ON source.cloud_optimized(source_key);
CREATE INDEX cloud_optimized_source_category_index ON source.cloud_optimized(source_category);


INSERT INTO source.info("key","name","source_category") VALUES ('9269d343-c2c9-b175-a9ac-c6f668ebfc78','矢量',1);
INSERT INTO source.info("key","name","source_category") VALUES ('584422c0-1acb-4295-a549-e43723120c8d','影像',2);

INSERT INTO source.original(key, name, storage_category,
                                path)
VALUES ('9269d343-c2c9-b175-a9ac-c6f668ebfc78','矢量',2,'/vector');

INSERT INTO source.original(key,name, storage_category,
                                path)
VALUES ('584422c0-1acb-4295-a549-e43723120c8d','影像',2,'/imagery');