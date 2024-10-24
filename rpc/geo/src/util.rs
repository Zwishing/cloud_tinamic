use std::default::Default;
use std::{ffi::CString, path::Path, ptr::{self, null, null_mut}};
use std::io::Write;
use std::sync::{Arc, Mutex};
use gdal::Dataset;
use gdal::vector::Feature;
use gdal_sys::{GDALVectorTranslate,GDALVectorTranslateOptions};
use crate::programs::vector::vector_translate::{vector_translate,VectorTranslateOptions};
use crate::config::get_settings;
use crate::programs::destination::DatasetDestination;
// use geoarrow::io::{gdal::read_gdal,parquet::write_geoparquet_async};
// use geoarrow::io::parquet::GeoParquetWriterOptions;
use minio::s3::types::FileHeaderInfo::NONE;
use tokio::io::AsyncWrite;

pub fn zipshp2sql<P: AsRef<Path>>(url: P, out: &Path, schema: &str, table: &str) -> Result<(), anyhow::Error> {
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
            "-lco", "GEOM_COLUMN_POSITION=END",
            "-lco", "LAUNDER=NO",
            "-makevalid",
        ]
            .try_into()?
    );
    vector_translate(&[src], out.try_into()?, opts)?;
    Ok(())
}

pub fn to_geoparquet<P: AsRef<Path>>(url: P, out: &Path)-> Result<(), anyhow::Error>{
    let src = Dataset::open(url)?;
    let opts = Some(
        vec![
            "-f", "Parquet",
            "-t_srs", "EPSG:4326",
            "-lco", "FID=gid",
            "-makevalid",
            // "skipfailures"
        ]
            .try_into()?
    );
    vector_translate(&[src], out.try_into()?, opts)?;
    Ok(())
}

// pub async fn gdal_to_geoparquet<P: AsRef<Path>, F: AsyncWrite + Unpin + Send>(url: P, f: F) -> Result<(), anyhow::Error> {
//     let dataset = Dataset::open(url)?;
//
//     // 用内层作用域限制 layer 的生命周期
//     let reader = {
//         // 获取图层
//         let mut layer = dataset.layer(0)?;
//         // 从 layer 中读取数据
//         let geoarrow_data = read_gdal(&mut layer, None)?;
//
//         // 检查并将 geoarrow.wkb 解析为 GeoArrow 原生类型
//         geoarrow_data
//             .into_iter()
//             .map(|feature| match feature {
//
//             })
//             .collect::<Result<Vec<_>, _>>()? // 错误处理
//     };
//
//     // 将解析后的数据写入 GeoParquet 文件
//     write_geoparquet_async(reader, f, &GeoParquetWriterOptions::default()).await?;
//     Ok(())
// }


pub async fn vector_to_pg(url: &str, schema: &str, table: &str)->Result<(), anyhow::Error>{
    let schema = format!("SCHEMA={}", schema);
    let dst_connection  = get_settings().await.lock().await.to_pg_string();
    let dst_connection = CString::new(dst_connection).unwrap();
    let src = Dataset::open(url)?;
    // gdal::config::set_config_option("PG_USE_COPY", "YES").unwrap();
    let opts: Option<VectorTranslateOptions> = Some(
        vec![
            "-t_srs", "EPSG:4326",
            "-nln", table,
            "-lco", "OVERWRITE=NO",
            "-lco", "GEOMETRY_NAME=geom",
            "-lco", "FID=gid",
            "-lco", schema.as_str(),
            "-lco", "LAUNDER=NO",
            "-makevalid",
        ]
            .try_into()?
    );

    vector_translate(&[src], DatasetDestination::Path(dst_connection), opts)?;

    
    // let c_options = options
    //     .as_ref()
    //     .map(|x| x.c_options() as *const GDALVectorTranslateOptions)
    //     .unwrap_or(null());
    //
    // let datasets = [&src];
    //         // .iter()
    //         // .map(|x|x.borrow())
    //         // .collect::<Vec<&Dataset>>();
    //
    //         let dataset_out = unsafe {
    //             // Get raw handles to the datasets
    //             let mut datasets_raw: Vec<gdal_sys::GDALDatasetH> =
    //                 datasets.iter().map(|x| x.c_dataset()).collect();
    //
    //             let data = GDALVectorTranslate(
    //                 dst_connection.as_ptr(),
    //                 ptr::null_mut(),
    //                 // only 1 supported currently
    //                 1,
    //                 datasets_raw.as_mut_ptr(),
    //                 c_options,
    //                 null_mut(),
    //             );
    //
    //             // GDAL takes the ownership of `h_dst_ds`
    //             // dest.do_no_drop_dataset();
    //             data
    //         };
    
    Ok(())
}

