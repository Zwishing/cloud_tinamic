CREATE SCHEMA IF NOT EXISTS service;

-- 创建源信息表
CREATE TABLE IF NOT EXISTS service.info (
    id SERIAL PRIMARY KEY, -- 每个服务信息条目的唯一标识符
    service_key UUID NOT NULL UNIQUE, -- 服务的唯一键
    title VARCHAR(255) NOT NULL, -- 数据名称
    source_key UUID NOT NULL UNIQUE, -- 源数据的唯一键
    source_schema VARCHAR(255) NOT NULL, -- 源数据所在的模式
    source_category integer NOT NULL, -- 源数据的类型
    sird INT, -- 空间参考标识符
    created_at TIMESTAMP DEFAULT NOW(), -- 创建时间戳
    updated_at TIMESTAMP DEFAULT NOW(), -- 最后更新时间戳
    deleted_at TIMESTAMP DEFAULT NULL -- 软删除时间戳
);

COMMENT ON TABLE service.info IS '包含服务信息的表';
COMMENT ON COLUMN service.info.id IS '每个服务信息条目的唯一标识符';
COMMENT ON COLUMN service.info.service_key IS '服务的唯一键';
COMMENT ON COLUMN service.info.title IS '服务名称';
COMMENT ON COLUMN service.info.source_key IS '源数据的唯一键';
COMMENT ON COLUMN service.info.source_schema IS '源数据所在的模式';
COMMENT ON COLUMN service.info.source_category IS '源数据的类型';
COMMENT ON COLUMN service.info.sird IS '空间参考标识符';
COMMENT ON COLUMN service.info.created_at IS '创建时间戳';
COMMENT ON COLUMN service.info.updated_at IS '最后更新时间戳';
COMMENT ON COLUMN service.info.deleted_at IS '软删除时间戳';

-- 信息表的索引
CREATE INDEX idx_info_service_key ON service.info(service_key);
CREATE INDEX idx_info_source_key ON service.info(source_key);

-- 创建集合表
CREATE TABLE IF NOT EXISTS service.collection (
    id SERIAL PRIMARY KEY, -- 每个集合条目的唯一标识符
    service_key UUID NOT NULL UNIQUE, -- 服务的唯一键
    title VARCHAR(255) NOT NULL, -- 服务名称
    bbox GEOMETRY(Polygon, 4326), -- 集合的几何边界
    center GEOMETRY(Point, 4326), -- 集合的几何中心点
    srid INT, -- 几何的空间参考标识符
    service_category INT NOT NULL, -- 服务类别:mvt、feature、wmts、wms、tms
    description TEXT, -- 服务描述
    thumbnail BYTEA, -- 服务的缩略图
    created_at TIMESTAMP DEFAULT NOW(), -- 创建时间戳
    updated_at TIMESTAMP DEFAULT NOW(), -- 最后更新时间戳
    deleted_at TIMESTAMP DEFAULT NULL -- 软删除时间戳
);

COMMENT ON TABLE service.collection IS '包含服务集合的表';
COMMENT ON COLUMN service.collection.id IS '每个集合条目的唯一标识符';
COMMENT ON COLUMN service.collection.service_key IS '服务的唯一键';
COMMENT ON COLUMN service.collection.title IS '集合名称';
COMMENT ON COLUMN service.collection.bbox IS '集合的几何边界';
COMMENT ON COLUMN service.collection.center IS '集合的几何中心点';
COMMENT ON COLUMN service.collection.srid IS '几何的空间参考标识符';
COMMENT ON COLUMN service.collection.service_category IS '服务类别';
COMMENT ON COLUMN service.collection.description IS '服务描述';
COMMENT ON COLUMN service.collection.thumbnail IS '服务的缩略图';
COMMENT ON COLUMN service.collection.created_at IS '创建时间戳';
COMMENT ON COLUMN service.collection.updated_at IS '最后更新时间戳';
COMMENT ON COLUMN service.collection.deleted_at IS '软删除时间戳';


CREATE INDEX idx_collection_service_key ON service.collection(service_key);
-- 创建 GIST 索引
CREATE INDEX idx_collection_bounds ON service.collection USING GIST(bbox);
CREATE INDEX idx_collection_center ON service.collection USING GIST(center);

