use std::path::Path;

use tempfile::NamedTempFile;
use crate::{config::get_pool, util,minio};
use anyhow::anyhow;

pub async fn store_postgis(sql: &str) -> Result<(), anyhow::Error> {
    let pool = get_pool().await.lock().await;

    let client = pool.get().await  // 获取数据库客户端连接
        .map_err(|e| anyhow::anyhow!("Failed to get a client: {:?}", e))?;

    tracing::info!("Sending postgis");
    client.batch_execute(sql)
        .await
        .map_err(|e| anyhow::anyhow!("Failed to execute SQL: {:?}", e))
}

pub async fn to_geoparquet_and_upload<P: AsRef<Path>>(url:P,bucket_name:&str, name:&str)->Result<u64,anyhow::Error>{
    let temp = NamedTempFile::new().map_err(|e| anyhow!(e.to_string()))?;
    util::to_geoparquet(url,temp.path()).map_err(|e| anyhow!(e.to_string()))?;

    // 读取临时文件的大小
    let file_size = temp.as_file().metadata().map(|m| m.len()).unwrap_or(0);

    // 上传文件到minio
    minio::upload_to_minio(temp.path(), bucket_name, name).await?;
    Ok(file_size)
}
