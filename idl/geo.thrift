namespace rs data.storage
namespace go data.storage

include "base.thrift"

struct StoreRequest {
    1: required string schema,
    2: required string table,
    3: required string name,
    4: required string url,
    5: required string ext,
}

struct StoreResponse {
    1: required base.BaseResp base,
}

service StoreService {
    StoreResponse VectorStorage(1:StoreRequest req),
}