use std::io::Read;
use anyhow::{anyhow, Result};
use tempfile::NamedTempFile;
use crate::{service, util};

pub async fn store_vector(url: &str, schema: &str, table: &str) -> Result<()> {
    let mut temp = NamedTempFile::new().map_err(|e| anyhow!(e.to_string()))?;
    println!("222222");
    util::zipshp2sql(url, temp.path(), schema, table)
        .map_err(|e| anyhow!(e.to_string()))?;

    let mut sql = String::new();
    temp.read_to_string(&mut sql)
        .map_err(|e| anyhow!(e.to_string()))?;
    // println!("{:?}",sql);
    service::store_postgis(&sql).await
        .map_err(|e| anyhow!(e.to_string()))?;

    Ok(())
}
