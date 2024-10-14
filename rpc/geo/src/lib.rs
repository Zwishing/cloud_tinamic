mod db;
mod handler;
mod util;
mod service;
mod programs;
mod config;
mod minio;

use std::future::Future;
use volo_thrift::ServerError;
use volo_gen::data::storage::{StoreService, StoreRequest, StoreResponse, ToGeoParquetStorageRequest, ToGeoParquetStorageResponse};
use volo_gen::base::{BaseResp,Code};
use crate::handler::store_vector;
pub struct S;

impl StoreService for S {
    async fn vector_storage(
        &self,
        req: StoreRequest,
    ) -> Result<
        StoreResponse,
        volo_thrift::ServerError,
    > {
        let url = util::add_prefix_from_ext(&req.url, &req.ext);
        let base = match store_vector(&url, &req.schema, &req.table).await {
            Ok(_) => BaseResp {
                code: Code::SUCCESS,
                msg: format!("success to store vector to postgis in {}.{}", req.schema, req.table).into(),
            },
            Err(e) => BaseResp {
                code: Code::FAIL,
                msg: e.to_string().into(),
            },
        };

        Ok(StoreResponse { base })
    }

    fn to_geo_parquet_storage(&self, req: ToGeoParquetStorageRequest) -> impl Future<Output=Result<ToGeoParquetStorageResponse, ServerError>> + Send {
        todo!()
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



