
use std::path::Path;
use minio::s3::builders::ObjectContent;

use crate::config::get_minio_client;

pub async fn upload_to_minio(file_path: &Path, bucket: &str, object_name: &str) -> Result<(), anyhow::Error> {
    let client = get_minio_client().await.lock().await;
    // // 检查桶是否存在
    // if !client.bucket_exists(&BucketExistsArgs::new(bucket)?).await? {
    //     client.make_bucket(&MakeBucketArgs::new(bucket)?).await?;
    // }

    let content = ObjectContent::from(file_path);
    // Put an object
    client
        .put_object_content(bucket, object_name, content)
        .send()
        .await?;

    tracing::info!("文件上传成功！");

    Ok(())
}

#[cfg(test)]
mod tests {
    use std::path::Path;

    use crate::minio::upload_to_minio;

    #[tokio::test]
    async fn test_upload_to_minio(){
        // let path = "/vsizip//vsicurl/http://39.101.164.253:9000/vector/city.zip";
        let result = upload_to_minio(Path::new("/Users/zwishing/code/go/cloud_tinamic/rpc/geo/city.parquet"),"cloud-optimized-source","city.parquet").await;
        match result {
            Ok(_) => println!("Test upload successful"),
            Err(e) => println!("Test upload failed: {:?}", e),
        }
    }
}
