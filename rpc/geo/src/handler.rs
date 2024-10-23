use std::{io::Read, path::Path};
use anyhow::{anyhow, Result};
use tempfile::NamedTempFile;
use crate::{service, util};


pub async fn store_vector(url: &str, schema: &str, table: &str) -> Result<()> {
    let mut temp = NamedTempFile::new().map_err(|e| anyhow!(e.to_string()))?;
    
    util::zipshp2sql(url, temp.path(), schema, table)
        .map_err(|e| anyhow!(e.to_string()))?;

    let mut sql = String::new();
    temp.read_to_string(&mut sql)
        .map_err(|e| anyhow!(e.to_string()))?;
    
    tracing::info!("successfully read {} to sql",url);
    
    service::store_postgis(&sql).await
        .map_err(|e| anyhow!(e.to_string()))?;
    
    tracing::info!("Successfully stored vector to postgis in {}.{}", schema, table);
    Ok(())
}

pub async fn store_large_vector(url: &str, schema: &str, table: &str)->Result<()>{
    util::vector_to_pg(url,schema,table).await
}



pub async fn unified_vector<P: AsRef<Path>>(url: P, bucket_name:&str, name:&str) -> Result<u64>{
    service::to_geoparquet_and_upload(url, bucket_name, name).await
}
