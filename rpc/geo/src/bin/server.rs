use std::net::SocketAddr;

use rgeo::S;

#[volo::main]
async fn main() {
    let addr: SocketAddr = "[::]:8080".parse().unwrap();
    let addr = volo::net::Address::from(addr);

    volo_gen::geo::storage::VectorStoreServiceServer::new(S)
        .run(addr)
        .await
        .unwrap();
}
