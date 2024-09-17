
use std::net::SocketAddr;

use geo_storage::{S,LogLayer};

use volo::net::Address;
use volo_gen::data::storage::StoreServiceServer;

#[volo::main]
async fn main() {
    let addr: SocketAddr = "0.0.0.0:8089".parse().unwrap();
    let addr = Address::from(addr);

    StoreServiceServer::new(S)
        .layer_front(LogLayer)
        .run(addr)
        .await
        .unwrap();
}
