use std::path::Path;
use gdal::Dataset;
use crate::programs::vector::vector_translate;

pub fn zipshp2sql(url: &str, out: &Path, schema: &str, table: &str) -> Result<(), Box<dyn std::error::Error>> {
    let schema = format!("SCHEMA={}", schema);
    let src = Dataset::open(url)?;
    let opts = Some(vec![
        "-f", "PGDump",
        "-t_srs", "EPSG:4326",
        "-nln", table,
        "-lco", "GEOMETRY_NAME=geom",
        "-lco", "FID=gid",
        "-lco", &schema,
        "-lco", "CREATE_SCHEMA=OFF",
        "-lco", "GEOM_COLUMN_POSITION=END"
    ]);
    vector_translate(&[src], out, opts)?;
    Ok(())
}

pub fn add_prefix_from_ext(url: &str, ext: &str) -> String {
    let prefix = match ext.to_lowercase().as_str() {
        "shp" => "/vsicurl/",
        "zip" => "/vsizip//vsicurl/",
        _ => "/vsicurl/",
    };
    format!("{}{}", prefix, url)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_zipshp2sql() {
        let path = "/vsizip//vsicurl/http://39.101.164.253:9000/vector/九段线.zip";
        // TODO: Implement actual test
    }

    #[test]
    fn test_add_prefix_from_ext() {
        let url = "http://example.com/data";
        let ext = "SHP";
        let result = add_prefix_from_ext(url, ext);
        assert_eq!("/vsicurl/http://example.com/data", result);
    }
}
