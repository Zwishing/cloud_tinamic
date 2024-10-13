namespace go map.processor
namespace py map.processor

include "base.thrift"

struct VectorThumbnailRequest{
    1:string file_path,
    2:i32 width,
    3:i32 height,
}

struct VectorThumbnailRespose {
    1:base.BaseResp base,
    2:binary thumbnail,
}

service MapProcessorService {
    VectorThumbnailRespose VectorThumbnail(1: VectorThumbnailRequest req)
}
