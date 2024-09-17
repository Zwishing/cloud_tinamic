mod db;
mod handler;
mod util;
mod service;
mod programs;

use volo_gen::data::storage::{StoreService,StoreRequest,StoreResponse};
use volo_gen::base::{BaseResp,Code};
use crate::handler::{store_vector};
pub struct S;

impl StoreService for S {
    async fn vector_storage(
        &self,
        _req: StoreRequest,
    ) -> Result<
        StoreResponse,
        volo_thrift::ServerError,
    > {
        let url = _req.url.as_str();
        let schema = _req.schema.as_str();
        let table = _req.table.as_str();
        let ext = _req.ext.as_str();
        
        // 添加调用的前缀
        let url = util::add_prefix_from_ext(url,ext);

        let mut resp = StoreResponse::default();
        let mut base = BaseResp::default();
        match store_vector(url.as_str(), schema, table).await {
            Ok(_)=>{
                base.code = Code::SUCCESS;
                base.msg = String::from("存储成功").parse().unwrap();
            },
            Err(e)=>{
                base.code = Code::FAIL;
                base.msg = e.to_string().parse().unwrap();
            }
        }
        
        resp.base = base;
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



