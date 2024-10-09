use std::{borrow::Borrow, ffi::CString, path::Path, ptr::{self, null, null_mut}};
use gdal::Dataset;
use gdal_sys::{GDALVectorTranslate,GDALVectorTranslateOptions};
use crate::programs::vector::vector_translate::{vector_translate,VectorTranslateOptions};
use crate::db::get_settings;

pub fn zipshp2sql(url: &str, out: &Path, schema: &str, table: &str) -> Result<(), Box<dyn std::error::Error>> {
    let schema = format!("SCHEMA={}", schema);
    let src = Dataset::open(url)?;
    let opts = Some(
        vec![
            "-f", "PGDump",
            "-t_srs", "EPSG:4326",
            "-nln", table,
            "-lco", "GEOMETRY_NAME=geom",
            "-lco", "FID=gid",
            "-lco", schema.as_str(),
            "-lco", "CREATE_SCHEMA=OFF",
            "-lco", "GEOM_COLUMN_POSITION=END"
        ]
            .try_into()?
    );
    vector_translate(&[src], out.try_into()?, opts)?;
    Ok(())
}

pub async fn vector_to_pg(url: &str, schema: &str, table: &str)->Result<(), anyhow::Error>{
    let schema = format!("SCHEMA={}", schema);
    // let pg = format!("PG:{}", "postgresql://postgres:admin321@1.92.113.25:5432/tinamic");
    let pg = format!("PG:{}", "dbname=tinamic host=1.92.113.25 port=5432 user=postgres password=admin321");
    // let dst_connection = CString::new("PG: host=1.92.113.25 user=postgres password=admin321 dbname=tinamic").unwrap();
    let dst_connection  = get_settings().lock().await.to_pg_string();
    let dst_connection = CString::new(dst_connection).unwrap();
    let src = Dataset::open(url)?;
    gdal::config::set_config_option("PG_USE_COPY", "YES").unwrap();
    let options: Option<VectorTranslateOptions> = Some(
        vec![
            "-t_srs", "EPSG:4326",
            "-nln", table,
            "-lco", "GEOMETRY_NAME=geom",
            "-lco", "FID=gid",
            "-lco", schema.as_str(),
            "-makevalid",
            // "-skipinvalid",
            // "-lco", "CREATE_SCHEMA=OFF",
            // "-lco", "GEOM_COLUMN_POSITION=END",
        ]
            .try_into()?
    );
    
    let c_options = options
        .as_ref()
        .map(|x| x.c_options() as *const GDALVectorTranslateOptions)
        .unwrap_or(null());

    let datasets = [&src];
            // .iter()
            // .map(|x|x.borrow())
            // .collect::<Vec<&Dataset>>();

            let dataset_out = unsafe {
                // Get raw handles to the datasets
                let mut datasets_raw: Vec<gdal_sys::GDALDatasetH> =
                    datasets.iter().map(|x| x.c_dataset()).collect();
        
                let data = GDALVectorTranslate(
                    dst_connection.as_ptr(),
                    null_mut(),
                    // only 1 supported currently
                    1,
                    datasets_raw.as_mut_ptr(),
                    c_options,
                    null_mut(),
                );
        
                // GDAL takes the ownership of `h_dst_ds`
                // dest.do_no_drop_dataset();
                data
            };
    
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
    // #[test]
    // fn test_generate_shp_thumbnail() {
    //     let path = "/vsizip//vsicurl/http://39.101.164.253:9000/vector/九段线.zip";
    //     generate_shp_thumbnail(path,"thumbnail.png",200,200);
    // }

    #[tokio::test]
    async fn test_vector_to_pg() {
        // let path = "/vsizip//vsicurl/http://39.101.164.253:9000/vector/九段线.zip";
        let path = "/mnt/d/U盘文件/石漠化监测数据/石漠化一期数据2005_2_fix.shp";
        // let path = "/mnt/d/U盘文件/断陷盆地/断陷盆地轮廓.shp";
        
        let err = vector_to_pg(path, "public", "sandy2").await;
        println!("{:?}", err.err())
    }
}
