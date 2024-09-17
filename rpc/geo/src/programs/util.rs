use std::ffi::{CStr, CString};
use std::path::Path;
use gdal::errors::GdalError;
use libc::c_char;

pub fn _last_null_pointer_err(method_name: &'static str) -> GdalError {
    let last_err_msg = _string(unsafe { gdal_sys::CPLGetLastErrorMsg() });
    unsafe { gdal_sys::CPLErrorReset() };
    GdalError::NullPointer {
        method_name,
        msg: last_err_msg,
    }
}

pub fn _string(raw_ptr: *const c_char) -> String {
    let c_str = unsafe { CStr::from_ptr(raw_ptr) };
    c_str.to_string_lossy().into_owned()
}

pub fn _path_to_c_string(path: &Path) -> gdal::errors::Result<CString> {
    let path_str = path.to_string_lossy();
    CString::new(path_str.as_ref()).map_err(Into::into)
}