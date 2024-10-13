from base.ttypes import BaseResp
from map.processor import ttypes, MapProcessorService
from thumbnail import create_thumbnail_from_vector
from util import preprocess_path
import asyncio


class MapProcessorServiceHandler:
    def VectorThumbnail(self, req):
        file_path = preprocess_path(req.file_path)
        thumbnail = asyncio.get_event_loop().run_until_complete(
            create_thumbnail_from_vector(vector_path=file_path, width=req.width, height=req.height))
        base_resp = BaseResp(code=0, msg="success")
        return ttypes.VectorThumbnailRespose(base=base_resp, thumbnail=thumbnail)
