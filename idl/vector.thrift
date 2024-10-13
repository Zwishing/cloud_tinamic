namespace go service.vector

include "base.thrift"
include "source.thrift"
include "collection.thrift"


struct QueryParameters {
    1: i32 limit,        // 限制返回的记录数
    2: list<string> properties,  // 属性字段列表
    3: i32 resolution,   // 分辨率
    4: i32 buffer,       // 缓冲区大小
    5: string filter,    // 过滤条件
    6: i32 filterCrs     // 过滤条件的坐标参考系（CRS）
}

service VectorService{
    binary GetTile(1:string service_key,2:i32 x,3:i32 y,4:i8 zoom,5:string ext,6:QueryParameters params)
}