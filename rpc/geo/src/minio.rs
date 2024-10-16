use minio::s3::args::{BucketExistsArgs, MakeBucketArgs};
use std::path::Path;
use minio::s3::builders::ObjectContent;

use crate::config::get_minio_client;

pub async fn upload_to_minio(file_path: &Path, bucket: &str, object_name: &str) -> Result<(), anyhow::Error> {
    let client = get_minio_client().lock().await;
    // 检查桶是否存在
    if !client.bucket_exists(&BucketExistsArgs::new(bucket)?).await? {
        client.make_bucket(&MakeBucketArgs::new(bucket)?).await?;
    }

    let content = ObjectContent::from(file_path);
    // Put an object
    client
        .put_object_content(bucket, object_name, content)
        .send()
        .await?;

    tracing::info!("文件上传成功！");

    Ok(())
}
