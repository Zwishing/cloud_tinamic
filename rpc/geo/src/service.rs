use crate::db::get_pool;

pub async fn store_postgis(sql: &str) -> Result<(), anyhow::Error> {
    let pool = get_pool().await.lock().await;
<<<<<<< HEAD
    let client = pool.get().await  // 获取数据库客户端连接
        .map_err(|e| anyhow::anyhow!("Failed to get a client: {:?}", e))?;
    tracing::info!("Sending postgis");
=======

    let client = pool.get().await  // 获取数据库客户端连接
        .map_err(|e| anyhow::anyhow!("Failed to get a client: {:?}", e))?;

    tracing::info!("Sending postgis");

>>>>>>> refs/remotes/origin/main
    client.batch_execute(sql)
        .await
        .map_err(|e| anyhow::anyhow!("Failed to execute SQL: {:?}", e))
}


