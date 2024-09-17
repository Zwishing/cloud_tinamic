use std::ffi::CString;
use std::ptr::{null, null_mut};
use libc::c_char;
use gdal_sys::{ GDALTranslate, GDALTranslateOptions, GDALTranslateOptionsFree};
use gdal::{Dataset,errors::{Result,GdalError}};
use super::super::destination::DatasetDestination;
use super::super::util::_last_null_pointer_err;

pub struct TranslateOptions{
    c_options: *mut GDALTranslateOptions
}

impl TranslateOptions{
    /// See [GDALTranslateOptionsNew].
    ///
    /// [GDALTranslateOptionsNew]: https://gdal.org/api/gdal_utils.html#_CPPv429GDALTranslateOptionsNewPPcP35GDALTranslateOptionsForBinary
    pub fn new<S:Into<Vec<u8>>,I:IntoIterator<Item=S>>(args:I)-> Result<Self> {
        let cstr_args = args
            .into_iter()
            .map(CString::new)
            .collect::<std::result::Result<Vec<_>,_>>()?;
        let mut c_args = cstr_args
            .iter()
            .map(|x| x.as_ptr() as *mut c_char)
            .chain(std::iter::once(null_mut()))
            .collect::<Vec<_>>();

        unsafe {
            Ok(Self {
                c_options: gdal_sys::GDALTranslateOptionsNew(c_args.as_mut_ptr(), null_mut()),
            })
        }
    }
}

impl Drop for TranslateOptions {
    fn drop(&mut self) {
        unsafe {
            GDALTranslateOptionsFree(self.c_options);
        }
    }
}

impl TryFrom<Vec<&str>> for TranslateOptions {
    type Error = GdalError;

    fn try_from(value: Vec<&str>) -> Result<Self> {
        TranslateOptions::new(value)
    }
}



/// Converts simple features data between file formats.
///
/// Wraps [GDALTranslate].
/// See the [program docs] for more details.
///
/// [GDALTranslate]: https://gdal.org/api/gdal_utils.html#_CPPv419GDALTranslatePKc12GDALDatasetHiP12GDALDatasetHPK26GDALTranslateOptionsPi
/// [program docs]: https://gdal.org/programs/ogr2ogr.html
///
pub fn translate(src: &Dataset, dest: DatasetDestination, options: Option<TranslateOptions>) -> Result<Dataset> {
    _translate(
        src
        ,dest,
        options
    )
}

fn _translate(datasets: &Dataset, mut dest: DatasetDestination,options:Option<TranslateOptions>)-> Result<Dataset> {

    let (psz_dest_option, _h_dst_ds) = match &dest {
        DatasetDestination::Path(c_path) => (Some(c_path), null_mut()),
        DatasetDestination::Dataset { dataset, .. } => (None, dataset.c_dataset()),
    };

    let psz_dest = psz_dest_option.map(|x| x.as_ptr()).unwrap_or_else(null);

    let c_options = options
        .as_ref()
        .map(|x| x.c_options as *const GDALTranslateOptions)
        .unwrap_or(null());

    let dataset_out = unsafe {
        // Get raw handles to the datasets

        let data = GDALTranslate(
            psz_dest,
            datasets.c_dataset(),
            c_options,
            null_mut(),
        );

        // GDAL takes the ownership of `h_dst_ds`
        dest.do_no_drop_dataset();

        data

    };

    if dataset_out.is_null() {
        return Err(_last_null_pointer_err("GDALTranslate"));
    }

    let result = unsafe { Dataset::from_c_dataset(dataset_out) };

    Ok(result)
}

#[cfg(test)]
mod tests {
    use std::path::Path;
    use gdal::Dataset;
    use super::translate;

    #[test]
    fn test_translate(){
        let path = "fixtures/tinymarble_cog.img";
        let dataset = &Dataset::open(Path::new(path)).unwrap();
        let out = "fixtures/tinymarble_cog.tif";
        translate(dataset,out.try_into().unwrap(),None).unwrap();
        
    }
}