namespace go data.source

include "base.thrift"

enum SourceType {
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
    5: string modified_time,
    6: string path
}

struct UploadRequest{

}

struct UploadResponse{
    1: base.BaseResp base,
}

struct PresignedUploadResquest{
    1: SourceType source_type,
    2: string path,
    3: string name,
}

struct PresignedUploadResponse{
    1: base.BaseResp base,
    2: string url
}

struct GetItemRequest {
    1: SourceType source_type,
    2: string path
    3: string key
}

struct GetItemResponse {
    1: base.BaseResp base,
    2: list<Item> items
}

struct CreateFolderRequest {
    1: SourceType source_type,
    2: string name
    3: string path
}

struct CreateFolderResponse {
    1: base.BaseResp base,
}

struct DeleteItemRequest{
    1: string key
    2: string path
}

struct DeleteItemResponse{
    1: base.BaseResp base,
}

struct AddItemRequest{
    1: SourceType source_type,
    2: string currentFolder,
    3: Item item,
}

struct AddItemResponse{
    1: base.BaseResp base,
}

service SourceService{
    UploadResponse Upload(1:UploadRequest req),
    PresignedUploadResponse PresignedUpload(1:PresignedUploadResquest req),
    GetItemResponse GetItem(1:GetItemRequest req),
    CreateFolderResponse CreateFolder(1:CreateFolderRequest req),
    DeleteItemResponse DeleteItem(1:DeleteItemRequest req),
    AddItemResponse AddItem(1:AddItemRequest req),
}
