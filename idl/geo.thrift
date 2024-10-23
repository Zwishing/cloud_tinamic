namespace rs data.storage
namespace go data.storage

include "base.thrift"

struct VectorToPGStorageRequest {
    1: required string schema,
    2: required string table,
    3: required string name,
    4: required string cloud_optimized_bucket_name,
    5: required string cloud_optimized_path,
    // 6: optional i64 size, // 文件大小
}

struct VectorToPGStorageResponse {
    1: required base.BaseResp base,
}

struct ToGeoParquetStorageRequest{
    1:required string source_bucket
    2:required string source_path,
    3:required string dest_bucket,
    4:required string dest_path,
}

struct ToGeoParquetStorageResponse {
    1: required base.BaseResp base,
    2: string dest_path,
    3: i64 size,
}



service StoreService {
    VectorToPGStorageResponse VectorToPGStorage(1:VectorToPGStorageRequest req),
    ToGeoParquetStorageResponse ToGeoParquetStorage(1:ToGeoParquetStorageRequest req),
}