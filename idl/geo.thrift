namespace rs geo.storage

include "base.thrift"

struct VectorStoreRequest {
    1: string schema,
    2: string table,
    3: string name,
    4: string url,
    5: string ext,
}

struct VectorStoreResponse {
    1: base.BaseResp base,
}

service VectorStoreService {
    VectorStoreResponse Storage(1:VectorStoreRequest req),
}