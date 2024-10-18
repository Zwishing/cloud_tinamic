
use std::net::SocketAddr;

use geo_storage::{S,LogLayer, gdal_config::init_gdal_config};


use volo::net::Address;
use volo_gen::data::storage::StoreServiceServer;

#[volo::main]
async fn main() {
    init_logging();
    let addr: SocketAddr = "0.0.0.0:8089".parse().unwrap();
    let addr = Address::from(addr);
    match init_gdal_config() {
        Ok(_) => tracing::info!("GDAL configuration initialized successfully"),
        Err(e) => tracing::error!("Failed to initialize GDAL configuration: {}", e),
    };
    
    StoreServiceServer::new(S)
        .layer_front(LogLayer)
        .run(addr)
        .await
        .unwrap();
}



fn init_logging() {
    tracing_subscriber::fmt()
        .with_max_level(tracing::Level::INFO) // 设置最大日志等级
        .init();
}