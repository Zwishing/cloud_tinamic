use minio::s3::{
    args::{BucketExistsArgs, MakeBucketArgs, PutObjectArgs},
};
use std::fs::File;
use std::io::Read;

use crate::config::get_minio_client;

async fn upload_to_minio(file_path: &str, bucket: &str, object_name: &str) -> Result<(), anyhow::Error> {
    let client = get_minio_client().lock().await;
    // Check if bucket exists
    if !client.bucket_exists(&BucketExistsArgs::new(bucket)?).await? {
        client.make_bucket(&MakeBucketArgs::new(bucket)?).await?;
    }

    // Read file content
    let mut file = File::open(file_path)?;
    let file_size = file.metadata()?.len();

    // Upload file
    client
        .put_object(&mut PutObjectArgs::new(bucket, object_name, &mut file, Some(file_size as usize), None)?)
        .await?;

    println!("File uploaded successfully!");
    Ok(())
}
