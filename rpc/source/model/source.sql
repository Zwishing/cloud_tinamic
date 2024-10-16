CREATE SCHEMA IF NOT EXISTS data_source;

CREATE TABLE IF NOT EXISTS data_source.base_info(
    id serial PRIMARY KEY,
    key uuid NOT NULL UNIQUE,
    name varchar(255) NOT NULL,
    source_category integer NOT NULL,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp DEFAULT NULL
);

COMMENT ON TABLE data_source.base_info IS '存储数据源的基础信息表';

COMMENT ON COLUMN data_source.base_info.id IS '自增ID';
COMMENT ON COLUMN data_source.base_info.key IS '数据源的唯一标识';
COMMENT ON COLUMN data_source.base_info.name IS '数据源的名称';
COMMENT ON COLUMN data_source.base_info.source_category IS '数据源的数据类型：1-矢量，2-影像';


CREATE INDEX origin_info_id_index ON data_source.base_info(id);
CREATE INDEX origin_info_uuid_index ON data_source.base_info(key);
CREATE INDEX origin_info_category_index ON data_source.base_info(source_category);

CREATE TABLE IF NOT EXISTS data_source.storage(
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

COMMENT ON TABLE data_source.storage IS '存储所有源数据的表';

COMMENT ON COLUMN data_source.storage.id IS '唯一标识id';
COMMENT ON COLUMN data_source.storage.name IS '数据名称';
COMMENT ON COLUMN data_source.storage.storage_category IS '类型：文件-1和文件夹-2';
COMMENT ON COLUMN data_source.storage.key IS '数据的唯一标识';
COMMENT ON COLUMN data_source.storage.size IS '文件大小，单位kb';
COMMENT ON COLUMN data_source.storage.modified_time IS '修改时间';
COMMENT ON COLUMN data_source.storage.parent_key IS '所属的父文件夹的key,可以为空';
COMMENT ON COLUMN data_source.storage.path IS '数据源的minio的存储路径';

CREATE INDEX storage_id_index ON data_source.storage(id);
CREATE INDEX storage_uuid_index ON data_source.storage(key);
CREATE INDEX storage_parent_uuid_index ON data_source.storage(parent_key);

-- cloud_optimized_source
CREATE TABLE IF NOT EXISTS data_source.cloud_optimized(
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

COMMENT ON TABLE data_source.cloud_optimized IS '统一数据存储表';

COMMENT ON COLUMN data_source.cloud_optimized.id IS '自增ID';
COMMENT ON COLUMN data_source.cloud_optimized.key IS '数据的唯一标识';
COMMENT ON COLUMN data_source.cloud_optimized.source_key IS '数据源的唯一标识';
COMMENT ON COLUMN data_source.cloud_optimized.source_category IS '数据源的数据类型：1-矢量，2-影像';
COMMENT ON COLUMN data_source.cloud_optimized.size IS '文件大小，单位kb';
COMMENT ON COLUMN data_source.cloud_optimized.path IS '存储路径';
COMMENT ON COLUMN data_source.cloud_optimized.modified_time IS '修改时间';
COMMENT ON COLUMN data_source.cloud_optimized.created_at IS '创建时间';
COMMENT ON COLUMN data_source.cloud_optimized.updated_at IS '更新时间';
COMMENT ON COLUMN data_source.cloud_optimized.deleted_at IS '删除时间';

CREATE INDEX cloud_optimized_id_index ON data_source.cloud_optimized(id);
CREATE INDEX cloud_optimized_key_index ON data_source.cloud_optimized(key);
CREATE INDEX cloud_optimized_source_key_index ON data_source.cloud_optimized(source_key);
CREATE INDEX cloud_optimized_source_category_index ON data_source.cloud_optimized(source_category);


INSERT INTO data_source.base_info("key","name","source_category") VALUES ('9269d343-c2c9-b175-a9ac-c6f668ebfc78','矢量',1);
INSERT INTO data_source.base_info("key","name","source_category") VALUES ('584422c0-1acb-4295-a549-e43723120c8d','影像',2);

INSERT INTO data_source.storage(key, name, storage_category,
                                path)
VALUES ('9269d343-c2c9-b175-a9ac-c6f668ebfc78','矢量',2,'/vector');

INSERT INTO data_source.storage(key,name, storage_category,
                                path)
VALUES ('584422c0-1acb-4295-a549-e43723120c8d','影像',2,'/imagery');