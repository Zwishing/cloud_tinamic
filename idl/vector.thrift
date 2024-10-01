namespace go service.vector

include "base.thrift"
include "source.thrift"
include "service.thrift"

struct SpatialExtent {
    1: list<double> bbox
    2: string crs
}

struct Extent {
    1: SpatialExtent spatial_extent
}

struct Collection{
    1:string service_key
    2:string title
    3:source.SourceCategory source_category
    4:service.ServiceCategory service_category
    5:string crs
    6:Extent extent
    7:binary thumbnail
}

struct QueryParameters {
    1: i32 limit,        // 限制返回的记录数
    2: list<string> properties,  // 属性字段列表
    3: i32 resolution,   // 分辨率
    4: i32 buffer,       // 缓冲区大小
    5: string filter,    // 过滤条件
    6: i32 filterCrs     // 过滤条件的坐标参考系（CRS）
}

struct GetCollectionsResponse {
    1: base.BaseResp base,
    2: list<Collection> collections
    3: i64 number_returned
}

service VectorService{
    GetCollectionsResponse GetCollections(1:i64 pageSize,2:i64 page)
    void Publish(1:string req)
    binary GetTile(1:string service_key,2:i32 x,3:i32 y,4:i8 zoom,5:string ext,6:QueryParameters params)
}