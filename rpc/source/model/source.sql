CREATE SCHEMA IF NOT EXISTS data_source;

CREATE TABLE IF NOT EXISTS data_source.base_info(
    id serial PRIMARY KEY,
    source_id uuid NOT NULL UNIQUE,
    name varchar(255) NOT NULL,
    source_category integer NOT NULL,
    deleted_at bool DEFAULT NULL,
    created_at timestamp,
    updated_at timestamp
);

COMMENT ON TABLE data_source.base_info IS '存储数据源的基础信息表';

COMMENT ON COLUMN data_source.base_info.id IS '自增ID';
COMMENT ON COLUMN data_source.base_info.source_id IS '数据源的唯一标识';
COMMENT ON COLUMN data_source.base_info.name IS '数据源的名称';
COMMENT ON COLUMN data_source.base_info.source_category IS '数据源的数据类型：1-矢量，2-影像';


CREATE INDEX origin_info_id_index ON data_source.base_info(id);
CREATE INDEX origin_info_uuid_index ON data_source.base_info(source_id);

CREATE TABLE IF NOT EXISTS data_source.storage(
    id serial PRIMARY KEY,
    source_id uuid NOT NULL UNIQUE,
    parent_id uuid,
    name varchar(255),
    storage_category int2 NOT NULL,
    key varchar(255),
    size bigint default 0,
    modified_time timestamptz,
    path varchar(255) NOT NULL,
    deleted_at bool DEFAULT NULL,
    created_at timestamp,
    updated_at timestamp
);

COMMENT ON TABLE data_source.storage IS '存储所有源数据的表';

COMMENT ON COLUMN data_source.storage.id IS '唯一标识id';
COMMENT ON COLUMN data_source.storage.name IS '数据名称';
COMMENT ON COLUMN data_source.storage.source_id IS 'uuid';
COMMENT ON COLUMN data_source.storage.storage_category IS '类型：文件-1和文件夹-2';
COMMENT ON COLUMN data_source.storage.key IS '数据的唯一标识';
COMMENT ON COLUMN data_source.storage.size IS '文件大小，单位kb';
COMMENT ON COLUMN data_source.storage.modified_time IS '修改时间';
COMMENT ON COLUMN data_source.storage.parent_id IS '所属的父文件夹的id,可以为空';
COMMENT ON COLUMN data_source.storage.path IS '数据源的minio的存储路径';

CREATE INDEX storage_id_index ON data_source.storage(id);
CREATE INDEX storage_uuid_index ON data_source.storage(source_id);
CREATE INDEX storage_parent_uuid_index ON data_source.storage(parent_id);

INSERT INTO data_source.base_info("source_id","name","source_category")VALUES ('9269d343-c2c9-b175-a9ac-c6f668ebfc78','矢量',1);

INSERT INTO data_source.storage(source_id, name, storage_category,
                                 path)
VALUES ('9269d343-c2c9-b175-a9ac-c6f668ebfc78','矢量',2,'/vector');