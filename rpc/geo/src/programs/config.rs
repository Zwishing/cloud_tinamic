use std::{ffi::CString, path::Path};
use anyhow::Result;
use gdal_sys::{CPLFreeConfig, CPLLoadConfigOptionsFromFile};

pub struct CPLConfigLoader;

impl CPLConfigLoader {
    pub fn load_config_options_from_file<P: AsRef<Path>>(file_path: P, is_override_env_vars: Option<bool>) -> Result<()> {
        let config_file_cstr = CString::new(file_path.as_ref().to_str().ok_or_else(|| anyhow::anyhow!("Invalid path"))?)?;
        let b_override_env_vars = if is_override_env_vars.unwrap_or(false) { 1 } else { 0 };
        unsafe {
            CPLLoadConfigOptionsFromFile(config_file_cstr.as_ptr(), b_override_env_vars);
        }
        Ok(())
    }
}

impl Drop for CPLConfigLoader {
    fn drop(&mut self) {
        unsafe {
            CPLFreeConfig();
        }
    }
}
