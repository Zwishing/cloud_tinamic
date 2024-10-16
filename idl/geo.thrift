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

struct ToGeoParquetStorageRequest{
    1:required string source_path,
    2:required string bucket_name,
    3:required string storage_name,
}

struct ToGeoParquetStorageResponse {
    1: required base.BaseResp base,
    2: string dest_path,
    3: i64 size,
}



service StoreService {
    StoreResponse VectorStorage(1:StoreRequest req),
    ToGeoParquetStorageResponse ToGeoParquetStorage(1:ToGeoParquetStorageRequest req),
}