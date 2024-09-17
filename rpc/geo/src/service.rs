use crate::db::POOL;

pub async fn store_postgis(sql: &str) -> Result<(), anyhow::Error> {
    let pool = POOL.lock().await;
    
    let client = pool.get()
        .await
        .map_err(|e| anyhow::anyhow!("Failed to get a client: {:?}", e))?;

    client.batch_execute(sql)
        .await
        .map_err(|e| anyhow::anyhow!("Failed to execute SQL: {:?}", e))
}
