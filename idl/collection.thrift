namespace go service.collection

include "base.thrift"
include "source.thrift"

enum ServiceCategory {
    FEATURE = 0,
    MVT = 1,
    WMTS = 2,
}

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
    4:ServiceCategory service_category
    5:string crs
    6:Extent extent
    7:binary thumbnail
}

struct GetCollectionsResponse {
    1: base.BaseResp base,
    2: list<Collection> collections
    3: i64 number_returned
}

struct AddCollectionResponse {
    1: base.BaseResp base,
    2: list<string> service_keys
}

struct PublishRequest {
    1: source.SourceCategory source_category,
    2: string source_key
    3: ServiceCategory service_category
}

service ServiceCollection{
    GetCollectionsResponse GetCollections(1:i64 pageSize, 2:i64 page)
    AddCollectionResponse AddCollection(1:string source_key, 2:string title)
    void Publish(1:PublishRequest req)
}