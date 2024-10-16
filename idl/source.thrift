namespace go data.source

include "base.thrift"

enum SourceCategory {
    VECTOR =  1
    IMAGERY = 2
    TERRAIN = 3
    POINTCLOUD = 4
    PHOTOGRAMMETRY = 5
    SURFACEMODEL = 6
    BIM = 7
}

enum ItemType {
    FILE = 1
    FOLDER = 2
}

struct Item {
    1: string name,
    2: ItemType item_type,
    3: string key,
    4: i64 size,
    5: i64 modified_time,
    6: string path
    7: string parent_key
}

struct UploadRequest{
    1: SourceCategory source_category,
    2: string key,
    3: string name,
    4: i64 size,
    5: binary file_data,
}

struct UploadResponse{
    1: base.BaseResp base,
    2: string key,
}

struct PresignedUploadResquest{
    1: SourceCategory source_category,
    2: string path,
    3: string name,
}

struct PresignedUploadResponse{
    1: base.BaseResp base,
    2: string url
}

struct GetItemsRequest {
    1: SourceCategory source_category,
    2: string key
}

struct GetItemsResponse {
    1: base.BaseResp base,
    2: required string key
    3: required string name
    4: list<Item> items
}

struct GetHomeItemsRequest {
    1: SourceCategory source_category,
}


struct CreateFolderRequest {
    1: required SourceCategory source_category,
    2: required string key,
    3: required string name,
    4: required string path,
}

struct DeleteItemsRequest{
    1: list<string> keys
}

struct DeleteItemsResponse{
    1: base.BaseResp base,
}

struct AddItemRequest{
    1: SourceCategory source_category,
    2: string currentFolder,
    3: Item item,
}

struct AddItemResponse{
    1: base.BaseResp base,
    2: Item item,
}

service SourceService{
    GetItemsResponse GetNextItems(1:GetItemsRequest req),
    GetItemsResponse GetPreviousItems(1:GetItemsRequest req),
    GetItemsResponse GetHomeItems(1:GetHomeItemsRequest req),
    DeleteItemsResponse DeleteItems(1:DeleteItemsRequest req),
    AddItemResponse AddItem(1:AddItemRequest req),
    AddItemResponse CreateFolder(1:CreateFolderRequest req),
    UploadResponse Upload(1:UploadRequest req),
    PresignedUploadResponse PresignedUpload(1:PresignedUploadResquest req),
    string GetSourcePath(1:string key),
    string GetUnifiedSourcePath(1:string source_key),
}
