def is_url(path):
    """检查路径是否是 URL"""
    return path.startswith(("http://", "https://"))


def preprocess_s3path(bucket: str, path: str) -> str:
    """处理输入路径，根据需要添加前缀
    "s3://cloud-optimized-source/vector/116329e3-b8df-49eb-826a-75272f49c4cd.parquet"
    """
    return f"s3://{bucket}{path}"
