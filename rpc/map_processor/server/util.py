def is_url(path):
    """检查路径是否是 URL"""
    return path.startswith(("http://", "https://"))


def preprocess_path(path):
    """处理输入路径，根据需要添加前缀"""
    if is_url(path):
        prefix = "/vsizip//vsicurl/" if path.endswith(".zip") else "/vsicurl/"
        return f"{prefix}{path}"
    return path
