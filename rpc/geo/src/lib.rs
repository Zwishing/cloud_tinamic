pub struct S;

impl volo_gen::geo::storage::VectorStoreService for S {
    async fn storage(
        &self,
        _req: volo_gen::geo::storage::VectorStoreRequest,
    ) -> ::core::result::Result<
        volo_gen::geo::storage::VectorStoreResponse,
        ::volo_thrift::ServerError,
    > {
        ::std::result::Result::Ok(Default::default())
    }
}
