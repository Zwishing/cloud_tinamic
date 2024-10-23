mod handler;
mod util;
mod service;
mod programs;
mod config;
mod minio;
pub mod gdal_config;

use std::future::Future;
use volo_thrift::ServerError;
use volo_gen::data::storage::{StoreService, ToGeoParquetStorageRequest, ToGeoParquetStorageResponse, VectorToPgStorageRequest, VectorToPgStorageResponse};
use volo_gen::base::{BaseResp,Code};
use crate::handler::{store_large_vector, store_vector, unified_vector};
pub struct S;

impl StoreService for S {
    async fn vector_to_pg_storage(&self, req: VectorToPgStorageRequest) -> Result<
        VectorToPgStorageResponse,
        ServerError,
    > {
        let s3path = util::add_prefix_from_ext(&req.cloud_optimized_path, &req.cloud_optimized_bucket_name, None);
        // let base = match store_vector(&s3path, &req.schema, &req.table).await {
        //     Ok(_) => BaseResp {
        //         code: Code::SUCCESS,
        //         msg: format!("success to store vector to postgis in {}.{}", req.schema, req.table).into(),
        //     },
        //     Err(e) => BaseResp {
        //         code: Code::FAIL,
        //         msg: e.to_string().into(),
        //     },
        // };

        let base = match store_large_vector(&s3path, &req.schema, &req.table).await {
            Ok(_) => BaseResp {
                code: Code::SUCCESS,
                msg: format!("success to store vector to postgis in {}.{}", req.schema, req.table).into(),
            },
            Err(e) => BaseResp {
                code: Code::FAIL,
                msg: e.to_string().into(),
            },
        };

        Ok(VectorToPgStorageResponse { base })
    }

    async fn to_geo_parquet_storage(&self, req: ToGeoParquetStorageRequest) -> Result<ToGeoParquetStorageResponse, ServerError> {
        let s3url = util::add_prefix_from_ext(&req.source_path, &req.source_bucket,None);
        let resp = match unified_vector(&s3url,&req.dest_bucket,&req.dest_path).await{
            Ok(size)=> {
                tracing::info!("success to unified_vector in {}",req.dest_path);
                ToGeoParquetStorageResponse{
                base:BaseResp {
                    code: Code::SUCCESS,
                    msg: format!("success to unified_vector in {}",req.dest_path).into(),
                },
                dest_path: Some(format!("{}",req.dest_path).into()),
                size:Some(size as i64),
            }},
            Err(e)=>{
                tracing::error!("{}", e.to_string());
                ToGeoParquetStorageResponse{
                base: BaseResp {
                    code: Code::FAIL,
                    msg: e.to_string().into(),
                },
                dest_path: None,
                size: None
            }}
        };
        Ok(resp)

    }
}

pub struct LogLayer;

impl<S> volo::Layer<S> for LogLayer {
    type Service = LogService<S>;

    fn layer(self, inner: S) -> Self::Service {
        LogService(inner)
    }
}


#[derive(Clone)]
pub struct LogService<S>(S);

#[volo::service]
impl<Cx, Req, S> volo::Service<Cx, Req> for LogService<S>
where
    Req: std::fmt::Debug + Send + 'static,
    S: Send + 'static + volo::Service<Cx, Req> + Sync,
    S::Response: std::fmt::Debug,
    S::Error: std::fmt::Debug,
    Cx: Send + 'static,
{
    async fn call(&self, cx: &mut Cx, req: Req) -> Result<S::Response, S::Error> {
        let now = std::time::Instant::now();
        tracing::debug!("Received request {:?}", &req);
        let resp = self.0.call(cx, req).await;
        tracing::debug!("Sent response {:?}", &resp);
        tracing::info!("Request took {}ms", now.elapsed().as_millis());
        resp
    }
}



