import unittest
import os
import asyncio
from PIL import Image
import io

from thumbnail import create_thumbnail_from_vector, generate_thumbnail_from_parquet


class TestVectorThumbnail(unittest.TestCase):
    def setUp(self):
        self.vector_data_path1 = "/vsizip//vsicurl/http://39.101.164.253:9000/vector/city.zip"
        self.vector_data_path = "http://39.101.164.253:9000/vector/city.parquet"

    def test_create_thumbnail_from_vector(self):
        async def run_test():
            import time
            start_time = time.time()

            result = await create_thumbnail_from_vector(self.vector_data_path1, sample_size=10000, width=400, height=400)

            end_time = time.time()
            processing_time = end_time - start_time
            print(f"创建缩略图所用时间: {processing_time:.2f} 秒")

            self.assertIsNotNone(result)
            self.assertIsInstance(result, bytes)

            # 读取byte生成图片
            image = Image.open(io.BytesIO(result))
            self.assertIsInstance(image, Image.Image)

            # 选择保存图片进行进一步验证
            image.save("thumbnail_test.png")

        asyncio.run(run_test())

    def test_generate_thumbnail_from_parquet(self):
        import time

        start_time = time.time()
        result = generate_thumbnail_from_parquet(self.vector_data_path, width=400, height=400, cmap=["lightblue"])
        end_time = time.time()
        processing_time = end_time - start_time
        print(f"创建缩略图所用时间: {processing_time:.2f} 秒")

        # 读取byte生成图片
        image = Image.open(io.BytesIO(result))

        # 选择保存图片进行进一步验证
        image.save("thumbnail_test1.png")


if __name__ == '__main__':
    unittest.main()