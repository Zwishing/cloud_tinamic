import io
import logging
import random
import asyncio

import datashader as ds
import datashader.transfer_functions as tf
import geopandas as gpd
from dask import delayed, compute, dataframe
from osgeo import ogr
from shapely.wkb import loads
import fsspec

# 配置日志输出
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')


def create_thumbnail(s3_path, plot_width=400, plot_height=400,cmap=["lightblue"], add_border=True):
    import time
    s1 = time.time()
    gdf = gpd.read_parquet(s3_path, storage_options={
        'key': "ugNa8yDGzk4gESCATs06",
        'secret': "ruiC05DkvnxxNZrMba5kUbgux8oJLreYuulXhryw",
        'client_kwargs': {
            'endpoint_url': "http://39.101.164.253:9000",
        }
    })
    s2 = time.time()
    # 创建 Datashader Canvas
    canvas = ds.Canvas(plot_width=plot_width, plot_height=plot_height)
    # 获取几何类型并根据类型处理
    geom_type = gdf.geometry.iloc[0].geom_type

    if geom_type in ['Point', 'MultiPoint']:  # 如果是点或多点数据
        agg = canvas.points(gdf, 'geometry', agg=ds.count())
    elif geom_type in ['LineString', 'MultiLineString']:  # 如果是线或多线数据
        agg = canvas.line(gdf, 'geometry', agg=ds.count())
    elif geom_type in ['Polygon', 'MultiPolygon']:  # 如果是面或多面数据
        agg = canvas.polygons(gdf, 'geometry', agg=ds.count())
    else:
        raise ValueError(f"Unsupported geometry type: {geom_type}")
    s3 = time.time()
    img = tf.shade(agg, cmap=cmap, how='eq_hist')
    # 如果是多边形，且需要添加边界线
    if geom_type in ['Polygon', 'MultiPolygon'] and add_border:
        contour = canvas.line(gdf, geometry='geometry', agg=ds.count(), line_width=0.5)
        contour_img = tf.shade(contour, cmap=["black"], how='linear')
        img = tf.stack(img, contour_img, how="over", alpha=0.9)  # 叠加边界到缩略图上
    s4 = time.time()
    # 将 Datashader 图像转换为 PIL 图像对象
    pil_img = img.to_pil()
    s5 = time.time()
    # 将图片转换为 8-bit 调色板模式，确保是8-bit
    pil_img = pil_img.convert("P")  # 使用适应性调色板，最多256种颜色

    img_byte_arr = io.BytesIO()
    pil_img.save(img_byte_arr, format='PNG')
    img_byte_arr.seek(0)  # 移动到缓冲区的开始位置
    img_binary = img_byte_arr.getvalue()  # 读取二进制数据
    s6 = time.time()
    logging.info("8-bit缩略图的二进制数据已生成")
    print(s2 - s1, s3 - s2, s4 - s3, s5 - s4, s6 - s5)
    return img_binary  # 返回二进制数据