pub fn add_prefix_from_ext(path: &str, bucket_name:&str, ext: Option<&str>) -> String {
    let ext = ext.map(|e| e.to_lowercase())
        .or_else(|| path.split('.').last().map(|s| s.to_lowercase()));

    let prefix = match ext.as_deref() {
        Some("shp") => "/vsis3/",
        Some("parquet") => "/vsis3/",
        Some("zip") => "/vsizip//vsis3/",
        _ => "/vsis3/",
    };
    format!("{}{}{}", prefix, bucket_name, path)
}

#[cfg(test)]
mod tests {
    use std::path;
    use tokio::fs::File;
    use crate::gdal_config;

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
        let result = add_prefix_from_ext(url, "",Some(ext));
        assert_eq!("/vsicurl/http://example.com/data", result);
    }
    // #[test]
    // fn test_generate_shp_thumbnail() {
    //     let path = "/vsizip//vsicurl/http://39.101.164.253:9000/vector/九段线.zip";
    //     generate_shp_thumbnail(path,"thumbnail.png",200,200);
    // }

    #[tokio::test]
    async fn test_vector_to_pg() {
        //let path = "/vsizip//vsicurl/http://39.101.164.253:9000/vector/九段线.zip";
        // let path = "/vsizip//vsicurl/http://39.101.164.253:9000/vector/九段线.zip";
        let path = "/mnt/d/U盘文件/石漠化监测数据/石漠化一期数据2005_2_fix.shp";
        // let path = "/mnt/d/U盘文件/断陷盆地/断陷盆地轮廓.shp";
        
        let err = vector_to_pg(path, "public", "sandy2").await;
        println!("{:?}", err.err())
    }
    #[tokio::test]
    async fn test_to_geoparquet() {
        let _ = gdal_config::init_gdal_config();
        // gdal::config::set_config_option("AWS_VIRTUAL_HOSTING", "FALSE").unwrap();
        // gdal::config::set_config_option("AWS_HTTPS", "NO").unwrap();
        // gdal::config::set_config_option("AWS_S3_ENDPOINT", "39.101.164.253:9000").unwrap();
        // gdal::config::set_config_option("AWS_REGION", "us-east-1").unwrap();
        // gdal::config::set_config_option("AWS_SECRET_ACCESS_KEY", "ruiC05DkvnxxNZrMba5kUbgux8oJLreYuulXhryw").unwrap();
        // gdal::config::set_config_option("AWS_ACCESS_KEY_ID", "ugNa8yDGzk4gESCATs06").unwrap();
        // gdal::config::set_config_option("AWS_HTTPS", "NO").unwrap();
       
        let path = "/vsizip//vsis3/original-source/vector/city01.zip";
        to_geoparquet(path, Path::new("city.parquet")).expect("Failed to convert to GeoParquet");
    }

    // #[tokio::test]
    // async fn test_gdal_to_geoparquet(){
    //     let _ = gdal_config::init_gdal_config();
    //     let path = "/vsizip//vsis3/original-source/vector/city01.zip";
    //     let f = File::create("city2.parquet").await.unwrap();
    //     gdal_to_geoparquet(path,f).await.unwrap();
    // }
}
