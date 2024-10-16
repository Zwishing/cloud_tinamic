import io
import logging
import random
import asyncio

import datashader as ds
import datashader.transfer_functions as tf
import geopandas as gpd
from dask import delayed, compute
from osgeo import ogr
from shapely.wkb import loads
import fsspec

# 配置日志输出
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')


async def load_feature_geometry(feature):
    """异步加载要素几何并转换为 shapely 几何对象，返回 None 表示失败"""
    geom_ref = feature.GetGeometryRef()

    if geom_ref is not None:
        geom_wkb = geom_ref.ExportToWkb()
        if isinstance(geom_wkb, (bytearray, bytes)):
            return loads(bytes(geom_wkb))  # 将WKB转换为shapely几何对象
    logging.warning("几何数据无效或缺失")
    return None


async def load_geometries_from_dataset(dataset_path, sample_size=50000):
    """异步并行加载指定路径的数据集，并返回采样后的几何数据列表"""
    try:
        dataset = ogr.Open(dataset_path)
        if dataset is None:
            raise FileNotFoundError(f"无法打开数据集：{dataset_path}")

        layer = dataset.GetLayer()
        total_features = layer.GetFeatureCount()
        sample_size = min(sample_size, total_features)  # 限制采样数量
        sample_indices = random.sample(range(total_features), sample_size)

        # 使用异步方式并行加载几何数据
        tasks = [asyncio.create_task(load_feature_geometry(layer.GetFeature(idx))) for idx in sample_indices]
        geometries = await asyncio.gather(*tasks)

        # 过滤掉 None 值
        geometries = [geom for geom in geometries if geom is not None]

        logging.info(f"成功加载 {len(geometries)} 个几何对象")
        return geometries

    except Exception as e:
        logging.error(f"加载数据集失败: {e}")
        return None


async def generate_thumbnail(geometries, width, height, cmap,
                             add_border=True):
    """异步生成包含加粗边界的缩略图并保存为8-bit图片"""
    try:
        # 将几何数据转换为 GeoDataFrame
        gdf = gpd.GeoDataFrame(geometry=geometries)

        # 创建 Datashader Canvas
        canvas = ds.Canvas(plot_width=width, plot_height=height)

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

        img = tf.shade(agg, cmap=cmap, how='eq_hist')

        # 如果是多边形，且需要添加边界线
        if geom_type in ['Polygon', 'MultiPolygon'] and add_border:
            contour = canvas.line(gdf, geometry='geometry', agg=ds.count(), line_width=0.5)
            contour_img = tf.shade(contour, cmap=["black"], how='linear')
            img = tf.stack(img, contour_img, how="over", alpha=0.9)  # 叠加边界到缩略图上

        # 将 Datashader 图像转换为 PIL 图像对象
        pil_img = img.to_pil()

        # 将图片转换为 8-bit 调色板模式，确保是8-bit
        pil_img = pil_img.convert("P")  # 使用适应性调色板，最多256种颜色

        img_byte_arr = io.BytesIO()
        pil_img.save(img_byte_arr, format='PNG')
        img_byte_arr.seek(0)  # 移动到缓冲区的开始位置
        img_binary = img_byte_arr.getvalue()  # 读取二进制数据

        logging.info("8-bit缩略图的二进制数据已生成")
        return img_binary  # 返回二进制数据

    except Exception as e:
        logging.error(f"生成缩略图失败: {e}")
        return None


async def create_thumbnail_from_vector(vector_path, sample_size=50000, width=300, height=800, cmap=["lightblue"],
                                       add_border=True):
    """从矢量数据异步生成缩略图的通用函数，支持点、线、面及多几何类型"""
    geometries = await load_geometries_from_dataset(vector_path, sample_size)
    if geometries:
        return await generate_thumbnail(geometries, width, height, cmap, add_border)
    else:
        logging.error("无法生成缩略图，几何数据加载失败")
        return None


def generate_thumbnail_from_parquet(path, width, height, cmap,
                             add_border=True):
    """异步生成包含加粗边界的缩略图并保存为8-bit图片"""
    try:
        with fsspec.open(path) as file:
            gdf = gpd.read_parquet(file)
        # 将几何数据转换为 GeoDataFrame
        #     gdf = gpd.read_parquet(path,columns=None)

        # 创建 Datashader Canvas
        canvas = ds.Canvas(plot_width=width, plot_height=height)

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

        img = tf.shade(agg, cmap=cmap, how='eq_hist')

        # 如果是多边形，且需要添加边界线
        if geom_type in ['Polygon', 'MultiPolygon'] and add_border:
            contour = canvas.line(gdf, geometry='geometry', agg=ds.count(), line_width=0.5)
            contour_img = tf.shade(contour, cmap=["black"], how='linear')
            img = tf.stack(img, contour_img, how="over", alpha=0.9)  # 叠加边界到缩略图上

        # 将 Datashader 图像转换为 PIL 图像对象
        pil_img = img.to_pil()

        # 将图片转换为 8-bit 调色板模式，确保是8-bit
        pil_img = pil_img.convert("P")  # 使用适应性调色板，最多256种颜色

        img_byte_arr = io.BytesIO()
        pil_img.save(img_byte_arr, format='PNG')
        img_byte_arr.seek(0)  # 移动到缓冲区的开始位置
        img_binary = img_byte_arr.getvalue()  # 读取二进制数据

        logging.info("8-bit缩略图的二进制数据已生成")
        return img_binary  # 返回二进制数据

    except Exception as e:
        logging.error(f"生成缩略图失败: {e}")
        return None