-- 创建矢量表
CREATE TABLE IF NOT EXISTS service.vector (
    id SERIAL PRIMARY KEY, -- 每个矢量条目的唯一标识符
    source_key UUID NOT NULL UNIQUE, -- 服务的唯一键
    title VARCHAR(255) NOT NULL, -- 矢量服务名称
    geometry_category VARCHAR(255), -- 几何类型: 点、线、面
    geometry_field VARCHAR(255), -- 几何字段名称
    field_count INT, -- 字段数
    record_count INT, -- 记录数
    properties jsonb, -- 记录字段的名称和类型
    created_at TIMESTAMP DEFAULT NOW(), -- 创建时间戳
    updated_at TIMESTAMP DEFAULT NOW(), -- 最后更新时间戳
    deleted_at TIMESTAMP DEFAULT NULL -- 软删除时间戳
    );

COMMENT ON TABLE service.vector IS '包含矢量服务的表';
COMMENT ON COLUMN service.vector.id IS '每个矢量条目的唯一标识符';
COMMENT ON COLUMN service.vector.source_key IS '服务的唯一键';
COMMENT ON COLUMN service.vector.title IS '矢量服务名称';
COMMENT ON COLUMN service.vector.geometry_category IS '几何类型';
COMMENT ON COLUMN service.vector.geometry_field IS '几何字段名称';
COMMENT ON COLUMN service.vector.field_count IS '字段数';
COMMENT ON COLUMN service.vector.record_count IS '记录数';
COMMENT ON COLUMN service.vector.created_at IS '创建时间戳';
COMMENT ON COLUMN service.vector.updated_at IS '最后更新时间戳';
COMMENT ON COLUMN service.vector.deleted_at IS '软删除时间戳';

-- 创建影像表
-- CREATE TABLE IF NOT EXISTS service.imagery (
--     id SERIAL PRIMARY KEY, -- 每个影像条目的唯一标识符
--     service_key UUID NOT NULL UNIQUE, -- 服务的唯一键
--     title VARCHAR(255) NOT NULL, -- 影像服务名称
--     created_at TIMESTAMP DEFAULT NOW(), -- 创建时间戳
--     updated_at TIMESTAMP DEFAULT NOW(), -- 最后更新时间戳
--     deleted_at TIMESTAMP DEFAULT NULL -- 软删除时间戳
-- );
--
-- COMMENT ON TABLE service.imagery IS '包含影像服务的表';
-- COMMENT ON COLUMN service.imagery.id IS '每个影像条目的唯一标识符';
-- COMMENT ON COLUMN service.imagery.service_key IS '服务的唯一键';
-- COMMENT ON COLUMN service.imagery.title IS '影像服务名称';
-- COMMENT ON COLUMN service.imagery.created_at IS '创建时间戳';
-- COMMENT ON COLUMN service.imagery.updated_at IS '最后更新时间戳';
-- COMMENT ON COLUMN service.imagery.deleted_at IS '软删除时间戳';

-- 创建样式表
-- CREATE TABLE IF NOT EXISTS service.style (
--     id SERIAL PRIMARY KEY, -- 每个样式条目的唯一标识符
--     service_key UUID NOT NULL UNIQUE, -- 服务的唯一键
--     style_key UUID NOT NULL UNIQUE, -- 样式的唯一键
--     title VARCHAR(255) NOT NULL, -- 样式名称
--     min_zoom INT, -- 样式的最小缩放级别
--     max_zoom INT, -- 样式的最大缩放级别
--     created_at TIMESTAMP DEFAULT NOW(), -- 创建时间戳
--     updated_at TIMESTAMP DEFAULT NOW(), -- 最后更新时间戳
--     deleted_at TIMESTAMP DEFAULT NULL -- 软删除时间戳
-- );
--
-- COMMENT ON TABLE service.style IS '包含服务样式的表';
-- COMMENT ON COLUMN service.style.id IS '每个样式条目的唯一标识符';
-- COMMENT ON COLUMN service.style.service_key IS '服务的唯一键';
-- COMMENT ON COLUMN service.style.style_key IS '样式的唯一键';
-- COMMENT ON COLUMN service.style.title IS '样式名称';
-- COMMENT ON COLUMN service.style.min_zoom IS '样式的最小缩放级别';
-- COMMENT ON COLUMN service.style.max_zoom IS '样式的最大缩放级别';
-- COMMENT ON COLUMN service.style.created_at IS '创建时间戳';
-- COMMENT ON COLUMN service.style.updated_at IS '最后更新时间戳';
-- COMMENT ON COLUMN service.style.deleted_at IS '软删除时间戳';
--
-- -- 样式表的索引
-- CREATE INDEX idx_style_service_key ON service.style(service_key);
-- CREATE INDEX idx_style_key ON service.style(style_key);
