from base.ttypes import BaseResp
from map.processor import ttypes, MapProcessorService
from thumbnail import create_thumbnail
from util import preprocess_s3path


class MapProcessorServiceHandler:
    def VectorThumbnail(self, req):
        s3path = preprocess_s3path(req.cloud_optimized_bucket_name,req.cloud_optimized_path)
        thumbnail = create_thumbnail(s3path,req.width,req.height)
        base_resp = BaseResp(code=0, msg="success")
        return ttypes.VectorThumbnailRespose(base=base_resp, thumbnail=thumbnail)
