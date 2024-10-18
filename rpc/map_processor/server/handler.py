from base.ttypes import BaseResp
from map.processor import ttypes, MapProcessorService
from thumbnail import create_thumbnail
from util import preprocess_path


class MapProcessorServiceHandler:
    def VectorThumbnail(self, req):
        file_path = preprocess_path(req.file_path)
        thumbnail = create_thumbnail(file_path,req.width,req.height)
        base_resp = BaseResp(code=0, msg="success")
        return ttypes.VectorThumbnailRespose(base=base_resp, thumbnail=thumbnail)
