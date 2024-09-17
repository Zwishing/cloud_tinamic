use std::io::{Read};
use anyhow::anyhow;
use tempfile::NamedTempFile;
use crate::{service, util};

pub async fn store_vector(url:&str, schema:&str, table:&str) -> Result<(),anyhow::Error> {
    // Create a temporary file
    let mut temp = match NamedTempFile::new() {
        Ok(file) => file,
        Err(e) => return Err(anyhow!(e.to_string()))
    };

    // Call your utility function (assuming it returns a Result)
    if let Err(e) = util::zipshp2sql(url, temp.path(), schema, table){
        return Err(anyhow!(e.to_string()));
    }

    // Read the temporary file into a string
    let mut sql = String::new();
    if let Err(e) = temp.read_to_string(&mut sql) {
        return Err(anyhow!(e.to_string()));
    }

    // Store the SQL in PostGIS
    if let Err(e) = service::store_postgis(sql.as_str()).await {
        return Err(anyhow!(e.to_string()));
    }

    // Respond with a success message using ApiResponse
   return Ok(())
}
