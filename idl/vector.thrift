namespace go service.vector

include "base.thrift"

struct Tile {
    1: i64 x,
    2: i64 y,
    3: i8 z,
}

struct GetTileRequest{
    1: string service_id
    2: Tile tile
}

service VectorService{
    void Publish(1:string req)
    binary GetTile(1:GetTileRequest req)
}