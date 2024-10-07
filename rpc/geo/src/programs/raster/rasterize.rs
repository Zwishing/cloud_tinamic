// use gdal::{cpl::CslStringList, errors::GdalError};

// #[derive(Clone, Copy,Debug)]
// pub enum Dtype {
//     Byte,
//     Int8,
//     Int16,
//     Uint16,
//     Float16
// }
// #[derive(Clone, Copy,Debug)]
// pub enum Size {
//     Resolution([u32;2]),
//     WidthHeight([u32;2]),
// }

// #[derive(Clone, Copy,Debug)]
// pub struct RasterizeOptions<'a> {
//     bands: &'a [u32],
//     burn: &'a [f64],
//     extent:[f64;4],
//     size: Size,
//     dtype: Dtype,
//     nodata:f32
// }

// impl Default for RasterizeOptions {
//     fn default() -> Self {
//         RasterizeOptions {
//         }
//     }
// }

// impl TryFrom<RasterizeOptions> for CslStringList {
//     type Error = GdalError;

//     fn try_from(value: RasterizeOptions) -> Result<Self, Self::Error> {
//         todo!()
//     }
// }

// pub fn rasterize(){

// }