namespace go map.processor
namespace py map.processor

include "base.thrift"

struct VectorThumbnailRequest{
    1:string cloud_optimized_path,
    2:string cloud_optimized_bucket_name,
    3:i32 width,
    4:i32 height,
}

struct VectorThumbnailRespose {
    1:base.BaseResp base,
    2:binary thumbnail,
}

service MapProcessorService {
    VectorThumbnailRespose VectorThumbnail(1: VectorThumbnailRequest req)
}
